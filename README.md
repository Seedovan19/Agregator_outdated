# Агрегатор

<h2>Для запуска проекта локально:</h2>
1. Установить PostgreSQL, создать пустую базу данных, назвать ее
<br><br>
2. Cоздать файл ".env" в корневой директории проекта и вписать в него следующее:
<br>
<p>HOST="localhost"</p>
<p>DBPORT="5432"</p>
<p>USER= *здесь имя пользователя для доступа к базе данных*</p>
<p>PASSWORD= *пароль пользователя*</p>
<p>DBNAME= *имя созданной базы данных*</p>
<p>SSLMODE="disabled"</p>
<p>PORT=*номер порта для сервера, если занят по умолчанию (по умолчанию 8080)*</p>
<br>
3. В консоли внутри директории проекта выполнить команды:
<br>

```console
$ go run main.go
```
