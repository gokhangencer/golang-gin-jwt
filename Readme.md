$ go env GO111MODULE

# If it is not set to "auto", then run:

$ go env -w GO111MODULE=auto

$ cd wsgin

$ go mod init wsgin

# force Go to behave the GOPATH way, even outside of GOPATH

$ go env -w GO111MODULE=off

# install packages

$ go get . # fail so moved project to $GOPATH\src
$ go get github.com/golang-jwt/jwt

// gin framework
go get -u github.com/gin-gonic/gin
// ORM library
go get -u github.com/jinzhu/gorm
// package that we will be used to authenticate and generate our JWT
go get -u github.com/dgrijalva/jwt-go
// to help manage our environment variables
go get -u github.com/joho/godotenv
// to encrypt our user's password
go get -u golang.org/x/crypto

# run service

go run .

f5 => installs go debugger

# ==============

# go playground

# https://go.dev/play/

# thanks to

# https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817

Golang
Jwt
Rest Api
Go
Gin
