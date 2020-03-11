package epub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rolex01/epub_workers/utils"
	"github.com/valyala/fasthttp"
	"log"
)

type EpubParams struct {
	DateRaw			int64		`json:"date"`
	Parent			int64		`json:"parent"`
	RequestedStats	[]string	`json:"requested_stats"`
}

func EpubParse(ctx *fasthttp.RequestCtx) {
	var res []byte
	var err error
	textError := ""
	filename, err := utils.GetFilename(ctx)



	e, err := Open(filename)
	if err != nil {
		textError := fmt.Sprintln("ERROR Open file:", err.Error())
		log.Println(textError)
	}
	defer e.Close()

	styles := e.GetStyle()

	// CSS Save
	for i, val := range styles {
		fmt.Println("Styles:", i, string(val))
	}

	// HTTML Save
	q, _ := e.Spine()
	for true {
		fmt.Println(q.IsFirst(), q.IsLast(), q.URL())
		r, err := q.Open()
		defer r.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		fmt.Println("File:", buf.String())

		if q.IsLast() {
			break
		}
		err = q.Next()
		if err != nil {
			nowError := fmt.Sprintln("ERROR next Spine: ", err.Error())
			log.Fatal(nowError)
			textError += nowError
		}
	}

	if textError != "" {
		res, _ = json.Marshal(map[string]string {"error": textError})
	} else {
		res, _ = json.Marshal(map[string]string {"status": "OK"})
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(res)
}
