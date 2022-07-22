package main

import (
	"fiber-api-example/app/api"
	"fiber-api-example/app/config"
	"fiber-api-example/app/server"
	"fiber-api-example/app/server/middleware"
	_ "fiber-api-example/docs"
)

func main() {
	config.Init()
	app := server.Create()
	middleware.RegisterMiddlewares(app)
	api.SwaggerRoute(app)
	api.SetupRoutes(app)
	server.StartServerWithGracefulShutdown(app)
}
