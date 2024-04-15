package main

import (
	log "github.com/sirupsen/logrus"
	"service_topic/config"
	_ "service_topic/docs"
	"service_topic/internal/routes"
	"service_topic/internal/tools"
)

func main() {
	//	@title		User API
	//	@version	1.0.0

	// 	@description 	Это topic_service с использованием свагера
	// 	@termsOfService  тут были-бы условия использования, еслибы я их мог обозначить
	// 	@contact.url    https://t.me/Timuchin3
	// 	@contact.email  tima.gorenskiy@mail.ru

	// 	@securityDefinitions.apikey  ApiKeyAuth
	//  @in header
	//  @name Authorization

	//	@host localhost:8800

	config.CheckFlagEnv()
	tools.InitLogger()
	tools.InfoLogs()

	// config.InitPgSQL инициализируем подключение к базе данных
	err := config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Panic(err)
	}
	tools.INFO.WithField("cmd", "initialization").Info("подключение к базе успешно")

	r := routes.SetupRouter()
	tools.INFO.WithField("cmd", "initialization").Info("ручки управления api получены")

	// запуск сервера
	tools.INFO.WithField("cmd", "initialization").Info("запуск сервера")
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Panic(err)
	}
}
