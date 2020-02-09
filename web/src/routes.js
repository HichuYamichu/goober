import Index from './views/index.svelte';
import Login from './views/login.svelte';
import MainLayout from './MainLayout.svelte';

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
    onlyIf: { guard: loggedIn, redirect: '/auth' }
  },
  {
    name: 'auth/:inviteID',
    component: Login, 
    layout: MainLayout
  },
];

export { routes };
