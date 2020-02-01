import Index from './views/index.svelte';
import Login from './views/login.svelte';

function userIsAdmin() {
  return true;
}

const routes = [
  {
    name: '/',
    component: Index,
    onlyIf: { guard: userIsAdmin, redirect: '/login' }
  },
  {
    name: 'login',
    component: Login
  }
];

export { routes };
