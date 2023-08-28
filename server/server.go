package server

import (
	"strconv"

	"github.com/Kamalesh-Seervi/url/models"
	"github.com/Kamalesh-Seervi/url/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllRedirects(c *fiber.Ctx) error {
	urls, err := models.GetAllUrl()
	if err != nil {
		return c.Status(503).JSON(&fiber.Map{"message": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(urls)
}

func getUrl(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(503).JSON(&fiber.Map{"message": "Internal Server Error"})
	}

	url, err := models.GetOneUrl(id)
	if err != nil {
		return c.Status(503).JSON(&fiber.Map{"message": "Could Not return Url from DB"})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func createUrl(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var url models.Url
	err:=c.BodyParser(&url)
	if err != nil {
		return c.Status(503).JSON(&fiber.Map{"message": "Could Not return Url from DB"})
	}

	if url.Random{
		url.Url=utils.Base62(8)

	}
	err=models.CreateUrl(url)
	if err != nil {
		return c.Status(503).JSON(&fiber.Map{"message": "Could Not create url"})
	}
	return c.Status(fiber.StatusOK).JSON(url)

}

func Router() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	router.Get("/urls", getAllRedirects)
	router.Get("/url/:id", getUrl)
	router.Post("/url", createUrl)


	router.Listen(":3333")

}
