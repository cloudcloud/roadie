import axios from 'axios';

const el = document.getElementById('config');
const c = JSON.parse(el.innerHTML);
const client = axios.create({
  baseURL: c.hostname,
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

  pushCopy(payload) {
    return this.perform('post', '/api/v1/execute', payload);
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
