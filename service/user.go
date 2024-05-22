package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"tests/lib"
	"tests/model"
)

var UserService = User{}

type User struct{}

type LoginInfo struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func (*User) Login(c *gin.Context) {
	input := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := lib.VerifyMulti(input, lib.UserVerify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	password, err := lib.Decode(input.Password, lib.KeyEncode)
	if err != nil || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid username or password"})
		return
	}
	if ok, err := model.UserModel.Login(input.Username, password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	} else if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid username or password"})
		return
	}

	jwt := lib.NewJWT()
	token, err := jwt.CreateToken(jwt.CreateClaims(lib.BaseClaims{
		Uid: model.UserInfo.Id,
	}))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, LoginInfo{
		Id:    model.UserInfo.Id,
		Token: token,
	})
}

func (*User) Info(c *gin.Context) {
	claims, _ := c.Get(lib.KeyXClaims)
	claim := claims.(*lib.CustomClaims)

	user, err := model.UserModel.Info(claim.BaseClaims.Uid)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
