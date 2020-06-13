<template>
  <div>
    <todo v-for="todo in todos" v-bind:todo="todo" v-bind:key="todo.id"></todo>
  </div>
</template>

<script type = "text/javascript" >
import Todo from "./todo";
export default {
  props: ["todos"],
  components: {
    Todo
  }
};
</script>