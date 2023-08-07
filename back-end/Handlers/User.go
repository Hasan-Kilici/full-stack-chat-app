package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"chat/Database"
	"chat/Utils"
)

func HelloWorld(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"hello":"world",
	})
	return nil
}

func Register(c *fiber.Ctx) error {
	form := new(RegisterForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	if !Utils.ValidatePassword(form.Password) {
		return c.Status(400).SendString("Şifre yeterince güvenli değil")
	}

	if !Utils.ValidateUsername(form.Username) {
		return c.Status(400).SendString("Kullanıcı Adı Boşluk içermemeli")
	}

	if !Database.Register(form.Username, form.Password) {
		return c.Status(400).SendString("Bu Kullanıcı adı kullanılıyor.")
	}

	login, _ := Database.Login(form.Username,form.Password)

	c.JSON(fiber.Map{
		"Token": login,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	form := new(LoginForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	user := Database.FindUserByName(form.Username)
	if !user {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}
	
	login, err := Database.Login(form.Username,form.Password)
	if err != nil{
		return c.Status(400).SendString("Kullanıcı adı yada şifre yanlış")
	}
	
	c.JSON(fiber.Map{
		"Token": login,
	})
	return nil
}

func CreateRoom(c *fiber.Ctx) error {
	form := new(CreateRoomForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	Token := c.Cookies("Token")
	User,err := Database.FindUserByToken(Token)
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}

	Room,err := Database.CreateRoom(form.Name, User.ID)
	if err != nil {
		return c.Status(500).SendString("Bir sorun oluştu")
	}

	c.JSON(fiber.Map{
		"token":Room,
	})
	return nil
}