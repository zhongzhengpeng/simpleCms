package user

import (
	"fmt"
	"net/http"
	"simpleCms/app/common"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey string = "simpleCms"

func Login(ctx *gin.Context) {
	// 1.获取用户名和密码
	ul := &UserLogin{}
	if err := ctx.ShouldBind(ul); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	//2.查询用户
	user := &User{}
	if result := common.DB.Where("username=?", ul.Username).First(user); result.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "用户不存在",
		})
		return
	}
	fmt.Println(password(ul.Password))
	//3.核对密码
	if user.Password != password(ul.Password) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "密码错误",
		})
		return
	}
	//4.生成token
	claims := jwt.StandardClaims{
		Id:        fmt.Sprintf("%d", user.ID),
		ExpiresAt: time.Now().Add(60 * 24 * 3600 * time.Second).Unix(),
		Audience:  "SimpleSMS",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "SimpleSMS",
		NotBefore: 0,
		Subject:   "SimpleSMS-User-Auth",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "token 生成错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"token":   tokenStr,
	})
	return
}
