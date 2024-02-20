import { defineStore } from "pinia";

export type AppTab = "schema" | "csv";

type AppStore = {
  appTab: AppTab;
};

export const useAppStore = defineStore("app", {
  state: (): AppStore => ({
    appTab: "schema",
  }),
  actions: {
    setAppTab(tab: "schema" | "csv") {
      this.appTab = tab;
    },
  },
});
