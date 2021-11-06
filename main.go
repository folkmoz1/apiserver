package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
	"awesomeProject/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

func setUpRoutes(app *fiber.App)  {
	api := app.Group("/api")
	routes.MessagesRoute(api.Group("/messages"))
}


func main() {
	app := fiber.New()
	key := encryptcookie.GenerateKey()

	app.Static("/uploads/image/", "./public/images")

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: key,
	}))
	//app.Use(csrf.New())
	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()
	
	setUpRoutes(app)

	app.Post("/api/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}

		uniqueId, err := utils.GenRandomBytes(12)
		if err != nil {
			panic(err)
		}
		fileExt := strings.Split(file.Filename, ".")[1]
		imageName := fmt.Sprintf("%x.jpg", uniqueId)

		src, _ := file.Open()
		imgCus, err := utils.Loader(src, fileExt)
		log.Println(fileExt)
		if err != nil {
			return err
		}

		m := resize.Resize(1000, 0, imgCus, resize.Lanczos3)
		out, errCreate := os.Create(fmt.Sprintf("./public/images/%s", imageName))
		if errCreate != nil {
			return errCreate
		}
		defer out.Close()

		errEncode := jpeg.Encode(out, m, nil)
		if errEncode != nil {
			return errEncode
		}

		return c.Status(201).JSON(fiber.Map{
			"success": true,
			"url": fmt.Sprintf("http://localhost:8080/uploads/image/%s", imageName),
		})
	})


	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Error app failed to start")
	}
}