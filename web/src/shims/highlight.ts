// discord-markdown pulls in the full highlight.js bundle (~800KB minified,
// every language) just to colorize fenced code blocks in the preview.
// Vite aliases "highlight.js" to this module so only the core engine ships.
// Register languages here if code highlighting in the preview matters:
//
//   import javascript from "highlight.js/lib/languages/javascript";
//   hljs.registerLanguage("javascript", javascript);
import hljs from "highlight.js/lib/core";

export default hljs;
