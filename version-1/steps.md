# STEPS

Go Full-Stack Pizza Tracker Admin Dashboard | Real-Time Updates (Gin, GORM, SSE, Golang)
[pixxa-track](https://www.youtube.com/watch?v=8XRTAPWMO2E)

go mod init github.com/carloscfgos1980/pizza-tracker
mkdir -p internal/models templates/static data cmd

go get gorm.io/gorm
go get github.com/teris-io/shortid
go get gorm.io/driver/sqlite
go get github.com/gin-gonic/gin
go get github.com/go-playground/validator/v10
