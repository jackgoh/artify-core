package main

import (
	"github.com/NghiaTranUIT/artify-core/constant"
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/NghiaTranUIT/artify-core/resources"
	"github.com/NghiaTranUIT/artify-core/service"
	"github.com/NghiaTranUIT/artify-core/utils"
	"net/http"
)

func main() {

	// Default configuration
	config := constant.Config{
		IsEnablePostgreSQL:       true,
		IsEnablePostgreSQLLogger: true,
	}

	// Resource
	r, err := resources.Init(config)
	if err != nil {
		utils.LogError("Initialed Resource ...", err)
		return
	}
	defer r.Close()

	// Migration
	r.PostgreSQL.AutoMigrate(&model.Author{}, &model.Photo{})

	// Hello
	utils.LogInfo("Hello, world.")

	// App
	app := service.GetEngine(config)
	s := &http.Server{
		Addr: constant.AppHost + ":" + constant.AppPort,
	}

	// Start
	app.Logger.Fatal(app.StartServer(s))
}
