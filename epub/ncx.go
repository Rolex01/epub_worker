package epub

import (
	"io"
)

type xmlNCX struct {
	NavMap []navpoint `xml:"navMap>navPoint"`
}
type navpoint struct {
	Text     string     `xml:"navLabel>text"`
	Content  content    `xml:"content"`
	NavPoint []navpoint `xml:"navPoint"`
}
type content struct {
	Src string `xml:"src,attr"`
}

func parseNCX(ncx io.Reader) (*xmlNCX, error) {
	var n xmlNCX
	err := decodeXML(ncx, &n)
	if err != nil {
		return nil, err
	}

	return &n, nil
}

func (ncx xmlNCX) navMap() []navpoint {
	return ncx.NavMap
}

func (point navpoint) Title() string {
	return point.Text
}

func (point navpoint) URL() string {
	return point.Content.Src
}

func (point navpoint) Children() []navpoint {
	return point.NavPoint
}
