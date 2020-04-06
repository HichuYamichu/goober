import { user } from '../store';

let userValue;

user.subscribe((value) => {
  userValue = value;
});

export class HTTPClient {
  constructor() {
    this.json = 'application/json';
  }

  async get(endpoint) {
    return await fetch(endpoint, {
      method: 'GET',
      headers: {
        Accept: this.json,
        token: userValue.token,
      },
    });
  }

  async post(endpoint, body) {
    const headers = {
      Accept: this.json,
      token: userValue.token,
      'Content-Type': this.json,
    };
    if (body instanceof FormData) {
      delete headers['Content-Type'];
    }
    return await fetch(endpoint, { method: 'POST', headers, body });
  }

  async delete(endpoint, body) {
    return await fetch(endpoint, {
      method: 'DELETE',
      body,
      headers: {
        Accept: this.json,
        token: userValue.token,
        'Content-Type': this.json,
      },
    });
  }
}
