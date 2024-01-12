package api

import (
	"encoding/json"
	"github.com/kataras/iris/v12/context"
	"github.com/zcgly/api/models"
	"golang.org/x/exp/maps"
)

func WriteString(ctx *context.Context, data string) {
	_, err := ctx.WriteString(data)
	checkErr(err)
}

func WriteData(ctx *context.Context, data any) {
	resp := &models.ApiResponse{Success: true, Code: 200, Data: data}
	checkErr(ctx.JSON(resp))
}

func WriteOK(ctx *context.Context) {
	resp := &models.ApiResponse{Success: true, Code: 200}
	checkErr(ctx.JSON(resp))
}

func WriteFailed(ctx *context.Context, data any) {
	resp := &models.ApiResponse{Success: false, Code: 400, Data: data}
	checkErr(ctx.JSON(resp))
}

func ReadJSON(ctx *context.Context, outPtr any) {
	checkErr(ctx.ReadJSON(outPtr))
}

func GetJsonFields(ctx *context.Context, key string) []string {
	bs, err := ctx.GetBody()
	checkErr(err)

	m := make(map[string]any)
	err = json.Unmarshal(bs, &m)
	checkErr(err)

	if key != "" {
		mm := m[key].(map[string]any)
		return maps.Keys(mm)
	}
	return maps.Keys(m)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
