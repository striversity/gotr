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
		Type string `json:"type,omitempty"` // hello, add, or remove
		Todo `json:"todo,omitempty"`
		ID   int `json:"id,omitempty"`
	}
	ClientResponse struct {
		Todos `json:"todos,omitempty"`
	}
)

var upgrader websocket.Upgrader
var todos Todos
var todoID int

func main() {
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
		log.Infof("Message from clinet: %v", clientReq)

		clientResp := &ClientResponse{}
		switch clientReq.Type {
		case "add":
			todoID++
			clientReq.Todo.ID = todoID
			todos = append(todos, clientReq.Todo)
			log.Infof("All todos: %v", todos)
		case "delete":
			removeTodo(clientReq.ID)
		}

		clientResp.Todos = todos
		log.Infof("Message ==> clinet: %v", clientResp)
		conn.WriteJSON(clientResp)
	}
}

func removeTodo(id int) {
	var tmp Todos
	for _, v := range todos {
		if id != v.ID {
			tmp = append(tmp, v)
		}
	}
	todos = tmp
}
