import { mdsvex, escapeSvelte } from 'mdsvex';
import { createHighlighter } from 'shiki';
import adapter from '@sveltejs/adapter-auto';

const theme = 'catppuccin-macchiato'
const highlighter = await createHighlighter({
	themes: [theme],
	langs: ['cpp']
})

const mdsvexOptions = {
	extensions: ['.md'],
		highlight: {
			highlighter: async (code, lang = 'text') => {
				await highlighter.loadLanguage('cpp')
				const html = escapeSvelte(highlighter.codeToHtml(code, { lang, theme }))
				return `{@html \`${html}\` }`
			}
	}
}

const config = {
	kit: { adapter: adapter() },
	preprocess: [mdsvex(mdsvexOptions)],
	extensions: ['.svelte', '.md'],
	optimizeDeps: {
		exclude: ["svelte-codemirror-editor", "codemirror", "@codemirror/language-cpp", "@codemirror/view", "@codemirror/state", "@codemirror/commands", "@codemirror/language"]
	}
};

export default config;