import fs from 'fs';
import path from 'path';
import hljs from 'highlight.js/lib/core';

hljs.registerLanguage('go', require('highlight.js/lib/languages/go'))

const codesDirectory = path.join(process.cwd(), 'posts', 'codes');

export function getAllCodeIds() {
    const fileNames = fs.readdirSync(codesDirectory);
    return fileNames.filter((filename) => filename.endsWith(".go"))
        .map((fileName) => {
            return {
                params: {
                    id: fileName.replace(/\.go$/, ''),
                },
            };
        });
}

export async function getCodeData(id) {
    const fullPath = path.join(codesDirectory, `${id}.go`);
    const fileContent = fs.readFileSync(fullPath, 'utf8');

    return {
        id,
        html: hljs.highlight(fileContent, {language: 'go'}).value
    };
}