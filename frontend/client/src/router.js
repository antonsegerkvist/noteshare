import Vue from 'vue'
import Router from 'vue-router'

import ViewDefault from '@/views/Default.vue'
import ViewLogin from '@/views/Login.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/en/login'
    },

    //
    // Login page.
    //

    {
      path: '/:lang/login',
      name: 'login',
      component: ViewLogin
    },

    //
    // Application routes.
    //

    {
      path: '/:lang',
      component: ViewDefault,
      children: [
        {
          path: '',
          name: 'home'
        },
        {
          path: 'me',
          name: 'me'
        },
        {
          path: 'user',
          name: 'user'
        },
        {
          path: 'file',
          name: 'file'
        },
        {
          path: 'shop',
          name: 'shop'
        }
      ]
    }
  ]
})
