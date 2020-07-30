package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var todoTable map[int]string

var id int

func init() {
	todoTable = make(map[int]string)
	id = 1
}

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "アクセスされたよ")
} )
	http.HandleFunc("/hoge", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hoge")
} )
	http.HandleFunc("/todo", todoHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		for i, todo := range todoTable {
			fmt.Fprintf(w, "%d\t%s\n", i, todo)
		}
        case "POST":
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()
		todoTable[id] = body
 		id++
		fmt.Fprint(w, "Success\n")
	case "PUT":
	case "DELETE":
	}
}