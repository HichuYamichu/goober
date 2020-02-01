import { writable } from 'svelte/store';

let token = writable('');

export { token };
