package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/internal/repository/models"
	"time"
)

var jwtKey = []byte("acking-you.xyz")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func ReleaseToken(user models.UserLogin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "douyin_pro_131",
			Subject:   "L_B__",
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}
	return nil, false
}

func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		if tokenStr == "" {
			c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 401, StatusMsg: "用户不存在"})
			c.Abort()
			return
		}
		tokenStruck, ok := ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, models.CommonResponse{
				StatusCode: 403,
				StatusMsg:  "token不正确",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, models.CommonResponse{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort()
			return
		}
		c.Set("user_id", tokenStruck.UserId)
		c.Next()
	}
}
