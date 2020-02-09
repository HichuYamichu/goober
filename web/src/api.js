class API {
  constructor() {
    this.accept = 'application/json';
    this.contentType = 'application/json';
  }

  get(endpoint) {
    return fetch(endpoint, {
      method: 'GET',
      headers: {
        Accept: this.accept,
        Authorization: `Bearer ${document.cookie}`
      }
    });
  }

  post(endpoint, body) {
    return fetch(endpoint, {
      method: 'POST',
      headers: {
        Accept: this.accept,
        Authorization: `Bearer ${document.cookie}`,
        'Content-Type': this.contentType
      },
      body: body
    });
  }
}

const api = new API();

export { api };
