import Index from './views/index.svelte';
import Login from './views/login.svelte';
import { token } from './store.js';

let tokenValue;

const unsubscribe = token.subscribe(token => {
  tokenValue = token;
});

function loggedIn() {
  if (tokenValue) {
    return true
  }
  return false;
}

const routes = [
  {
    name: '/',
    component: Index,
    onlyIf: { guard: loggedIn, redirect: '/login' }
  },
  {
    name: 'login',
    component: Login
  },
  {
    name: '/admin',
    component: Index,
    onlyIf: { guard: loggedIn, redirect: '/login' }
  },
];

export { routes };
