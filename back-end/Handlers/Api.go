package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"chat/Database"
)

func RoomsApi(c *fiber.Ctx) error {
	Rooms,err := Database.ListRoom()
	if err != nil {
		c.Status(404).SendString("Oda bulunamadı")
	}
	
	c.JSON(fiber.Map{
		"data":Rooms,
	})

	return nil
}

func RoomApi(c *fiber.Ctx) error {
	Token := c.Params("Token")
	Room,err := Database.FindRoom(Token)
	if err != nil {
		return c.Status(404).SendString("Oda bulunamadı")
	}

	c.JSON(fiber.Map{
		"data":Room,
	})

	return nil
}

func UserApi(c *fiber.Ctx) error {
	Token := c.Params("Token")
	User,err := Database.FindUserByToken(Token)
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}

	c.JSON(fiber.Map{
		"data":User,
	})

	return nil
}