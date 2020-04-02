import { HTTPClient } from './client';

class API {
  constructor() {
    this.client = new HTTPClient();
  }
  async upload(file) {
    const res = await this.client.post('/api/files', file);
    const data = res.json();
    return data;
  }

  async getFiles() {
    const res = await this.client.get('/api/files');
    const data = res.json();
    return data;
  }

  async deleteFile(filename) {
    const res = await this.client.delete(`/api/files/${filename}`);
    const data = res.json();
    return data;
  }

  async login(username, password) {
    const res = await this.client.post('/api/auth/login', JSON.stringify({ username, password }));
    const data = res.json();
    return data;
  }

  async register(username, password) {
    const res = await this.client.post(
      '/api/auth/register',
      JSON.stringify({ username, password })
    );
    const data = res.json();
    return data;
  }
  async changePassword(password) {
    const res = await this.client.post('/api/users/password/change', JSON.stringify({ password }));
    const data = res.json();
    return data;
  }

  async activateUser(id) {
    const res = await this.client.get(`/api/users/activate/${id}`);
    const data = res.json();
    return data;
  }

  async deleteUser(id) {
    const res = await this.client.delete(`/api/users/${id}`);
    const data = res.json();
    return data;
  }

  async getUsers() {
    const res = await this.client.get('/api/users');
    const data = res.json();
    return data;
  }
}

export const api = new API();
