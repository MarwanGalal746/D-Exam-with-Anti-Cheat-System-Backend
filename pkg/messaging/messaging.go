package messaging

import (
	"exam_service/pkg/domain/models"
	"exam_service/pkg/service"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
)

func DeleteCourseExams(repository models.ExamRepository) {
	log.Println("Consumer")

	rabbitmqUrl := "amqp://guest:guest@" +
		viper.GetString("RABBITMQ_HOST") + viper.GetString("RABBITMQ_PORT")

	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer conn.Close()

	log.Println("Successfully connected to the RabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("RabbitMQ practice",
		"",
		true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	messages := make(chan string)
	go func() {
		for d := range msgs {
			courseId := string(d.Body[:])
			log.Println("Received msg: ", courseId)
			examService := service.NewExamService(repository)
			examService.DelCourseExams(courseId)
		}
	}()
	log.Println("Successfully consumed msg to the queue")
	<-messages

}
