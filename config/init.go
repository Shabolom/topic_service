package config

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// env Структура для хранения переменных среды
type env struct {
	Host            string
	Port            string
	DbHost          string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbName          string
	LocalApi        string
	ConnectionApi   string
	ConnectionLogin string
	Production      bool
	SecretKey       string
}

// Env глобальная переменная для доступа к переменным среды
var Env env

// CheckFlagEnv Метод проверяющий флаги
func CheckFlagEnv() {

	var host string
	var port string
	var dbHost string
	var dbPort string
	var dbUser string
	var dbPassword string
	var dbName string
	var localApi string
	var connectionApi string
	var connectionLogin string
	var production bool
	var secretKey string

	// сканируем env файл
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	var flagHost = flag.String("h", "", "host")
	var flagPort = flag.String("p", "", "port")
	var flagDbHost = flag.String("dh", "", "dbHost")
	var flagDbPort = flag.String("dp", "", "dbPort")
	var flagDbUser = flag.String("du", "", "dbUser")
	var flagDbPassword = flag.String("dpa", "", "dbPassword")
	var flagDbName = flag.String("dn", "", "dbName")
	var flagProduction = flag.Bool("pr", false, "production")
	var flagSecretKey = flag.String("ske", "", "secret key for jwt")

	flag.Parse()

	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	} else {
		host = "localhost"
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	if os.Getenv("DB_HOST") != "" {
		dbHost = os.Getenv("DB_HOST")
	} else {
		dbHost = ""
	}

	if os.Getenv("DB_PORT") != "" {
		dbPort = os.Getenv("DB_PORT")
	} else {
		dbPort = ""
	}

	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	} else {
		dbUser = ""
	}

	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	} else {
		dbPassword = ""
	}

	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	} else {
		dbName = ""
	}

	if os.Getenv("LOCAL_API") != "" {
		localApi = os.Getenv("LOCAL_API")
	} else {
		localApi = ""
	}

	if os.Getenv("CONNECTION_API_AUT") != "" {
		connectionApi = os.Getenv("CONNECTION_API_AUT")
	} else {
		connectionApi = ""
	}

	if os.Getenv("CONNECTION_API_LOGIN") != "" {
		connectionLogin = os.Getenv("CONNECTION_API_LOGIN")
	} else {
		connectionLogin = ""
	}

	if os.Getenv("PRODUCTION") != "" {
		production, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))
	} else {
		production = false
	}

	if os.Getenv("SECRET_KEY") != "" {
		secretKey = os.Getenv("SECRET_KEY")
	} else {
		secretKey = ""
	}

	if *flagHost != "" {
		host = *flagHost
	}

	if *flagPort != "" {
		port = *flagPort
	}

	if *flagDbHost != "" {
		dbHost = *flagDbHost
	}

	if *flagDbPort != "" {
		dbPort = *flagDbPort
	}

	if *flagDbUser != "" {
		dbUser = *flagDbUser
	}

	if *flagDbPassword != "" {
		dbPassword = *flagDbPassword
	}

	if *flagDbName != "" {
		dbName = *flagDbName
	}

	if *flagProduction != false {
		production = *flagProduction
	}

	if *flagSecretKey != "" {
		secretKey = *flagSecretKey
	}

	Env = env{
		Host:            host,
		Port:            port,
		DbHost:          dbHost,
		DbPort:          dbPort,
		DbUser:          dbUser,
		DbPassword:      dbPassword,
		DbName:          dbName,
		LocalApi:        localApi,
		ConnectionApi:   connectionApi,
		ConnectionLogin: connectionLogin,
		Production:      production,
		SecretKey:       secretKey,
	}
}
