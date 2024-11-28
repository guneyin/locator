package main

import (
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guneyin/locator/config"
	"github.com/guneyin/locator/controller"
	"github.com/guneyin/locator/db"
	"github.com/guneyin/locator/mw"
	"github.com/guneyin/locator/util"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
	"time"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

type Application struct {
	Name       string
	Version    string
	Config     *config.Config
	HttpServer *fiber.App
	Controller *controller.Controller
	DB         *gorm.DB
}

func NewApplication(name string) (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDB(cfg.DBConn)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	httpServer := fiber.New(fiber.Config{
		ServerHeader:      fmt.Sprintf("%s HTTP Server", name),
		AppName:           name,
		EnablePrintRoutes: true,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return mw.Error(ctx, err)
		},
	})

	httpServer.Use(cors.New())
	httpServer.Use(recover.New())
	httpServer.Use(favicon.New())
	httpServer.Use(swagger.New(swagger.Config{
		BasePath: "/api/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))

	api := httpServer.Group("/api")
	cnt := controller.New(db, api)

	return &Application{
		Name:       name,
		Version:    util.GetVersion().Version,
		Config:     cfg,
		HttpServer: httpServer,
		Controller: cnt,
		DB:         db,
	}, nil
}

func (app *Application) Run() error {
	util.SetLastRun(time.Now())

	return app.HttpServer.Listen(fmt.Sprintf(":%d", app.Config.Port))
}

var cmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		runApp()
	},
}

func runApp() {
	app, err := NewApplication("Locator")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run())
}

// @title Locator App API Doc
// @version 1.0
// @description Simple location marker app

// @contact.name Hüseyin Güney
// @contact.url https://github.com/guneyin
// @contact.email guneyin@gmail.com

// @host localhost:8080
// @BasePath /api/
// @schemes http
func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
