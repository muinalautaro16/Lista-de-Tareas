// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes = {

    '/default': wrap({ asyncComponent: () => import('./default/Modul.svelte') }),

};

export default routes;
