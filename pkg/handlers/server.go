package handlers

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/messaging"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()

	//this CORS to enable frontend request to the backend endpoints
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	pkg.InitActiveStudents()

	go messaging.AddActiveStudent()

	router.HandleFunc("/ws", WsEndpoint)

	log.Fatal(http.ListenAndServe(viper.GetString(viper.GetString("SERVER_PORT")), handlers.CORS(headers, methods, origins)(router)))

}
