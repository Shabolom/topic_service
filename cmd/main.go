package main

import (
	log "github.com/sirupsen/logrus"
	"service_topic/config"
	"service_topic/internal/routes"
	"service_topic/internal/tools"
)

func main() {
	config.CheckFlagEnv()
	tools.InitLogger()

	// config.InitPgSQL инициализируем подключение к базе данных
	err := config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Panic(err)
	}

	// конфигурация (инициализация) end point или ручка (можно назвать имя запроса)
	// (как api student) URLов пример (localhost, 8080) конфигурация всех URLов которые будет
	// обрабатывать сервер
	r := routes.SetupRouter()

	// запуск сервера
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Panic(err)
	}
}
