<template>
  <div class="container">
    <h1>Vue Todo App</h1>

    <!-- Add Todo -->
    <input v-model="newTodo" @keyup.enter="addTodo" placeholder="Add a task" />

    <!-- Todo List -->
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        <input type="checkbox" v-model="todo.completed" @change="updateTodo(todo)" />
        {{ todo.title }}
        <button @click="deleteTodo(todo.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      todos: [],
      newTodo: "",
    };
  },
  created() {
    this.getTodos();
  },
  methods: {
    // GET all todos
    getTodos() {
      axios.get("http://127.0.0.1:8081/todos")
        .then(res => this.todos = res.data);
    },

    // POST new todo
    addTodo() {
      if (!this.newTodo) return;
      axios.post("http://127.0.0.1:8081/todos", {
        title: this.newTodo,
        completed: false
      }).then(() => {
        this.newTodo = "";
        this.getTodos();
      });
    },

    // PUT update todo
    updateTodo(todo) {
      axios.put(`http://127.0.0.1:8081/todos/${todo.ID}`, {
        title: todo.Title,
        completed: todo.Completed
      }).then(() => this.getTodos());
    },

    // DELETE todo
    deleteTodo(id) {
      axios.delete(`http://127.0.0.1:8081/todos/${id}`)
        .then(() => this.getTodos());
    }
  }
};
</script>

<style>
.container {
  width: 400px;
  margin: auto;
  padding-top: 50px;
}
input[type="text"] {
  width: 70%;
  padding: 5px;
}
button {
  margin-left: 10px;
}
</style>
