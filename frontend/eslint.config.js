import eslintPluginSvelte from 'eslint-plugin-svelte';
import eslintConfigPrettier from 'eslint-config-prettier';
import * as svelteParser from 'svelte-eslint-parser';
import * as typescriptParser from '@typescript-eslint/parser';

export default [
	...eslintPluginSvelte.configs['flat/recommended'],
	eslintConfigPrettier,
	{
		files: ['**/*.svelte', '**/*.ts', '**/*.js'],
		languageOptions: {
			parser: svelteParser,
			parserOptions: {
				parser: typescriptParser,
				project: './tsconfig.json',
				extraFileExtensions: ['.svelte'],
			},
		},
		ignores: ['./.svelte-kit/**/*', './node_modules/**/*'],
	},
];
