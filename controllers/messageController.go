package controllers

import (
	"context"
	"log"
	"time"

	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type Username  string

func AddNewMessage(c *fiber.Ctx) error {
	userCollection := config.MI.DB.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	data := new(models.Message)

	if err := c.BodyParser(&data); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error": err,
		})
	}

	username := c.Query("username")
	if username == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Missing username query parameter",
		})
	}

	bangkok, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(bangkok)
	data.Timestamp = now

	filter := bson.M{"username": username}
	update := bson.M{
		"$push": bson.M{"messages": data},
	}
	err := userCollection.FindOneAndUpdate(ctx, filter, update)
	if err.Err() != nil {
		log.Println(err.Err())
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Message failed to send",
			"error":   err.Err(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "message sent successfully",
	})
}