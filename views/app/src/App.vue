<template>
  <todo-list v-bind:todos="todos"></todo-list>
</template>

<script>
import TodoList from "./components/TodoList";
export default {
  name: "app",
  components: {
    TodoList
  },

  data() {
    return {
      todos: [
        {
          id: 1,
          title: "Todo A",
          project: "Project A",
          done: false
        },
        {
          id: 2,
          title: "Todo B",
          project: "Project B",
          done: true
        }
      ]
    };
  }
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