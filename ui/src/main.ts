import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import BaseButtonVue from "./components/Buttons/BaseButton.vue";
import Router from "./router/index";
import { configure } from "vee-validate";
import axios from "axios";

axios.defaults.withCredentials = true

configure({
    validateOnBlur: true,
    validateOnChange: true,
    validateOnInput: true,
})

const app = createApp(App);
app.component("baseButton", BaseButtonVue);
app.use(Router).mount("#app");
