<template>
  <div>
    <todo
      v-on:delete-todo="deleteTodo"
      v-on:update-todo="updateTodo"
      v-for="todo in filteredTodos"
      v-bind:todo="todo"
      v-bind:visibility="visibility"
      v-bind:key="todo.id"
    ></todo>
    <footer class="footer ui centered card" v-show="todos.length" v-cloak>
      <div class="content">
        <span class="todo-count">
          <strong>{{ remaining }}</strong>
          {{ remaining | pluralize }} left
        </span>
        <div>
          <a
            v-on:click="setVisibility('all')"
            :class="{ selected: getVisibility == 'all' }"
            >All</a
          >
        </div>
        <div>
          <a
            v-on:click="setVisibility('active')"
            :class="{ selected: getVisibility == 'active' }"
            >Active</a
          >
        </div>
        <div>
          <a
            v-on:click="setVisibility('completed')"
            :class="{ selected: getVisibility == 'completed' }"
            >Completed</a
          >
        </div>
        <button
          class="clear-completed"
          @click="removeCompleted"
          v-show="todos.length > remaining"
        >
          Clear completed
        </button>
      </div>
    </footer>
  </div>
</template>

<script type="text/javascript">
import Todo from "./todo";
export default {
  props: ["todos", "filters", "visibility"],
  components: {
    Todo,
  },

  methods: {
    deleteTodo(todo) {
      this.$store.dispatch("tds/removeTodo", todo);
    },
    updateTodo(todo) {
      this.$store.dispatch("tds/updateTodo", todo);
    },
    setVisibility(visibility) {
      this.$emit("set-visibility", visibility);
    },
    removeCompleted: function() {
      this.todos = this.filters.active(this.todos);
    },
  },
  computed: {
    filteredTodos: function() {
      return this.filters[this.visibility](this.todos);
    },
    getVisibility: function() {
      return this.visibility;
    },
    remaining: function() {
      return this.filters.active(this.todos).length;
    },
    allDone: {
      get: function() {
        return this.remaining === 0;
      },
      set: function(value) {
        this.todos.forEach(function(todo) {
          todo.completed = value;
        });
      },
    },
  },

  filters: {
    pluralize: function(n) {
      return n === 1 ? "item" : "items";
    },
  },
};
</script>
