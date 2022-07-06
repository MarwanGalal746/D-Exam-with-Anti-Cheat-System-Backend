package handlers

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/dataContainers"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func contains(a string, list []string) (bool, int) {
	for i := 0; i < len(list); i++ {
		if list[i] == a {
			return true, i
		}
	}
	return false, -1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func checkIfStudentCheat(conn *websocket.Conn, studentId string, result chan bool) {
	for {
		cond, ind := contains(studentId, dataContainers.CheatStudents)
		if cond {
			err := conn.WriteMessage(1, []byte("student has cheated"))
			if err != nil {
				log.Println(err)
			}
			err = conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "Connection closed because desktop app is not opened"), time.Now().Add(1*time.Second))
			if err != nil {
				fmt.Println(err)
			}
			dataContainers.CheatStudents = remove(dataContainers.CheatStudents, ind)
			result <- true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func checkIfStudentClosedTool(conn *websocket.Conn, studentId string, result chan bool) {
	for {
		cond, _ := contains(studentId, dataContainers.ActiveStudents)
		if !cond {
			err := conn.WriteMessage(1, []byte("student has closed the desktop tool"))
			if err != nil {
				log.Println(err)
			}
			err = conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "Connection closed because desktop app is not opened"), time.Now().Add(1*time.Second))
			if err != nil {
				fmt.Println(err)
			}
			result <- true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func reader(conn *websocket.Conn) {
	studentId := "no-id"
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		if studentId == "no-id" {
			studentId = string(p)
		} else {
			continue
		}
		// print out that message for clarity
		fmt.Println("socket: " + string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
		cond, _ := contains(string(p), dataContainers.ActiveStudents)
		if !cond {
			err = conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(1000, "Connection closed because desktop app is not opened"), time.Now().Add(1*time.Second))
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		result := make(chan bool, 1)
		go checkIfStudentCheat(conn, studentId, result)

		closeDesktopChannel := make(chan bool, 1)
		go checkIfStudentClosedTool(conn, studentId, closeDesktopChannel)

		value := <-result
		close(result)
		if value {
			return
		}
		isStudClosedTool := <-closeDesktopChannel
		close(closeDesktopChannel)
		if isStudClosedTool {
			return
		}

	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Student!"))
	if err != nil {
		log.Println(err)
	}

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	go reader(ws)
}
