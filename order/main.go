package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"order/application"
	"order/domain"
	"order/infrastructure"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type CreateOrderBody struct {
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
}

func main() {
	ctx := context.Background()

	orderRepository := infrastructure.OrderRepositroy{}

	redisHost := os.Getenv("REDIS_HOST")
	if len(redisHost) == 0 {
		redisHost = "localhost"
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", redisHost),
	})

	app := fiber.New()

	app.Use(logger.New())

	go func() {
		for msg := range redisClient.PSubscribe(ctx, "credit:*").Channel() {
			ev := map[string]interface{}{}
			if err := json.Unmarshal([]byte(msg.Payload), &ev); err != nil {
				log.Println(err)
			}

			switch msg.Channel {
			case "credit:limit_exceeded":
				id := ev["order_id"].(string)

				application.RejectOrder(id, orderRepository)
			case "credit:reserved":
				id := ev["order_id"].(string)

				application.ApproveOrder(id, orderRepository)
			}

		}
	}()

	app.Post("/orders", func(c *fiber.Ctx) error {

		body := new(CreateOrderBody)
		if err := c.BodyParser(body); err != nil {
			return fiber.ErrBadRequest
		}

		order, err := application.CreateOrder(body.CustomerID, orderRepository)
		if err != nil {
			log.Panicln(err)
		}

		msg, _ := json.Marshal(&domain.OrderCreated{Id: order.Id, Amount: order.Amount, CustomerID: order.CustomerID})
		redisClient.Publish(ctx, "order:created", msg)

		return c.Status(http.StatusCreated).JSON(order)
	})

	app.Get("/orders", func(c *fiber.Ctx) error {
		orders, _ := orderRepository.FindMany(nil)

		return c.JSON(&fiber.Map{"data": orders})
	})

	log.Fatal(app.Listen(":3000"))
}
