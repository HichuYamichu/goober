import Index from './views/index.svelte';
import Login from './views/login.svelte';

function loggedIn() {
  if (document.cookie) {
    return true;
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
  }
];

export { routes };
