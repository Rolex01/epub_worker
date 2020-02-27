package fb2

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
	"encoding/xml"
)

// get xlink from enclosed tag image
func parseImage(data []byte) string {
	result := ""
	quoteOpened := false
_loop:
	for _, v := range data {
		if quoteOpened {
			if v == '"' {
				break _loop
			}
			result += string(v)
		} else {
			if v == '"' {
				quoteOpened = true
			}
		}
	}
	return result
}

func decodeWin1251(i io.Reader) (r io.Reader) {
	decoder := charmap.Windows1251.NewDecoder()
	r = decoder.Reader(i)

	return
}

func (p *Parser) PrintXML() {
	var f interface{}

	fmt.Println(p.reader)

	err := xml.Unmarshal(p.book, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Begin printing...", f)

	// test write filename
	file, err := os.OpenFile("parse.json", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	printXML(f, file)
}

func printXML(v interface{}, file *os.File) {
	switch vv := v.(type) {
	case string:
		if _, err := file.WriteString(fmt.Sprint("is string", vv)); err != nil {
			panic(err)
		}
	case float64:
		if _, err := file.WriteString(fmt.Sprint("is float64", vv)); err != nil {
			panic(err)
		}
	case []interface{}:
		if _, err := file.WriteString("is an array:"); err != nil {
			panic(err)
		}
		for i, u := range vv {
			if _, err := file.WriteString(fmt.Sprint(i, " ")); err != nil {
				panic(err)
			}
			printXML(u, file)
		}
	case map[string]interface{}:
		if _, err := file.WriteString("is an object:"); err != nil {
			panic(err)
		}
		for i, u := range vv {
			if _, err := file.WriteString(fmt.Sprint(i, " ")); err != nil {
				panic(err)
			}
			printXML(u, file)
		}
	default:
		if _, err := file.WriteString("Unknown type"); err != nil {
			panic(fmt.Sprint("qwe", err))
		}
	}
}
