package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-fiber-websocket/domain/entities"
	"log"
	"text/template"

	"github.com/gofiber/websocket/v2"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *entities.Message
}

type IWebSocketService interface {
	HandleWebSocket(ctx *websocket.Conn)
	HandleMessages()
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.Message),
	}
}

func (sv *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {

	// Register a new Client
	sv.clients[ctx] = true
	defer func() {
		delete(sv.clients, ctx)
		ctx.Close()
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		// send the message to the broadcast channel
		var message entities.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Fatalf("Error Unmarshalling")
		}
		sv.broadcast <- &message
	}
}

func (sv *WebSocketServer) HandleMessages() {
	for {
		msg := <-sv.broadcast

		// Send the message to all Clients

		for client := range sv.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemplate(msg))
			if err != nil {
				log.Printf("Write  Error: %v ", err)
				client.Close()
				delete(sv.clients, client)
			}
			fmt.Println("Message Sent", msg.Text)
		}

	}
}

func getMessageTemplate(msg *entities.Message) []byte {
	tmpl, err := template.ParseFiles("frontend/views/message.html")
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	// Render the template with the message as data.
	var renderedMessage bytes.Buffer
	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
}
