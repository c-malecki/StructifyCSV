/* eslint-env node */
require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
  root: true,
  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "@vue/eslint-config-typescript",
    "@vue/eslint-config-prettier/skip-formatting",
  ],
  parserOptions: {
    ecmaVersion: "latest",
  },
  rules: {
    "no-console": import.meta.env.MODE === "production" ? "warn" : "off",
    "no-debugger": import.meta.env.MODE === "production" ? "error" : "off",
    "no-reserved-component-names": "off",
    "no-unused-vars": "off",
    "prettier/prettier": [
      "warn",
      {
        endOfLine: "auto",
      },
    ],
  },
};
