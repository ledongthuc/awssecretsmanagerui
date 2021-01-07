// import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap-table/dist/bootstrap-table.min.css";

import "./jquery.js";
import Vue from "vue";
// import "bootstrap";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-table/dist/bootstrap-table.js";
import BootstrapTable from "bootstrap-table/dist/bootstrap-table-vue.esm.js";

Vue.component("BootstrapTable", BootstrapTable);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
