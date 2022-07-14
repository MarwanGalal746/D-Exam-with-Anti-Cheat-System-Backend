package handlers

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/domain"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/errs"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/services"
	"encoding/json"
	"net/http"
)

type ImageHandlers struct {
	service services.ImageService
}

func (imageHandler ImageHandlers) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newImage domain.Image
	_ = json.NewDecoder(r.Body).Decode(&newImage)
	imgUrl, err := imageHandler.service.Upload(newImage)
	//handling errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrInternalServerError.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(imgUrl)
}
