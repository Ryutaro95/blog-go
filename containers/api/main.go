package main

import "github.com/Ryutaro95/app/presentation/router"

func main() {
	// DB settings
	// if err := config.InitEnviroment(); err != nil {
	// 	logger.Fatal(err.Error())
	// }

	e := router.Router()
	e.Logger.Fatal(e.Start(":8080"))
}
