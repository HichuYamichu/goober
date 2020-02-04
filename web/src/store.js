import { writable } from 'svelte/store';

let user = writable({
  username: '',
  quota: 0,
  admin: false
});

export { user };
