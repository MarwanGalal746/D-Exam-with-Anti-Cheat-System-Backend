package handlers

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/dataContainers"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/messaging"
	"log"
	"net/http"
	"os"
)

func StartServer() {

	dataContainers.InitActiveStudents()
	dataContainers.InitCheatStudents()

	go messaging.AddActiveStudent()

	http.HandleFunc("/ws", WsEndpoint)
	log.Fatal(http.ListenAndServe(os.Getenv("OVERSEER_SERVER_PORT"), nil))

}
