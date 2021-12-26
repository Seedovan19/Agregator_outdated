# Agregator

<h3>Для запуска проекта локально:</h3>
1. Установить PostgreSQL, создать пустую базу данных, назвать ее
<br><br>
2. Cоздать файл ".env" внутри проекта и вписать в него следующее:
<br>
<p>export DIALECT="postgres"</p>
<p>export HOST="localhost"</p>
<p>export DBPORT="5432"</p>
<p>export USER= *здесь имя пользователя для доступа к базе данных*</p>
<p>export NAME= *имя созданной базы данных*</p>
<p>export PASSWORD= *пароль пользователя*</p>
<br>
3. В консоли внутри директории проекта выполнить команды:
<br>
```console
$ source .env
$ go run main.go
```
<br>
4. Возможно, потребуются команды (но вряд ли, раз есть файл go.mod):
<br>
```console
$ go mod init
$ go mod tidy
```

```console
$ go get "gorm.io/driver/postgres"
$	go get "gorm.io/gorm"
```
