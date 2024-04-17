import { createRouter, createWebHistory } from 'vue-router';

import Home from './components/Home';
import Sources from './components/Sources';
import Source from './components/Source';
import Destinations from './components/Destinations';
import Destination from './components/Destination';
import Config from './components/Config';

export default createRouter({
  history: createWebHistory(),
  routes: [
    {path: '/', name: 'Home', component: Home},
    {path: '/sources', name: 'Sources', component: Sources},
    {path: '/sources/:source_name', name: 'Source', 'component': Source, props: true},
    {path: '/destinations', name: 'Destinations', component: Destinations},
    {path: '/destinations/:destination_name', name: 'Destination', component: Destination, props: true},
    {path: '/config', name: 'Config', component: Config},
  ]
})
