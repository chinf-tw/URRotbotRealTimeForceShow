package main

import (
	"NTNU_Prj/DeCodeURInterface"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var urAddress string

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")

}

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var poseForce = make(chan []float32)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// helpful log statement to show connections
	log.Println("Client Connected")
	var (
		id   = 0
		pose []float32
	)
	conn, err := net.Dial("tcp", urAddress)
	defer conn.Close()
	for {

		if err != nil {
			fmt.Println(err)
		}

		var tcpForceData []byte

		tcpForceData, err = DeCodeURInterface.RealTime(conn)
		if err != nil {
			fmt.Println(err)
		}

		if tcpForceData != nil {
			for i := 0; i <= 6; i++ {
				pose = append(pose, float32(DeCodeURInterface.Float64frombytes(tcpForceData[i*8:(i+1)*8])))
			}

			time.Sleep(time.Millisecond * 200)
		}
		fmt.Println("pose ready to go")
		dataJSON := ForceData{Force: pose}
		if WebSocketWriter(dataJSON, ws, 1, &id) {
			return
		}
		pose = nil
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	//const network = "tcp"
	//const address = ":21"

	if len(os.Args) != 3 {
		log.Fatal("Please give two args.")
	}

	fmt.Println(os.Args)

	urAddress = os.Args[1]
	listenPort := os.Args[2]
	log.Fatal(http.ListenAndServe(listenPort, nil))
}

// func forcePackagesHandler() {
// 	conn, err := net.Dial("tcp", "192.168.1.107:30003")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer conn.Close()
// 	var tcpForceData []byte
// 	var poseForceData []float32

// 	for {
// 		tcpForceData, err = DeCodeURInterface.RealTime(conn)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		go func() {
// 			if tcpForceData != nil {
// 				for i := 0; i <= 6; i++ {
// 					poseForceData = append(poseForceData, float32(DeCodeURInterface.Float64frombytes(tcpForceData[i*8:(i+1)*8])))
// 				}
// 				poseForce <- poseForceData
// 				time.Sleep(time.Millisecond * 400)
// 			}
// 		}()

// 	}

// }
// func testConnectHandler(c net.Conn) {
// 	defer c.Close()
// 	var pose []float32
// 	for {

// 		data := make([]byte, 1024)

// 		dataLen, err := c.Read(data)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		poseStr := string(data[:dataLen])
// 		pose, err = UR3Handler.PoseTypeToFloatList(poseStr)
// 		poseForce <- pose
// 		if err != nil {
// 			log.Println("UR3Handler.PoseTypeToFloatList : ", err)
// 			return
// 		}
// 		// 釋放重複讀取pose
// 		_, err = c.Read(data)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// 	// Shut down the connection.
// }
