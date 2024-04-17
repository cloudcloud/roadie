import '@mdi/font/css/materialdesignicons.css';
import Vue from 'vue';
import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'

Vue.use(vuetify);

export default new Vuetify({
  icons: {
    iconfont: 'mdi',
  },
});
