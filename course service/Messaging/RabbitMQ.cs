﻿using System.Text;
using RabbitMQ.Client;

namespace Messaging;

public class RabbitMq
{
    public static void Send(string msg)
    {
        var factory = new ConnectionFactory() { HostName = "localhost" };
        
        using (var connection = factory.CreateConnection())
            
        using (var channel = connection.CreateModel())
        {
            channel.QueueDeclare(queue: "course-exam", durable: false,
                exclusive: false, autoDelete: false, arguments: null);

            var messageBody = Encoding.UTF8.GetBytes(msg);
            
            channel.BasicPublish(exchange: "", routingKey: "course-exam",
                basicProperties: null, body: messageBody);
        }
    }
            
}