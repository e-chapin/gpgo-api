<template>
  <div>
    <todo
      v-on:delete-todo="deleteTodo"
      v-for="todo in todos"
      v-bind:todo="todo"
      v-bind:key="todo.id"
    ></todo>
  </div>
</template>

<script type = "text/javascript" >
import Todo from "./todo";
export default {
  props: ["todos"],
  components: {
    Todo
  },
  methods: {
    deleteTodo(todo) {
      const todoIndex = this.todos.indexOf(todo);
      this.todos.splice(todoIndex, 1);
    }
    // completeTodo(todo){
    //   todo.COMPL
    // }
  }
};
</script>