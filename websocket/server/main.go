package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
)

type (
	Todo struct {
		ID          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Done        bool   `json:"done"`
	}
	Todos         []Todo
	ClientRequest struct {
		Username string `json:"username,omitempty"`
		Type     string `json:"type,omitempty"` // hello, add, or remove
		Todo     `json:"todo,omitempty"`
		ID       int `json:"id,omitempty"`
	}
	ClientResponse struct {
		Todos `json:"todos,omitempty"`
	}
)

var upgrader websocket.Upgrader
var db map[string]Todos
var todoID int

func main() {
	db = make(map[string]Todos)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	http.HandleFunc("/ws", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		clientReq := &ClientRequest{}
		err := conn.ReadJSON(clientReq)
		if err != nil {
			log.Println(err)
			return
		}
		log.Infof("Message from clinet: %#v", clientReq)
		if 0 == len(clientReq.Username) {
			return
		}

		clientResp := &ClientResponse{}
		var todos Todos
		switch clientReq.Type {
		case "hello":
			todos = getTodos(clientReq.Username)
		case "add":
			todos = addTodo(clientReq.Username, clientReq.Todo)
		case "delete":
			todos = removeTodo(clientReq.Username, clientReq.ID)
		case "toggle.done":
			todos = toggleDone(clientReq.Username, clientReq.ID)
		}

		clientResp.Todos = todos
		log.Infof("Message ==> clinet: %v", clientResp)
		conn.WriteJSON(clientResp)
	}
}

func toggleDone(username string, id int) Todos{
	todos := db[username]
	for i, v := range todos {
		if id == v.ID {
			todos[i].Done = !v.Done
		}
	}
	return todos
}
func getTodos(username string) Todos {
	return db[username]
}
func addTodo(username string, todo Todo) Todos {
	todoID++
	todo.ID = todoID
	todos := db[username]
	todos = append(todos, todo)
	db[username] = todos
	return todos
}

func removeTodo(username string, id int) Todos{
	todos := db[username]
	var tmp Todos
	for _, v := range todos {
		if id != v.ID {
			tmp = append(tmp, v)
		}
	}
	db[username] = tmp
	return tmp
}
