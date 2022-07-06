package handlers

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/dataContainers"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/messaging"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func StartServer() {

	dataContainers.InitActiveStudents()
	dataContainers.InitCheatStudents()

	go messaging.AddActiveStudent()
	go messaging.AddCheatStudent()

	http.HandleFunc("/ws", WsEndpoint)
	log.Fatal(http.ListenAndServe(viper.GetString("SERVER_PORT"), nil))

}
