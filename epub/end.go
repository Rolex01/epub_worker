package epub

type xmlEND struct {
	List struct {} `xml:"html>body>nav>ol"` // to be continued
	Title string `xml:"html>body>nav>h1"`
}
