import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import "vuetify/styles";
import { createVuetify } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi";
import { createPinia } from "pinia";

const vuetify = createVuetify({
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
  defaults: {
    VAutocomplete: {
      variant: "outlined",
      density: "compact",
      bgColor: "#ffffff",
    },
    VSelect: {
      variant: "outlined",
      density: "compact",
      bgColor: "#ffffff",
    },
    VTextField: {
      variant: "outlined",
      density: "compact",
      bgColor: "#ffffff",
    },
    VTextarea: {
      variant: "outlined",
      density: "compact",
      bgColor: "#ffffff",
    },
  },
});
const pinia = createPinia();
const app = createApp(App);

app.use(vuetify).use(pinia).mount("#app");
