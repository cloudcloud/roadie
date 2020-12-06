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
  },
  mutations: {
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
  },
  getters: {
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
    getCopyState: state => {
      return state.copy;
    },
  },
  actions: {
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

    pushCopy({commit}, payload) {
      return apiClient.pushCopy(payload).then((data) => {
        commit('resetCopy', data.items);
      });
    },
  }
});
