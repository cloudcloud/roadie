import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import Sources from './components/Sources';
import Source from './components/Source';
import Destinations from './components/Destinations';
import Destination from './components/Destination';
import Config from './components/Config';

Vue.use(VueRouter);

export default new VueRouter({
  mode: 'history',
  routes: [
    {path: '/', name: 'Home', component: Home},
    {path: '/sources', name: 'Sources', component: Sources},
    {path: '/sources/:source_name', name: 'Source', 'component': Source, props: true},
    {path: '/destinations', name: 'Destinations', component: Destinations},
    {path: '/destinations/:destination_name', name: 'Destination', component: Destination, props: true},
    {path: '/config', name: 'Config', component: Config},
  ]
})
