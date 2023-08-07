package Routers

import(
	"chat/Middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"chat/Handlers"
)

func Initalize(app *fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", Handlers.HelloWorld)

	app.Get("/api/rooms", Handlers.RoomsApi)
	app.Get("/api/room/:Token", Handlers.RoomApi)
	app.Get("/api/user/:Token", Handlers.UserApi)

	app.Post("/login", Handlers.Login)
	app.Post("/register", Handlers.Register)
	app.Post("/create/room", Handlers.CreateRoom)

	app.Delete("/delete/room/:Token", Handlers.DeleteRoom)

	app.Use(Middleware.Security)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}