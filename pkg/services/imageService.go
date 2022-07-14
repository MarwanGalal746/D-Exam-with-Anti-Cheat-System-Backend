package services

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/domain"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/utils"
	"encoding/base64"
	"errors"
	"fmt"
)

type ImageService interface {
	Upload(domain.Image) (*domain.ImageUrl, error)
}

type DefaultImageService struct {
}

func (s DefaultImageService) Upload(image domain.Image) (*domain.ImageUrl, error) {

	err := errors.New("")
	imageDecoded, err := base64.StdEncoding.DecodeString(image.Img)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	url, err := utils.UploadBytesToBlob(imageDecoded, image.UserId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &domain.ImageUrl{ImgUrl: url}, nil
}

func NewImagService() DefaultImageService {
	return DefaultImageService{}
}
