# Simple Web Site
---
Create file .env to store variable: <b>MYSQL_PASSWORD={my_password}</b>

where <b>{my_password}</b> - is your password.

Set host in configs/config.yaml

***

### With **make**

To <b>build docker-container</b> use shell command:
```shell
make build_app
```

To <b>stop and remove docker-container</b> use shell command:
```shell
make stop_app
```

To <b>start application</b> use shell command:
```shell
make # or make start_app
```

***

### With docker-compose

To <b>create docker-container</b> use shell command:
```shell
docker-compose build
```

To <b>start application</b> use shell command:
```shell
docker-compose up
```

To <b>stop docker-container</b> use shell command:
```shell
docker-compose down
```
