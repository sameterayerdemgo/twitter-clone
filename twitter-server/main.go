package main

import (
	"cekirdek/controller"
	"cekirdek/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var logger *zap.SugaredLogger

func main() {
	// create zaplogger instance
	createLogger()

	if err := readConfigFile(); err != nil {
		logger.Fatal("unmarshalling yaml file ", err)
	}

	// migration process for creating table and sample data
	/*	if err := createDbMigration(); err != nil {
			logger.Fatal("migration error", err)
	} */

	router := gin.Default()
	gin.SetMode(model.Configurationins.Server.Mode)
	// postgroup := router.Group("/post")
	usergroup := router.Group("/users")

	usergroup.POST("/login", controller.Login)

	if err := router.Run(model.Configurationins.Server.Port); err != nil {
		logger.Fatal("server problem \t ", err)
	}
}

func readConfigFile() error {

	file, err := ioutil.ReadFile("resources/application.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &model.Configurationins)
	fmt.Print(model.Configurationins)
	return err
}

func createLogger() {
	zaploger, _ := zap.NewProduction()
	defer zaploger.Sync() // flushes buffer, if any
	logger = zaploger.Sugar()
}

func createDbMigration() error {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}
