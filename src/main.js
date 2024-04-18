import { createApp } from 'vue';
import App from './App.vue';

import 'vuetify/styles';

import vuetify from './plugins/vuetify';
import router from './routes';
import store from './store';

createApp(App).use(vuetify).use(router).use(store).mount('#app');
