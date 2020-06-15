import Vue from "vue";
import Vuex from "vuex";
import VuexPersist from "vuex-persist";
import axios from "axios";
import VueAxios from "vue-axios";
import createLogger from "vuex/dist/logger";
import todos from "./modules/todos";

const vuexPersist = new VuexPersist({
  key: "gpgo",
  storage: window.localStorage,
});

Vue.use(Vuex, VueAxios, axios);

// add these before Vue is instantiated
window.axios = axios;
axios.defaults.baseURL = "http://127.0.0.1:5000/";

const debug = process.env.NODE_ENV !== "production";
var plugins = debug
  ? [createLogger(), vuexPersist.plugin]
  : [vuexPersist.plugin];

export default new Vuex.Store({
  plugins: plugins,
  modules: {
    tds: todos,
  },
  strict: debug,
});
