package main

import (
	"fiber-api-example/app/api"
	"fiber-api-example/app/config"
	"fiber-api-example/app/server"
	"fiber-api-example/app/server/middleware"
)

func main() {
	config.Init()
	app := server.Create()
	middleware.RegisterMiddlewares(app)
	api.SetupRoutes(app)
	api.SwaggerRoute(app)
	server.StartServerWithGracefulShutdown(app)
}
