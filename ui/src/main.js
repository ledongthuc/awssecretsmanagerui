import Vue from "vue";
import VueRouter from "vue-router";
import App from "./App.vue";
import "./plugins/table.js";
import "./plugins/moment.js";
import "./plugins/axios.js";

import SecretTable from "./components/SecretTable.vue";
import SecretForm from "./components/SecretForm.vue";

Vue.use(VueRouter);
Vue.config.productionTip = false;

const router = new VueRouter({
  mode: "history",
  base: __dirname,
  routes: [
    { name: "home", path: "/", component: SecretTable },
    { name: "detail", path: "/detail", component: SecretForm },
  ],
});

new Vue({
  render: (h) => h(App),
  router,
}).$mount("#app");
