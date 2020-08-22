package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	// resp := []byte(`{"name":"Punyapat"}`)
	// w.Header().Add("content-type", "application/json")
	// w.Write(resp)
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
		}
		// todos = []Todo{t}
		todos = append(todos, t)
		fmt.Printf("body : %#v\n", t)
		fmt.Printf("todos : %#v\n", todos)
		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}
	if req.Method == "GET" {
		resp, err := json.Marshal(todos)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
		}
		w.Header().Add("content-type", "application/json")
		w.Write(resp)
		return
	}
	if req.Method == "PUT" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		t := Todo{}
		idx := -1

		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
		}
		for i := range todos {
			if todos[i].ID == t.ID {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Fprintf(w, "error item ID : %v not found", t.ID)
		}
		todos[idx] = t
		// todos[idx].Title = t.Title
		fmt.Printf("body : %#v\n", t)
		fmt.Printf("todos : %#v\n", todos)
		fmt.Fprintf(w, "hello %s created todos", "POST")
		return
	}
}

func main() {
	http.HandleFunc("/todos", helloHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
