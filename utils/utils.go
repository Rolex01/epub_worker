package utils

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"math/rand"
)

func GetFilename(ctx *fasthttp.RequestCtx) (filename string, err error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}

	filename = fmt.Sprintf("tmp_%d.tmp", rand.Intn(1000))
	err = fasthttp.SaveMultipartFile(file, filename)

	return
}