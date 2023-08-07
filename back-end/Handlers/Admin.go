package Handlers

import(
	"github.com/gofiber/fiber/v2"
	"chat/Database"
)

func DeleteRoom(c *fiber.Ctx) error {
	UserToken := c.Cookies("Token")
	RoomToken := c.Params("Token")
	User,err := Database.FindUserByToken(UserToken)
	if err != nil {
		return c.Status(404).SendString("Kullanıcı bulunamadı")
	}
	
	Room,err := Database.FindRoom(RoomToken)
	if err != nil {
		return c.Status(404).SendString("Oda bulunamadı")
	}

	if User.Perm == "Admin" || User.ID == Room.OwnerID {
		err := Database.DeleteRoom(RoomToken)
		if err != nil {
			return c.Status(500).SendString("Bir sorun oluştu")
		}
		return c.Status(200).SendString("Oda başarıyla silindi")
	}
	return nil
}