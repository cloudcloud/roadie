import axios from 'axios';

const client = axios.create({
  baseURL: 'http://localhost:8008',
  json: true,
})

const apiClient = {
  getDestinations() {
    return this.perform('get', '/api/v1/destinations');
  },

  getHistorical() {
    return this.perform('get', '/api/v1/historical');
  },

  getSource(source_name) {
    return this.perform('get', `/api/v1/sources/${source_name}`);
  },

  getSources() {
    return this.perform('get', '/api/v1/sources');
  },

  async perform(method, resource, data) {
    return client({
      method,
      url: resource,
      data,
      headers: {
        'X-Client': 'roadie-web-client 1.0',
      }
    }).then(req => {
      return req.data;
    })
  }
};

export default apiClient;
