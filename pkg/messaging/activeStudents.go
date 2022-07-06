package messaging

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/dataContainers"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"strings"
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
			message := strings.Split(string(d.Body[:]), " ")
			studentId := message[1]
			fmt.Println(studentId)
			if message[0] == "open" {
				dataContainers.ActiveStudents = append(dataContainers.ActiveStudents, studentId)
			} else if message[0] == "cheat" {
				dataContainers.CheatStudents = append(dataContainers.CheatStudents, studentId)
			} else if message[0] == "close" {
				ind := -1
				for i := 0; i < len(dataContainers.ActiveStudents); i++ {
					if dataContainers.ActiveStudents[i] == studentId {
						ind = i
						break
					}
				}
				if ind != -1 {
					dataContainers.ActiveStudents = append(dataContainers.ActiveStudents[:ind], dataContainers.ActiveStudents[ind+1:]...)
				}
			}
			fmt.Println(dataContainers.ActiveStudents)
			log.Println("Received msg: ", message)
		}
	}()
	log.Println("Successfully consumed msg from desktop-student queue")
	<-messages

}
