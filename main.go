package main

import (
	"MPPLProject/authentication/delivery/http"
	. "MPPLProject/authentication/models"
	"MPPLProject/authentication/repository"
	"MPPLProject/authentication/usecase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"net/url"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbUser := viper.GetString(`database.user`)
	dbPassword := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
	}

	defer dbConn.Close()

	dbConn.DropTableIfExists(&User{}, &Role{}, &Account{}, &UserRole{})

	dbConn.AutoMigrate(&User{}, &Role{}, &Account{}, &UserRole{}, &UserStatus{}, &StatusCustomer{})

	ur := repository.NewUserRepository(dbConn)

	uu := usecase.NewUserUseCase(ur)

	e := echo.New()
	http.NewUserHandler(e, uu)

	_ = e.Start(viper.GetString("server.address"))
}
