const state = () => ({
  visibility: "all",
  todos: [],
});

const getters = {
  // Here we will create a getter
};

const mutations = {
  setVisibility(state, visibility) {
    state.visibility = visibility;
  },

  addTodo(state, todo) {
    state.todos.push(todo);
  },

  removeTodo(state, todo) {
    state.todos.splice(state.todos.indexOf(todo), 1);
  },

  updateTodo(state, todo) {
    var index = state.todos.indexOf(todo);
    state.todos[index] = todo;
  },
};

const actions = {
  setVisibility(context, visibility) {
    context.commit("setVisibility", visibility);
  },

  addTodo(context, todo) {
    context.commit("addTodo", todo);
  },

  removeTodo(context, todo) {
    context.commit("removeTodo", todo);
  },

  updateTodo(context, todo) {
    context.commit("updateTodo", todo);
  },
  // Here we will create Larry
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
