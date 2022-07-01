using System.Text;
using RabbitMQ.Client;

namespace Messaging;

public static class RabbitMq
{
    public static void Send(string msg)
    {
        var factory = new ConnectionFactory() { HostName = Environment.GetEnvironmentVariable("RABBITMQ_HOST_NAME", EnvironmentVariableTarget.Process) };
        
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