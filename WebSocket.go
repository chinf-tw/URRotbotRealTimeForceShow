package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func WebSocketReader(conn *websocket.Conn) bool {

	// read in a message
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return true
	}
	// print out that message for clarity
	fmt.Println(messageType, string(p))

	return false

}

func WebSocketWriter(dataJSON ForceData, conn *websocket.Conn, messageType int, id *int) bool {
	*id += 1
	dataJSON.ID = *id
	data, err := json.Marshal(dataJSON)
	if err != nil {
		log.Println(err)
		return true
	}
	if err := conn.WriteMessage(messageType, data); err != nil {
		log.Println(err)
		return true
	}
	return false
}
