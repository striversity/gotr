import { Todo } from "./todo";

export class App {
  constructor() {
    var app = this;
    this.heading = 'Todos';
    this.todos = [];
    this.todoDescription = '';
    // Create WebSocket connection.
    this.socket = new WebSocket('ws://localhost:8081/ws');

    // on websocket error
    this.socket.addEventListener('error', function (event) {
      console.log(event);
    });

    // Connection opened
    this.socket.addEventListener('open', function (event) {
      app.socket.send('{"type":"hello"}');
    });

    // Listen for messages
    this.socket.addEventListener('message', function (event) {
      var msg = JSON.parse(event.data);
      app.todos =msg.todos;
    });
  }

  addTodo() {
    if (this.todoDescription) {
      var todo = new Todo(this.todoDescription);
      var msg = {"type":"add", "todo": todo}
      this.socket.send(JSON.stringify(msg));
      this.todoDescription = '';
    }
  }
  
  removeTodo(todo) {
    var msg = {"type":"delete", "id": todo.id}
    this.socket.send(JSON.stringify(msg));
  }
}
