package main

import fmt "fmt"

//Page structure
type Page struct {
	Title string
	Body  []byte
}

func main() {
	//fmt.Print("asd")

	k := &Page{Title: "aaaa", Body: nil}
	k.Title = "kkkk"
	fmt.Print(k.Title)

}
