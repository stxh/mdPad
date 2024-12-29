import './style.css'
import App from './App.svelte'

const app = new App({
  target: document.getElementById('app')
})

export default app

// interface Result {
//   value: string;
//   error: string;
// }