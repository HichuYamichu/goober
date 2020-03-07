export class HTTPClient {
  constructor() {
    this.accept = 'application/json';
  }

  async get(endpoint) {
    return await fetch(endpoint, {
      method: 'GET',
      headers: {
        Accept: this.accept,
        Authorization: `Bearer ${document.cookie}`
      }
    });
  }

  async post(endpoint, body) {
    const headers = {
      Accept: this.accept,
      Authorization: `Bearer ${document.cookie}`,
      'Content-Type': 'application/json'
    };
    if (body instanceof FormData) {
      delete headers['Content-Type'];
    }
    return await fetch(endpoint, { method: 'POST', headers, body });
  }

  async delete(endpoint) {
    return await fetch(endpoint, {
      method: 'DELETE',
      headers: {
        Accept: this.accept,
        Authorization: `Bearer ${document.cookie}`
      }
    });
  }
}
