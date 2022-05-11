package messaging

import (
	"exam_service/pkg/domain/models"
	"exam_service/pkg/service"
	"fmt"
	"github.com/streadway/amqp"
)

func DeleteCourseExams(repository models.ExamRepository) {
	fmt.Println("Consumer")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to the RabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("RabbitMQ practice",
		"",
		true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	messages := make(chan string)
	go func() {
		for d := range msgs {
			courseId := string(d.Body[:])
			fmt.Println("Received msg: ", courseId)
			examService := service.NewExamService(repository)
			examService.DelCourseExams(courseId)
		}
	}()
	fmt.Println("Successfully consumed msg to the queue")
	<-messages

}
