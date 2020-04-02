import { HTTPClient } from './client';

class API {
  constructor() {
    this.client = new HTTPClient();
  }
  async upload(file) {
    const res = await this.client.post('/api/files/upload', file);
    const data = res.json();
    return data;
  }

  async getFiles() {
    const res = await this.client.get('/api/files/list');
    const data = res.json();
    return data;
  }

  async deleteFile(filename) {
    const res = await this.client.delete(`/api/files/delete/${filename}`);
    const data = res.json();
    return data;
  }

  async login(username, password) {
    const res = await this.client.post('/api/auth/login', JSON.stringify({ username, password }));
    const data = res.json();
    console.log(data);
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
    const res = await this.client.post('/api/user/password/change', JSON.stringify({ password }));
    const data = res.json();
    return data;
  }

  async activateUser(id) {
    const res = await this.client.post('/api/user/activate', JSON.stringify({ id }));
    const data = res.json();
    return data;
  }

  async deleteUser(id) {
    const res = await this.client.delete(`/api/user/delete/${id}`);
    const data = res.json();
    return data;
  }

  async getUsers() {
    const res = await this.client.get('/api/user/list');
    const data = res.json();
    return data;
  }
}

export const api = new API();
