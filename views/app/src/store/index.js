import Vue from "vue";
import Vuex from "vuex";
// import createLogger from "vuex/dist/logger";
// import todos from "./modules/todos";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    visibility: "all", // The TV inventory
    todos: [],
  },

  getters: {
    // Here we will create a getter
  },

  mutations: {
    setVisibility(state, visibility) {
      state.visibility = visibility;
    },

    addTodo(state, todo) {
      state.todos.push(todo);
    },
  },

  actions: {
    setVisibility(context, visibility) {
      context.commit("setVisibility", visibility);
    },

    addTodo(context, todo) {
      context.commit("addTodo", todo);
    },
    // Here we will create Larry
  },
});
