import Vue from 'vue'
import vuetify from './plugins/vuetify';
import App from './App.vue'
import router from './routes'
import store from './store'

Vue.config.productionTip = process.env.NODE_ENV == 'production'

new Vue({
  vuetify,
  store,
  router,
  render: h => h(App),
}).$mount('#app');
