import Index from './views/index.svelte';
import Login from './views/login.svelte';
import Admin from './views/admin.svelte';
import { user } from './store.js';

let userValue;

const unsubscribe = user.subscribe(user => {
  userValue = user;
});

function loggedIn() {
  if (document.cookie) {
    return true
  }
  return false;
}

function isAdmin() {
  user.subscribe(user => {
    userValue = user;
  });
  console.log(userValue)
  return userValue.admin
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
    component: Admin,
    onlyIf: { guard: isAdmin, redirect: '/login' }
  },
];

export { routes };
