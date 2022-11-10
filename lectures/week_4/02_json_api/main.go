package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
)

type Todo struct {
	Name        string  `json:"name"`
	Description string  `json:"task_description"`
	Done        bool    `json:"done,omitempty"`
	SomeData    *string `json:"someData,omitempty"`
}

var (
	aa = "Some\nData\nHere\twritten\n"
	bb = `Some \n Data\n Here	
	written
asda"sd" '''asdasdasd'
 asdasdasd
asdasdasd
asdad
`
)

//type Todo struct {
//	Name        string
//	Description string
//	Done        bool
//	SomeData    []byte
//}

func main() {
	todos := []Todo{
		{"Выучить Go", "", false, &aa},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// здесь надо отдать статический файл, который будет общаться с API из браузера
		// открываем файл
		fileContents, err := os.ReadFile("index.html")
		w.Header().Add("Content-Type", "text/html")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// и выводим содержимое файла
		w.Write(fileContents)
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request ", r.URL.Path)
		defer r.Body.Close()

		// разные методы обрабатываются по-разному
		switch r.Method {
		// GET для получения данных
		case http.MethodGet:
			// преобразуем структуру в json
			productsJson, _ := json.Marshal(todos)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(productsJson)
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
	})

	http.ListenAndServe(":8080", nil)
}
