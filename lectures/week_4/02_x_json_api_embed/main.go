package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done,omitempty"`
}

var (
	bindAddress = flag.String("bind", "localhost", "network address to bind")
	port        = flag.Int("port", 8080, "network port to bind")
	firstToDo   = flag.String("first", "Learn Go", "just init TODO list with first phase")
)

/*
//go:embed index.html
var fileContents []byte
*/

//go:embed files
var fldr embed.FS

var todos = []Todo{
	{*firstToDo, false},
}

func main() {
	flag.Parse()
	todos[0].Name = *firstToDo
	http.HandleFunc("/", serveHtml)

	http.HandleFunc("/todos/", restAPI)
	log.Printf("serve at http://%s:%d", *bindAddress, *port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", *bindAddress, *port), nil); err != nil {
		log.Fatalf("can't start: %v", err)
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Content-Type", "text/html")
	fileContents, err := fldr.ReadFile("files/index.html")
	if err != nil {
		log.Printf("can't read: %v", err)
	}
	_, _ = w.Write(fileContents)
}

func restAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("request ", r.URL.String())
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	// разные методы обрабатываются по-разному
	switch r.Method {
	// GET для получения данных
	case http.MethodGet:
		// преобразуем структуру в json
		productsJson, _ := json.Marshal(todos)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(productsJson)
	// POST для добавления чего-то нового
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		todo := Todo{}
		// преобразуем json запрос в структуру
		err := decoder.Decode(&todo)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todos = append(todos, todo)
	// PUT для обновления существующей информации
	case http.MethodPut:
		id := r.URL.Path[len("/todos/"):]
		index, _ := strconv.ParseInt(id, 10, 0)
		todos[index].Done = true
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
