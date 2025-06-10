
import * as monaco from 'monaco-editor'

export function registerCustomThemes() {
  monaco.editor.defineTheme('solarized', {
    base: 'vs',
    inherit: true,
    rules: [{ background: 'FDF6E3' }],
    colors: {
      'editor.background': '#FDF6E3',
      'editor.foreground': '#657b83',
      'editorLineNumber.foreground': '#93a1a1'
    }
  })

  monaco.editor.defineTheme('dracula', {
    base: 'vs-dark',
    inherit: true,
    rules: [{ background: '282A36' }],
    colors: {
      'editor.background': '#282A36',
      'editor.foreground': '#F8F8F2',
      'editorLineNumber.foreground': '#6272A4'
    }
  })
}
