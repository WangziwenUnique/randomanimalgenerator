package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/aogen-fiber/backend/config"
	"github.com/aogen-fiber/backend/facade"
	"github.com/aogen-fiber/backend/infrastructure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Parse command line flags
	prod := flag.Bool("prod", false, "enable prod mode")
	port := flag.Int("port", 3000, "server port")
	flag.Parse()

	// Set production mode
	config.SetProd(*prod)

	// 确保在程序退出时关闭所有数据库连接
	defer infrastructure.GetDBManager().CloseAll()

	// Only log in master process when using prefork
	if !config.IsProd() || !fiber.IsChild() {
		log.Printf("Starting server in %s mode", map[bool]string{true: "production", false: "development"}[config.IsProd()])
		log.Printf("Server will listen on port %d", *port)
	}

	app := fiber.New(fiber.Config{
		ServerHeader:            "Go-Fiber-App",
		AppName:                 "Go-Fiber-App",
		EnableTrustedProxyCheck: true,
		ProxyHeader:             "X-Forwarded-For",
		Prefork:                 config.IsProd(),
		Concurrency:             runtime.NumCPU() * 1024,
		DisableStartupMessage:   config.IsProd() && fiber.IsChild(),
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders: "Content-Disposition",
		MaxAge:        300,
	}))

	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Register handlers
	api := app.Group("/api")
	facade.RegisterHandlers(api)

	// Static file service
	app.Static("/", "./dist")
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/index.html")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
