export function setTheme(theme = 'light') {
  const head = document.head
  const existingLink = document.getElementById('theme-link')

  if (existingLink) {
    head.removeChild(existingLink)
  }

  const link = document.createElement('link')
  link.id = 'theme-link'
  link.rel = 'stylesheet'
  link.href = `/src/styles/themes/${theme}.css`
  head.appendChild(link)

  // 记住用户主题
  localStorage.setItem('theme', theme)
}
