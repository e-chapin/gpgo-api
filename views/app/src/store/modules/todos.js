import axios from "axios";

const state = () => ({
  visibility: "all",
  todos: [],
});

const getters = {
  // Here we will create a getter
};

const mutations = {
  saveItems(state, items) {
    state.todos = items;
  },

  setVisibility(state, visibility) {
    state.visibility = visibility;
  },

  addTodo(state, todo) {
    if (state.todos != null) {
      state.todos.push(todo);
    } else {
      state.todos = [todo];
    }
  },

  updateTodo(state, todo) {
    var index = state.todos.indexOf(todo);
    var tds = state.todos;
    tds.forEach();
    state.todos[index] = todo;
  },

  removeTodo(state, todo) {
    state.todos.splice(state.todos.indexOf(todo), 1);
  },
};

const actions = {
  loadItems(context) {
    axios
      .get("/item/all")
      .then((result) => {
        context.commit("saveItems", result.data);
      })
      .catch((error) => {
        throw new Error(`API ${error}`);
      });
  },

  setVisibility(context, visibility) {
    context.commit("setVisibility", visibility);
  },

  addTodo(context, todo) {
    axios
      .post("/item/new", {
        Title: todo.Title,
        Description: todo.Description,
        Url: todo.Url,
      })
      .then((result) => {
        context.commit("addTodo", result.data);
      })
      .catch((error) => {
        throw new Error(`API ${error}`);
      });
  },

  updateTodo({ dispatch }, todo) {
    axios
      .post("/item/edit", {
        Id: parseInt(todo.Id),
        Title: todo.Title,
        Description: todo.Description,
        Url: todo.Url,
      })
      .then(() => {
        dispatch("loadItems");
      })
      .catch((error) => {
        throw new Error(`API ${error}`);
      });
  },

  removeTodo(context, todo) {
    axios
      .delete("/item/id/" + todo.Id)
      .then(() => {
        context.commit("removeTodo", todo);
      })
      .catch((error) => {
        throw new Error(`API ${error}`);
      });
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
