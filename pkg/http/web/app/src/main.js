import { createApp } from "vue";
import App from "./App.vue";
import "primeflex/primeflex.css";
import "primevue/resources/themes/saga-blue/theme.css";
import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";
import PrimeVue from "primevue/config";
import Menubar from "primevue/menubar";
import router from "./router";
import Card from "primevue/card";
import InputText from "primevue/inputtext";
import Divider from "primevue/divider";

const app = createApp(App);
app.use(PrimeVue, { ripple: true });
app.use(router);

app.component("Menubar", Menubar);
app.component("Card", Card);
app.component("InputText", InputText);
app.component("Divider", Divider);

app.mount("#app");
