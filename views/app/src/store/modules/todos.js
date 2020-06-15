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
    var oldTodo = state.todos.find(function(element) {
      return element.Id == todo.Id;
    });
    var index = state.todos.indexOf(oldTodo);
    // state.todos[index] = todo;
    state.todos.splice(index, 1, todo);
  },

  completeTodo(state, todo) {
    var oldTodo = state.todos.find(function(element) {
      return element.Id == todo.Id;
    });
    var index = state.todos.indexOf(oldTodo);
    // state.todos[index] = todo;
    state.todos.splice(index, 1, todo);
  },

  removeTodo(state, todo) {
    state.todos.splice(state.todos.indexOf(todo), 1);
  },
};

const actions = {
  setVisibility(context, visibility) {
    context.commit("setVisibility", visibility);
  },

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

  updateTodo(context, todo) {
    axios
      .post("/item/edit", {
        Id: parseInt(todo.Id),
        Title: todo.Title,
        Description: todo.Description,
        Url: todo.Url,
        Done: todo.Done,
      })
      .then((result) => {
        // todo: less lazy way
        // dispatch("loadItems");
        context.commit("updateTodo", result.data);
      })
      .catch((error) => {
        throw new Error(`API ${error}`);
      });
  },

  completeTodo(context, todo) {
    axios
      .post("/item/complete", {
        Id: parseInt(todo.Id),
        Done: todo.Done,
      })
      .then((result) => {
        // todo: less lazy way
        // dispatch("loadItems");
        context.commit("completeTodo", result.data);
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
