package middlewares

import (
	"github.com/kataras/iris/v12/context"
	"strings"
)

func Token(ctx *context.Context) {
	token := ctx.GetHeader("Authorization")
	fields := strings.Fields(token)
	if len(fields) >= 2 {
		token = fields[1]
		ctx.Values().Set("token", token)
	}
	ctx.Next()
}
