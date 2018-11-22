import Vue from 'vue'
import Router from 'vue-router'

import ViewLogin from '@/views/Login.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/en'
    },
    {
      path: '/:lang',
      component: ViewLogin
    }
  ]
})
