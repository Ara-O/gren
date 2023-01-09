import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import BaseButtonVue from "./components/Buttons/BaseButton.vue";
import Router from "./router/index";
const app = createApp(App);
app.component("baseButton", BaseButtonVue);
app.use(Router).mount("#app");
