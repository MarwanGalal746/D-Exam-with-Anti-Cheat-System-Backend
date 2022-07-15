package services

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/domain"
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ImageService interface {
	Upload(domain.Image) (*domain.ImageUrl, error)
}

type DefaultImageService struct {
}

func (s DefaultImageService) Upload(image domain.Image) (*domain.ImageUrl, error) {
	res1 := strings.Split(image.Img, ",")
	res2 := strings.Split(res1[0], "/")
	res3 := strings.Split(res2[1], ";")
	imgType := "." + res3[0]
	err := errors.New("")
	imageDecoded, err := base64.StdEncoding.DecodeString(res1[1])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	url, err := utils.UploadBytesToBlob(imageDecoded, string(time.Now().Unix())+imgType)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &domain.ImageUrl{ImgUrl: url}, nil
}

func NewImagService() DefaultImageService {
	return DefaultImageService{}
}
