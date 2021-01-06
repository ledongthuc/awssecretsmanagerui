import Vue from "vue";
import Vuex from "vuex";
import secrets from "./modules/secrets";

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== "production";

export default new Vuex.Store({
  modules: {
    secrets,
  },
  strict: debug,
});
