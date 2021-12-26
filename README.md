# Agregator

<h3>Для запуска проекта локально:</h3>
1. Установить PostgreSQL, создать пустую базу данных, назвать ее
2. Cоздать файл ".env" внутри проекта и вписать в него следующее:

export DIALECT="postgres"
export HOST="localhost"
export DBPORT="5432"
export USER= *здесь имя пользователя для доступа к базе данных*
export NAME= *имя созданной базы данных*
export PASSWORD= *пароль пользователя*

3. В консоли внутри директории проекта выполнить команды: 
```console
$ source .env
$ go run main.go
```
4. Возможно, потребуются команды (но вряд ли, раз есть файл go.mod):

```console
$ go mod init
$ go mod tidy
```

```console
$ go get "gorm.io/driver/postgres"
$	go get "gorm.io/gorm"
```
