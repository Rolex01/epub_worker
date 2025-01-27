package fb2

import (
	"encoding/xml"
)

// List of interfaces for integration

// FB2 represents FB2 structure
//proteus:generate
type FB2 struct {
	ID          string   `bson:"_id"`
	FictionBook xml.Name `xml:"FictionBook" bson:"FictionBook"`
	Stylesheet  []string `xml:"stylesheet" bson:"stylesheet"`
	Description struct {
		TitleInfo struct {
			Genre      []string     `xml:"genre" bson:"genre"`
			GenreType  []string     `xml:"genreType" bson:"genreType"`
			Author     []AuthorType `xml:"author" bson:"author"`
			BookTitle  string       `xml:"book-title" bson:"book-title"`
			Annotation string       `xml:"annotation" bson:"annotation"`
			Keywords   string       `xml:"keywords" bson:"keywords"`
			Date struct {
				Data	string `xml:",chardata" bson:"data"`
				Value	string `xml:"value,attr" bson:"value"`
			} `xml:"date" bson:"date"`
			Coverpage  struct {
				Image struct {
					Href string `xml:"xlink:href,attr" bson:"href"`
				} `xml:"image,allowempty" bson:"image"`
			} `xml:"coverpage" bson:"coverpage"`
			Lang       string     `xml:"lang" bson:"lang"`
			SrcLang    string     `xml:"src-lang" bson:"src-lang"`
			Translator AuthorType `xml:"translator" bson:"translator"`
			Sequence   struct {
				Name 	string	`xml:"name,attr" bson:"name"`
				Number  int64	`xml:"number,attr" bson:"number"`
			} `xml:"sequence,allowempty" bson:"sequence"`
		} `xml:"title-info" bson:"title-info"`
		DocumentInfo struct {
			Author      []AuthorType `xml:"author" bson:"author"`
			ProgramUsed string       `xml:"program-used" bson:"program-used"`
			Date struct {
				Data	string `xml:",chardata" bson:"data"`
				Value	string `xml:"value,attr" bson:"value"`
			} `xml:"date" bson:"date"`
			SrcURL      []string     `xml:"src-url" bson:"src-url"`
			SrcOcr      string       `xml:"src-ocr" bson:"src-ocr"`
			ID          string       `xml:"id" bson:"id"`
			Version     float64      `xml:"version" bson:"version"`
			History     string       `xml:"history" bson:"history"`
			Publisher struct {
				FirstName  string `xml:"first-name"`
				MiddleName string `xml:"middle-name"`
				LastName   string `xml:"last-name"`
				Id		   string `xml:"id"`
			} `xml:"publisher" bson:"publisher"`
		} `xml:"document-info" bson:"document-info"`
		PublishInfo struct {
			BookName  string `xml:"book-name" bson:"book-name"`
			Publisher string `xml:"publisher" bson:"publisher"`
			City      string `xml:"city" bson:"city"`
			Year      int    `xml:"year" bson:"year"`
			ISBN      string `xml:"isbn" bson:"isbn"`
			Sequence  string `xml:"sequence" bson:"sequence"`
		} `xml:"PublishInfo" bson:"PublishInfo"`
		CustomInfo []struct {
			InfoType string `xml:"info-type" bson:"info-type"`
		} `xml:"custom-info" bson:"custom-info"`
	} `xml:"description" bson:"description"`
	Body []struct {
		Name	string `xml:"name,attr" bson:"name"`
		Image struct {
			Href string `xml:"xlink:href,attr" bson:"href"`
		} `xml:"image,allowempty" bson:"image"`
		Title		string		`xml:"title" bson:"title"`
		Epigraphs	[]string	`xml:"epigraph" bson:"epigraph"`
		Sections []struct {
			Id		string `xml:"id,attr" bson:"id"`
			Value	string `xml:",chardata" bson:"value"`
			//P []string `xml:"p" bson:"p"`
		} `xml:"section" bson:"section"`
	} `xml:"body" bson:"body"` // <body> tag needs to be processed separately
	Binary []struct {
		ID          string `xml:"id,attr" bson:"id"`
		ContentType string `xml:"content-type,attr" bson:"content-type"`
		Value       string `xml:",chardata" bson:"value"`
	} `xml:"binary" bson:"binary"`
}

// UnmarshalCoverpage func
func (f *FB2) UnmarshalCoverpage(data []byte) {
	tagOpened := false
	coverpageStartIndex := 0
	coverpageEndIndex := 0
	// imageHref := ""
	tagName := ""
_loop:
	for i, v := range data {
		if tagOpened {
			switch v {
			case '>':
				if tagName != "p" && tagName != "/p" {
				}
				tagOpened = false
				if tagName == "coverpage" {
					coverpageStartIndex = i + 1
				} else if tagName == "/coverpage" {
					coverpageEndIndex = i - 11
					break _loop
				}
				tagName = ""
				break
			default:
				tagName += string(v)
			}
		} else {
			if v == '<' {
				tagOpened = true
			}
		}
	}

	if coverpageEndIndex > coverpageStartIndex {
		href := parseImage(data[coverpageStartIndex:coverpageEndIndex])
		f.Description.TitleInfo.Coverpage.Image.Href = href
	}
}

// AuthorType embedded fb2 type, represents author info
type AuthorType struct {
	FirstName  string `xml:"first-name"`
	MiddleName string `xml:"middle-name"`
	LastName   string `xml:"last-name"`
	Nickname   string `xml:"nickname"`
	HomePage   string `xml:"home-page"`
	Email      string `xml:"email"`
}
