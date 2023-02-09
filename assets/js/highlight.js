import hljs from 'highlight.js/lib/core';

import javascript from 'highlight.js/lib/languages/javascript';
import json from 'highlight.js/lib/languages/json';
import bash from 'highlight.js/lib/languages/bash';
import xml from 'highlight.js/lib/languages/xml';
import ini from 'highlight.js/lib/languages/ini';
import yaml from 'highlight.js/lib/languages/yaml';
import markdown from 'highlight.js/lib/languages/markdown';
import python from 'highlight.js/lib/languages/python';
import java from 'highlight.js/lib/languages/java'
import go from 'highlight.js/lib/languages/go'
import sql from 'highlight.js/lib/languages/sql'
import plaintext from 'highlight.js/lib/languages/plaintext.js'

hljs.registerLanguage('javascript', javascript);
hljs.registerLanguage('json', json);
hljs.registerLanguage('bash', bash);
hljs.registerLanguage('shell', bash);
hljs.registerLanguage('html', xml);
hljs.registerLanguage('ini', ini);
hljs.registerLanguage('conf', ini);
hljs.registerLanguage('toml', ini);
hljs.registerLanguage('yaml', yaml);
hljs.registerLanguage('md', markdown);
hljs.registerLanguage('python', python);
hljs.registerLanguage('java', java);
hljs.registerLanguage('go', go);
hljs.registerLanguage('mysql', sql);
hljs.registerLanguage('text', plaintext);

document.addEventListener('DOMContentLoaded', () => {
  document.querySelectorAll('pre code:not(.language-mermaid)').forEach((block) => {
    hljs.highlightElement(block);
  });
});
