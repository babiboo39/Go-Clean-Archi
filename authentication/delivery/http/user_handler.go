package http

import (
	"MPPLProject/authentication"
	. "MPPLProject/authentication/delivery/utils"
	"MPPLProject/authentication/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

var db *gorm.DB

type UserHandler struct {
	UserUseCase authentication.UseCase
}

func NewUserHandler(e *echo.Echo, userUseCase authentication.UseCase) {
	handler := &UserHandler{UserUseCase:userUseCase}

	e.GET("/users", handler.FetchUser)
	e.GET("/users/:id", handler.GetById)
	e.POST("/users/register", handler.Store)
	e.PUT("/users/:id", handler.Update)
	e.DELETE("/users/:id", handler.Delete)
}

func (uh *UserHandler) FetchUser(c echo.Context) error {
	listEl, err := uh.UserUseCase.Fetch()

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message:err.Error()})
	}

	return c.JSON(http.StatusOK, listEl)
}

func (uh *UserHandler) GetById(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id := uint(id_)

	el, err := uh.UserUseCase.GetById(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message:err.Error()})
	}

	return c.JSON(http.StatusOK, el)
}

func (uh *UserHandler) Update(c echo.Context) error {
	var user_ models.User
	err := c.Bind(&user_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uh.UserUseCase.Update(&user_)


	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message:err.Error()})
	}

	return c.JSON(http.StatusCreated, user_)
}
/* REGISTER FUNCTION */

func (uh *UserHandler) Store(c echo.Context) error {
	var user_ models.User

	user_.FirstName = c.FormValue("firstname")
	user_.LastName = c.FormValue("lastname")
	user_.Email = c.FormValue("email")
	user_.PhoneNumber = c.FormValue("phonenumber")
	user_.Password = c.FormValue("password")
	user := hashPassword(user_.Password)
	user_.Password = fmt.Sprintf("%s", user)

	//match := CheckPasswordHash(user_.Password, hash)
	//fmt.Println("Match		:", match)
	fmt.Println(user_.Password)
	err := c.Bind(&user_)

	fmt.Print(user_)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user_); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = uh.UserUseCase.Store(&user_ )

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user_)
}

/* END OF REGISTER FUNCTION */

func (uh *UserHandler) Delete(c echo.Context) error {
	id_, err := strconv.Atoi(c.Param("id"))
	id :=  uint(id_)

	err = uh.UserUseCase.Delete(id)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(u *models.User) (bool, error) {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func hashPassword(input string) string {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
