// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	"sync"
	"fmt"
	// Add any other necessary imports
)

type Message struct {
    Sender string
    Recipient string
    Content string
}

// Client represents a connected chat client
type Client struct {
	// TODO: Implement this struct
	// Hint: username, message channel, mutex, disconnected flag
	Username string
	incomming chan string
	outgoing chan string
	active bool
	
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	// TODO: Implement this method
	// Hint: thread-safe, non-blocking send
	if !c.active {
		return
	}

	select {
	case c.incomming <- message:
	default:
	}
	
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// TODO: Implement this method
	// Hint: read from channel, handle closed channel
	return <- c.outgoing
}

// ChatServer manages client connections and message routing
type ChatServer struct {
	// TODO: Implement this struct
	// Hint: clients map, mutex
	mu sync.RWMutex
	clients map[string]*Client
	input chan Message
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	// TODO: Implement this function
	server := &ChatServer{
		clients: make(map[string]*Client),
		input:   make(chan Message, 100),
	}
	
	go server.listen()
	
	return server
}

func (s *ChatServer) listen() {
    for msg := range s.input {
        if msg.Recipient == "" {
            s.broadcastMessage(msg)
        } else {
            _ = s.sendPrivateMessage(msg)
        }
    }
}

// Connect adds a new client to the chat server
func (s *ChatServer) Connect(username string) (*Client, error) {
	// TODO: Implement this method
	// Hint: check username, create client, add to map
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _, exists := s.clients[username]; exists {
	    return nil, ErrUsernameAlreadyTaken
	}
	
	client := &Client{
	    Username: username,
	    incomming: make(chan string, 100),
	    outgoing: make(chan string, 100),
	    active: true,
	}
	
	s.clients[username] = client
	
	go func(c *Client) {
	    for msg := range c.incomming {
	        c.outgoing <- msg
	    }
	}(client)
	
	return client, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// TODO: Implement this method
	// Hint: remove from map, close channels
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if _,ok := s.clients[client.Username]; ok {
	    close(client.incomming)
	    close(client.outgoing)
	    delete(s.clients, client.Username)
	}
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	// TODO: Implement this method
	// Hint: format message, send to all clients
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, exists := s.clients[sender.Username]; !exists {
		
	}

	s.input <- Message{Sender: sender.Username, Content: message}
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	// TODO: Implement this method
	// Hint: find recipient, check errors, send message
	s.mu.RLock()
	defer s.mu.RUnlock()

	// ✳️ Sender hâlâ bağlı mı?
	if _, exists := s.clients[sender.Username]; !exists {
		return fmt.Errorf("sender %s not connected", sender.Username)
	}

	// ✳️ Recipient kontrolü
	recipientClient, exists := s.clients[recipient]
	if !exists {
		return ErrRecipientNotFound
	}

	recipientClient.Send(fmt.Sprintf("[Private from %s]: %s", sender.Username, message))
	return nil
}

func (s *ChatServer) broadcastMessage(msg Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for username, client := range s.clients {
		if username != msg.Sender {
			client.Send(fmt.Sprintf("[Broadcast from %s]: %s", msg.Sender, msg.Content))
		}
	}
}

func (s *ChatServer) sendPrivateMessage(msg Message) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	recipient, exists := s.clients[msg.Recipient]
	if !exists {
		return ErrRecipientNotFound
	}

	recipient.Send(fmt.Sprintf("[Private from %s]: %s", msg.Sender, msg.Content))
	return nil
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
)
