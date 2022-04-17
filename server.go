package main

import (
	"github.com/azura-labs/middlewares"

	"github.com/azura-labs/databases"

	cm "github.com/azura-labs/common"

	controller "github.com/azura-labs/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	cm.InitConfig()

	db, err := databases.ConnectDb()
	if err == nil {
		defer db.Close()
		log.Info("MySQL Database Connected..")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middlewares.AllowOriginSkipper,
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	// USERS
	e.POST(cm.Config.RootURL+"/users", controller.PostUser)
	e.PUT(cm.Config.RootURL+"/users/:userId", controller.PutUser)
	e.DELETE(cm.Config.RootURL+"/users/:userId", controller.DeleteUser)
	e.GET(cm.Config.RootURL+"/users/:userId", controller.DetailUser)
	e.GET(cm.Config.RootURL+"/users", controller.ListUser)
	// TICKETS
	e.POST(cm.Config.RootURL+"/tickets", controller.PostTicket)
	e.PUT(cm.Config.RootURL+"/tickets/:ticketId", controller.PutTicket)
	e.DELETE(cm.Config.RootURL+"/tickets/:ticketId", controller.DeleteTicket)
	e.GET(cm.Config.RootURL+"/tickets/:ticketId", controller.DetailTicket)
	e.GET(cm.Config.RootURL+"/tickets", controller.ListTicket)
	// PAYMENTS
	e.POST(cm.Config.RootURL+"/payments", controller.PostPayment)
	e.GET(cm.Config.RootURL+"/payments/:paymentId", controller.DetailPayment)
	e.GET(cm.Config.RootURL+"/payments", controller.ListPayment)

	e.Use(middleware.RequestID())

	// Start serverlog.Info()
	log.Info("Staring server ...")

	e.Logger.Fatal(e.Start(":" + cm.Config.Port))
}
