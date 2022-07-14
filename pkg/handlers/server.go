package handlers

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := mux.NewRouter()
	//this CORS to enable frontend request to the backend endpoints
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	uploadImgHandler := ImageHandlers{services.NewImagService()}

	//sandwiches endpoints
	router.HandleFunc("/api/upload-img", uploadImgHandler.Upload).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(os.Getenv("UPLOAD_IMG_SERVER_PORT"), handlers.CORS(headers, methods, origins)(router)))

}
