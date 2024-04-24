import axios from 'axios';

const url = document.querySelector('#RoadieBaseURL').getAttribute('content');
const client = axios.create({
  baseURL: url,
  json: true,
});

const apiClient = {
  addConfig(type, payload) {
    return this.perform('post', `/api/v1/config/add/${type}`, payload);
  },

  editConfig(type, name, payload) {
    return this.perform('put', `/api/v1/config/edit/${type}/${name}`, payload);
  },

  getConfig() {
    return this.perform('get', '/api/v1/config');
  },

  getDestination(destination_name) {
    return this.perform('get', `/api/v1/destinations/${destination_name}`)
  },

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

  getSubSource(source_name, sub_name) {
    return this.perform('get', `api/v1/sources/${source_name}/${sub_name}`);
  },

  pushCopy(payload) {
    return this.perform('post', '/api/v1/execute', payload);
  },

  removeFromConfig(type, name) {
    return this.perform('delete', `/api/v1/config/remove/${type}/${name}`);
  },

  removeFile(payload) {
    return this.perform('delete', '/api/v1/remove', payload);
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
