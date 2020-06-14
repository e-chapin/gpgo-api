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

  created() {
    this.$store.dispatch("tds/loadItems");
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
