package middleware

import (
	"awesomeProject3/utils"
	"awesomeProject3/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"email"`
	//Password string `json:"password"`
	jwt.StandardClaims
}
var code int
//生成token
func SetToken (username string) (string,int) {
	expireTime := time.Now().Add(10 *time.Hour)
	SetClaims := MyClaims{
		username,
		///password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer :"showdoc",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token , err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "1",errmsg.ERROR
	}
	return token , errmsg.SUCCESE
}
//验证token
func CheckToken(token string)(*MyClaims,int){
	setToken, err := jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil

	})
	if err != nil{
		if ve,ok := err.(*jwt.ValidationError);ok{
			if ve.Errors&jwt.ValidationErrorMalformed != 0{
				return nil,errmsg.ERROR_TOKEN_WRONG
			}else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) !=0{
				return nil,errmsg.ERROR_TOKEN_RUNTIME
			}else{
				return nil,errmsg.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil{
		if key := setToken.Claims.(*MyClaims); setToken.Valid{
			return key,errmsg.SUCCESE
		}else{
			return nil, errmsg.ERROR_TOKEN_WRONG
		}
	}
	return nil,errmsg.ERROR_TOKEN_WRONG
}


//jwt 中间件
func JwtToken() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		//cookie,e := c.Request.Cookie("Cookie")

		if tokenHerder =="" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder," ",2)
		if len(checkToken) == 0{
			code =errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"status" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if len(checkToken)!=2 && checkToken[0]!="Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key ,tCode := CheckToken(checkToken[1])
		if tCode ==errmsg.ERROR{
			code =tCode
			c.JSON(http.StatusOK,gin.H{
				"code": code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//if e == nil{
		//	c.SetCookie(cookie.Name,cookie.Value,1000,cookie.Path,cookie.Domain,cookie.Secure,cookie.HttpOnly)
		//	c.Next()
		//} else {
		//	c.Abort()
		//}
		c.Set("username",key.Username)
		c.Next()
		//fmt.Println(errmsg.ERROR)
	}
}
