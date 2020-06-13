import Vue from "vue";
import Vuex from "vuex";
// import createLogger from "vuex/dist/logger";
// import todos from "./modules/todos";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    visibility: "all", // The TV inventory
  },

  getters: {
    // Here we will create a getter
  },

  mutations: {
    setVisibility(state, visibility) {
      state.visibility = visibility;
    },
  },

  actions: {
    setVisibility(context, visibility) {
      context.commit("setVisibility", visibility);
    },
    // Here we will create Larry
  },
});
