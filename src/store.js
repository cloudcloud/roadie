import Vue from 'vue';
import Vuex from 'vuex';

import apiClient from './api';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    sources: [],
    destinations: [],
    historical: [],
    source: {},
    copy: {
      loading: false,
      source: {},
      destination: {},
    },
    destination: [],
    subSource: {},
  },
  mutations: {
    resetConfig(state, config) {
      state.config = config;
    },
    resetHistorical(state, historical) {
      state.historical = historical;
    },
    resetSources(state, sources) {
      state.sources = sources;
    },
    resetDestinations(state, destinations) {
      state.destinations = destinations;
    },
    resetSource(state, source) {
      state.source = source;
    },
    resetDestination(state, destination) {
      state.destination = destination;
    },
    resetCopy(state, {source, destination}) {
      state.copy = {
        loading: false,
        source: source,
        destination: destination,
      };
    },
    resetSubSource(state, source) {
      state.subSource = source;
    },
  },
  getters: {
    allConfig: state => {
      return state.config;
    },
    allHistorical: state => {
      return state.historical;
    },
    allSources: state => {
      return state.sources;
    },
    allDestinations: state => {
      return state.destinations;
    },
    allSource: state => {
      return state.source;
    },
    allDestination: state => {
      return state.destination;
    },
    allSubSource: state => {
      return state.subSource;
    },
    getCopyState: state => {
      return state.copy;
    },
  },
  actions: {
    getConfig({commit}) {
      return new Promise((resolve) => {
        apiClient.getConfig().then((data) => {
          commit('resetConfig', data.items);
          resolve();
        });
      });
    },

    getDestinations({commit}) {
      return new Promise((resolve) => {
        apiClient.getDestinations().then((data) => {
          commit('resetDestinations', data.items);
          resolve();
        });
      });
    },

    getDestination({commit}, destination_name) {
      return new Promise((resolve) => {
        apiClient.getDestination(destination_name).then((data) => {
          commit('resetDestination', data.items);
          resolve();
        });
      });
    },

    getHistorical({commit}) {
      return new Promise((resolve) => {
        apiClient.getHistorical().then((data) => {
          commit('resetHistorical', data.items);
          resolve();
        });
      });
    },

    getSources({commit}) {
      return new Promise((resolve) => {
        apiClient.getSources().then((data) => {
          commit('resetSources', data.items);
          resolve();
        });
      });
    },

    getSource({commit}, source_name) {
      return new Promise((resolve) => {
        apiClient.getSource(source_name).then((data) => {
          commit('resetSource', data.items);
          resolve();
        });
      });
    },

    getSubSource({commit}, {source_name, sub_name}) {
      return new Promise((resolve) => {
        apiClient.getSubSource(source_name, sub_name).then((data) => {
          commit('resetSubSource', data.items);
          resolve();
        });
      });
    },

    pushCopy({commit}, payload) {
      return apiClient.pushCopy(payload).then((data) => {
        commit('resetCopy', data.items);
      });
    },

    removeFile(_, payload) {
      return apiClient.removeFile(payload);
    },
  }
});
