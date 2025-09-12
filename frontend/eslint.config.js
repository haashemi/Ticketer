import { defineConfig } from "@fullstacksjs/eslint-config";
import eslintPluginSvelte from "eslint-plugin-svelte";
import ts from "typescript-eslint";

import svelteConfig from "./svelte.config.js";

export default defineConfig({}, ...eslintPluginSvelte.configs.recommended, {
  files: ["**/*.svelte", "**/*.svelte.ts", "**/*.svelte.js"],
  languageOptions: {
    parserOptions: {
      projectService: true,
      extraFileExtensions: [".svelte"],
      parser: ts.parser,
      svelteConfig,
    },
  },
});
