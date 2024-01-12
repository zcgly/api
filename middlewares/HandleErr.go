package middlewares

import (
	"github.com/kataras/iris/v12/context"
	"github.com/zcgly/api/models"
)

func HandleErr(ctx *context.Context) {
	defer func() {
		reason := recover()
		if reason != nil {
			ctx.Application().Logger().Errorf("%+v", reason)
			var resp *models.ApiResponse
			switch reason.(type) {
			case string:
				resp = models.NewMsgResponse(reason.(string))
			case *models.TitledError:
				err := reason.(*models.TitledError)
				resp = models.NewMsgResponse(err.Error())
				resp.Title = err.Title()
			case error:
				resp = models.NewMsgResponse(reason.(error).Error())
			}
			err := ctx.JSON(resp)
			if err != nil {
				ctx.Application().Logger().Error(err.Error())
			}
		}
	}()
	ctx.Next()
}
