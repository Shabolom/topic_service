# Topic_service
___
### Подключение и настройка зависимостей :
  + **Установка `curl`  необходим для установки мираций**.
    ```go
    $ go get -u github.com/andelf/go-curl 
    ```
  + **Установка пакета миграций**.
    ```go 
     $ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
    ```
  + **Для автоматической докачки зависимостей:** используем команду.
    ```go
    $ go get ./..
    ```
  + **Создание файла миграции.**
    ```go
    $ migrate create -ext sql -dir (dir) -seq fileName
    ```
  + **Инициализация миграций.**
    ```go
    $ migrate -source file://file_path -database postgresql://dbUser:dbPassword@dbHost:dbPort/dbName?sslmode=disable
    ```
____
### Функционал Dockerfile и docker-compose.

Инициализация docker-compose файла.
```go
$ docker-compose --project-directory .\build\dev\ up --build
```
___
### Описание флагов.

```go 
// изменить хост
var flagHost = flag.String("h", "", "host")
// изменить порт
var flagPort = flag.String("p", "", "port")
// изменить хост db
var flagDbHost = flag.String("dh", "", "dbHost")
// изменить порт db
var flagDbPort = flag.String("dp", "", "dbPort")
// изменить пользователя db
var flagDbUser = flag.String("du", "", "dbUser")
// изменить пароль к db
var flagDbPassword = flag.String("dpa", "", "dbPassword")
// заменить имя подключения db
var flagDbName = flag.String("dn", "", "dbName")
// изменить изменить статус проэкта
var flagProduction = flag.Bool("pr", false, "production")
```
___
### Swagger 
>http://localhost:8800/docs/index.html#/

