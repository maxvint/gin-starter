# Gin Strter

[Gin](https://github.com/gin-gonic/gin) Framework with a structured starter project including [Protobuf](https://github.com/golang/protobuf), CURD operations, routing, pagination and more.

### Configured with
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/jinzhu/gorm)
- [Protobuf](https://github.com/golang/protobuf)

### Features

### Quickstart

##### Install Dependencies
```
$ dep ensure
```

##### Run Your Application
```bash
$ go run ./

# You can use realize to hot reload
$ realize start --run --nc 
```

### Config
##### Config MySQL
You will find the MySQL config file in `conf/app.yaml`. In which you can setup your database config information.
```yaml
db_address:
  database_name: test
  hostname: 127.0.0.1
  username: root
  password: 123456
  port: 3306
  options: charset=utf8mb4&parseTime=True&loc=UTC&interpolateParams=true
  ssl_client_key:
  ssl_client_cert:
  ssl_ca:
  max_open_conns: 100
  max_idle_conns: 10
  max_life_time: 1000
```

### Support & Suggestions

