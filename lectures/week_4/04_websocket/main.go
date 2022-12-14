package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type JokeChuck struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconUrl    string        `json:"icon_url"`
	Id         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	Url        string        `json:"url"`
	Value      string        `json:"value"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// "Шина" событий, регистрация клиентов и рассылка сообщений идет отсюда
type Bus struct {
	register  chan *websocket.Conn
	broadcast chan []byte
	clients   map[*websocket.Conn]bool
}

func (b *Bus) Run() {
	for {
		select {
		case message := <-b.broadcast:
			// каждому зарегистрированному клиенту шлем сообщение
			for client := range b.clients {
				w, err := client.NextWriter(websocket.TextMessage)
				if err != nil {
					// если достучаться до клиента не удалось, то удаляем его
					log.Printf("connection error: %v\n", err)
					delete(b.clients, client)
					continue
				}

				w.Write(message)
			}
		case client := <-b.register:
			// регистрируем клиентов в мапе клиентов
			log.Println("User registered")
			b.clients[client] = true
		}
	}
}

func NewBus() *Bus {
	return &Bus{
		register:  make(chan *websocket.Conn),
		broadcast: make(chan []byte),
		clients:   make(map[*websocket.Conn]bool),
	}
}

func runJoker(b *Bus) {
	for {
		// каждые 5 секунд ходим за шутками
		<-time.After(5 * time.Second)
		log.Println("Its joke time!")
		b.broadcast <- getJoke()
	}
}

func getJoke() []byte {
	c := http.Client{}
	resp, err := c.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return []byte("jokes API not responding")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("status")
	}
	body, _ := io.ReadAll(resp.Body)
	joke := JokeChuck{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		log.Fatal(err)
		return []byte("Joke error")
	}

	return []byte(joke.Value)
}

//go:embed index.html
var htmlData []byte

func main() {
	bus := NewBus()
	go bus.Run()
	go runJoker(bus)
	socketMux := &http.ServeMux{}
	socketMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// апгрейд соединения
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrader error:", err)
			return
		}

		bus.register <- ws
	})
	httpMux := &http.ServeMux{}

	httpMux.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Keep-Alive", "timeout=32, max=100")
		w.Header().Add("Connection", "Keep-Alive")
		w.Header().Add("Origin", "localhost")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(htmlData); err != nil {
			log.Println("ERROR:", err)
		}
	})
	go func() {
		if err := http.ListenAndServe(":8081", socketMux); err != nil {
			log.Fatal(err)
		}
	}()
	if err := http.ListenAndServe(":8080", httpMux); err != nil {
		log.Fatal(err)
	}
}
