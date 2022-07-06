package messaging

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/dataContainers"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
)

func AddCheatStudent() {
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

	msgs, err := ch.Consume("cheat-student",
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
			log.Println("Received cheat student id: ", studentId)
			dataContainers.CheatStudents = append(dataContainers.CheatStudents, studentId)
		}
	}()
	log.Println("Successfully consumed msg from cheat-student queue")
	<-messages

}
