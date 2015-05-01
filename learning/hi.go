package main

import fmt "fmt"

//Page structure
type Page struct {
	Title string
	Body  []byte
}

func main() {
	//fmt.Print("asd")

	k := &Page{Title: "aaaa", Body: []byte{'a', 'b', 'c'}}
	k.Title = "name"
	fmt.Println("title : " + k.Title)
	fmt.Println("body : " + string(k.Body))

}
