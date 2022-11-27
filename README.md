# Simple Web Site

Create file .env to store variable: <b>MYSQL_PASSWORD={my_password}</b>

where <b>{my_password}</b> - is your password.

Set host in configs/config.yaml

***
Befort starting application start docker container with command <b>(change {my_password})</b>:
```shell
docker run \
--name simple_website_db \
-e MYSQL_ROOT_PASSWORD={my_password} \
-e MYSQL_DATABASE=simple_website \
-p 3307:3306 \
-d --rm mysql:latest
```

To make migrations use https://github.com/golang-migrate/migrate
```shell
migrate create -ext sql -dir ./schema -seq init # create migration files
```

To make migrations <b>(change {my_password})</b>:
```shell
migrate -path ./db_schema -database 'mysql://root:{my_password}@tcp(127.0.0.1:3307)/simple_website?query' up # migrate to database
migrate -path ./db_schema -database 'mysql://root:{my_password}@tcp(127.0.0.1:3307)/simple_website?query' down # migrate to database
```
