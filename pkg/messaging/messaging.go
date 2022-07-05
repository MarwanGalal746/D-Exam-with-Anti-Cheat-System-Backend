package messaging

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
)

func AddActiveStudent() {
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

	msgs, err := ch.Consume("desktop-student",
		"",
		true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	messages := make(chan string)
	go func() {
		for d := range msgs {
			studentId := string(d.Body[:])
			log.Println("Received msg: ", studentId)
			pkg.ActiveStudents = append(pkg.ActiveStudents, studentId)
		}
	}()
	log.Println("Successfully consumed msg to the queue")
	<-messages

}
