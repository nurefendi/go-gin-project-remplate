# go gin project template

## Packages

- GO Gin (gin-gonic.com)
- Gorm (gorm.io)
- go-playground/validator (github.com/go-playground/validator)
- godotenv (github.com/joho/godotenv)
- Kafka (github.com/Shopify/sarama)

## Database driver
- MySQL

## Project stucture
```bash
├── src
│   ├── config
│   │   ├── database.config.go
│   |   ├── env.config.go
│   |   ├── kafka.config.go
│   |   └── ...
│   ├── constant 
│   |   ├── messages.go
│   |   └── ...
│   ├── controllers
│   |   └── ex: user.controller.go
│   ├── dto (data transfer object)
│   |   ├── request
│   |   └── response
│   ├── entity
│   ├── helper
│   ├── repository
│   ├── routes
│   ├── service
├── .env
├── .env.example
├── go.mod
├── go.sum
├── .gitignore
└── readme.md
```
## Run project
```bash
go run main.go
```

## Rename project

- Remove `go.mod` 
- Remove `go.sum`
- Run `go mod init <name project>`
- fix all import
- done

