# Simple Web Site

Create file .env to store variable: <b>MYSQL_PASSWORD={my_password}</b>

where <b>{my_password}</b> - is your password.

Set host in configs/config.yaml

***
To make migrations use https://github.com/golang-migrate/migrate
```shell
migrate create -ext sql -dir ./schema -seq init # create migration files
```
***
Before starting application <b>start docker container</b> with shell command:
```shell
make db
```

To <b>make migrations</b> use shell command:
```shell
make migrations # migrate to database
```

To <b>build executable file</b> use shell command:
```shell
make build
```

To <b>stop and remove docker-container</b> use shell command:
```shell
make stop
```

To <b>start application in terminal</b> use shell command:
```shell
make # or make start
```
