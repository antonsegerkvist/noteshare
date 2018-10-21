import Vue from 'vue'
import Router from 'vue-router'

import ViewDefault from '@/views/Default.vue'
import ViewLogin from '@/views/Login.vue'
import ViewSubviewFile from '@/views/subviews/File.vue'
import ViewSubviewHome from '@/views/subviews/Home.vue'
import ViewSubviewMe from '@/views/subviews/Me.vue'
import ViewSubviewShop from '@/views/subviews/Shop.vue'

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
          name: 'home',
          component: ViewSubviewHome
        },
        {
          path: 'me',
          name: 'me',
          component: ViewSubviewMe
        },
        {
          path: 'file',
          name: 'file',
          component: ViewSubviewFile
        },
        {
          path: 'shop',
          name: 'shop',
          component: ViewSubviewShop
        }
      ]
    }
  ]
})
