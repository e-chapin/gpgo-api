<template>
  <div>
    <create-todo v-on:create-todo="addTodo"></create-todo>
    <todo-list
      v-on:set-visibility="setVisibility"
      v-bind:todos="todos"
      v-bind:filters="filters"
      v-bind:visibility="visibility"
    ></todo-list>
  </div>
</template>

<script>
import TodoList from "./components/TodoList";
import CreateTodo from "./components/CreateTodo";

// var STORAGE_KEY = "todos-vuejs-2.0";
// var todoStorage = {
//   fetch: function() {
//     var todos = JSON.parse(localStorage.getItem(STORAGE_KEY) || "[]");
//     todos.forEach(function(todo, index) {
//       todo.id = index;
//     });
//     // todoStorage.uid = todos.length;
//     return todos;
//   },
//   save: function(todos) {
//     localStorage.setItem(STORAGE_KEY, JSON.stringify(todos));
//   },
// };

// visibility filters
var filters = {
  all: function(todos) {
    return todos;
  },
  active: function(todos) {
    return todos.filter(function(todo) {
      return !todo.done;
    });
  },
  completed: function(todos) {
    return todos.filter(function(todo) {
      return todo.done;
    });
  },
};

// handle routing

import { mapState } from "vuex";
export default {
  name: "app",
  components: {
    TodoList,
    CreateTodo,
  },

  methods: {
    addTodo(todo) {
      this.$store.dispatch("tds/addTodo", todo);
    },
    setVisibility(visibility) {
      this.$store.dispatch("tds/setVisibility", visibility);
    },
  },

  computed: {
    ...mapState({
      visibility: (state) => state.tds.visibility,
      todos: (state) => state.tds.todos,
    }),
  },

  data() {
    return {
      filters: filters,
      newTodo: "",
      editedTodo: null,
    };
  },
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
