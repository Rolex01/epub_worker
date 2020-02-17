package main

import (
	"github.com/rolex01/epub_workers/epub"
	"log"
)

func main() {
	e, err := epub.Open("Yanagihara Hanya. A Little Life - royallib.com.epub")
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()

	e.GetStyle()
	/*for i, val := range m {
		fmt.Println("Styles:", i, string(val))
	}*/

	//f, _ := e.OpenFile("style.css")
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(f)
	//fmt.Println("File:", buf.String())

	//MetadataAttr
	//for i, val := range e.MetadataFields() {
	//	meta, _ := e.MetadataAttr(val)
	//	for k, v := range meta {
	//		fmt.Println("MetadataAttr:", i, val, k, v)
	//	}
	//}

	//Navigation
	//n, _ := e.Navigation()
	//for true {
	//	fmt.Println(n.IsFirst(), n.IsLast(), n.HasChildren(), n.HasParents(), n.URL(), n.Title())
	//	if n.HasChildren() {
	//		// Recursive
	//	}
	//
	//	if n.IsLast() {
	//		break
	//	}
	//	err = n.Next()
	//	if err != nil {
	//		log.Fatal("NEXT Navigation: ", err)
	//	}
	//}

	//Metadata
	//for i, val := range e.MetadataFields() {
	//	meta, _ := e.Metadata(val)
	//	fmt.Println("Meta:", i, val, ":", meta)
	//}


	//xhtml Files
	//q, _ := e.Spine()
	//for true {
	//	fmt.Println(q.IsFirst(), q.IsLast(), q.URL())
	//	r, err := q.Open()
	//	defer r.Close()
	//	buf := new(bytes.Buffer)
	//	buf.ReadFrom(r)
	//	fmt.Println("File:", buf.String())
	//
	//	if q.IsLast() {
	//		break
	//	}
	//	err = q.Next()
	//	if err != nil {
	//		log.Fatal("NEXT Spine: ", err)
	//	}
	//}
}