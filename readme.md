# go gin project template

## Packages

- GO Gin (gin-gonic.com)
- Gorm (gorm.io)
- go-playground/validator (github.com/go-playground/validator)
- godotenv (github.com/joho/godotenv)
- Kafka (github.com/Shopify/sarama)

## Project stucture
```bash
├── config
│   ├── database.config.go
|   ├── env.config.go
|   ├── kafka.config.go
|   └── ...
├── constant 
|   ├── messages.go
|   └── ...
├── controllers
|   └── ex: user.controller.go
├── dto (data transfer object)
|   ├── request
|   └── response
├── entity
├── helper
├── repository
├── routes
├── service
├── .env
├── .env.example
├── go.mod
├── go.sum
├── .gitignore
└── readme.md
```