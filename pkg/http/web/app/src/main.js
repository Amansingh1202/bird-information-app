import { createApp } from "vue";
import App from "./App.vue";
import "primeflex/primeflex.css";
import "primevue/resources/themes/saga-blue/theme.css";
import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import PrimeVue from "primevue/config";
import Menubar from "primevue/menubar";
import router from "./router";

const app = createApp(App);
app.use(PrimeVue, { ripple: true });
app.use(router);

app.component("Menubar", Menubar);

app.mount("#app");
