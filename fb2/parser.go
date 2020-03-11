package fb2

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/rolex01/epub_workers/utils"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
	"os/exec"
)

// Parser struct
type Parser struct {
	book   []byte
	reader io.Reader
}

// New creates new Parser
func New(data []byte) *Parser {
	return &Parser{
		book: data,
	}
}

// NewReader creates new Parser from reader
func NewReader(data io.Reader) *Parser {
	return &Parser{
		reader: data,
	}
}

// CharsetReader required for change encodings
func (p *Parser) CharsetReader(c string, i io.Reader) (r io.Reader, e error) {
	switch c {
	case "windows-1251":
		r = decodeWin1251(i)
	}
	return
}

// Unmarshal parse data to FB2 type
func (p *Parser) Unmarshal() (result FB2, err error) {
	if p.reader != nil {
		decoder := xml.NewDecoder(p.reader)
		decoder.CharsetReader = p.CharsetReader
		if err = decoder.Decode(&result); err != nil {
			return
		}

		result.UnmarshalCoverpage(p.book)

		return
	}
	reader := bytes.NewReader(p.book)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = p.CharsetReader
	if err = decoder.Decode(&result); err != nil {
		return
	}

	result.UnmarshalCoverpage(p.book)

	return
}

func FB2Parse(ctx *fasthttp.RequestCtx) {
	var res []byte
	textError := ""
	filename, err := utils.GetFilename(ctx)

	defer func(filename string) {
		if err := os.Remove(filename); err != nil {
			log.Fatal("ERROR remove fb2 file:", err)
		}
	}(filename)

	html, err := xml2html(filename)
	if err != nil {
		textError = "ERROR parse fb2 file:"
		fmt.Println(textError, err)
		res, _ = json.Marshal(map[string]string {"error": textError})
	} else {
		res, _ = json.Marshal(map[string]string {"status": "OK"})
	}

	fmt.Println(string(html)) // HTML Save

	ctx.SetContentType("application/json")
	ctx.SetBody(res)
}

func xml2html(xmlPath string) ([]byte, error) {
	cmd := exec.Cmd{
		Args: []string{"xsltproc", "fb2/style/stylesheet.xsl", xmlPath},
		Env:  os.Environ(),
		Path: "/usr/bin/xsltproc",
	}

	htmlOut, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return []byte(htmlOut), nil
}
