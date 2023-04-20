import '../../theme/theme.css'
import './app.css'
import App from './App.svelte'

let appElement = document.getElementById('app') || new Element()

const app = new App({
  target: appElement,
})

export default app
