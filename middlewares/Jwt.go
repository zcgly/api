package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

const JwtKey = "ipingpang"

var j *jwt.Middleware

func init() {
	j = jwt.New(jwt.Config{
		Extractor: jwt.FromAuthHeader,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func Jwt(ctx iris.Context) {
	j.Serve(ctx)
	ctx.Next()
}
