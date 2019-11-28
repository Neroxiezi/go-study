GO语言学习

### 环境搭建

##### 基于docker-composer的go环境搭建



- docker-composer.yml

```dockerfile
version: "0.1"
services:
  golang:
    build: ./golang
    ports:
      - "8088:8088"
    links:
      - "mysql"
      - "redis"
    volumes:
      - $HOME/www/code/go-study/code:/go
    tty: true
  mysql:
    image: mysql:5.7
    ports:
      - "3306:33066"
    volumes:
      -   /home/pfinal/www/code/go-study/data:/var/lib/mysql/
    environment:
      MYSQL_ROOT_PASSWORD: 123456
  redis:
    image: redis
    ports:
      - "6379:63791"   
```



- 目录结构:

```
|-- docker-composer.yml
|-- golang
         |--Dockerfile
```

- golang 目录下面的 Dockerfile内容:

```dockerfile
FROM golang
RUN apt-get update && apt-get install -y vim
WORKDIR $GOPATH/src
 
EXPOSE 8088
```



- 构建容器:

```
docker-compose up -d
```

> 注意: 构建完成以后,docker ps 查看一下

- 进入golang容器 

```shell
docker exec -it compose-golang_golang_1  /bin/bash
```

