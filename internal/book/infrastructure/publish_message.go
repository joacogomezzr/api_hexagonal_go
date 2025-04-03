package infrastructure

import (
    "api-joaquin/internal/book/domain"
    "context"
    "encoding/json"
    "log"

    "time"

    amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func SendMessageToRabbit(payload *domain.Book) {
    conn, err := amqp.Dial("amqp://joaquin:123456@98.82.195.253:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    err = ch.ExchangeDeclare(
        "exchange",   // name
        "direct", // type
        true,     // durable
        false,    // auto-deleted
        false,    // internal
        false,    // no-wait
        nil,      // arguments
    )
    failOnError(err, "Failed to declare an exchange")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    body, err := json.Marshal(payload)
    failOnError(err, "Failed to marshal payload to JSON")

    err = ch.PublishWithContext(ctx,
        "exchange", // exchange
        "",         // routing key
        false,      // mandatory
        false,      // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
    failOnError(err, "Failed to publish a message")

    log.Printf(" [x] Sent %s", body)
}
