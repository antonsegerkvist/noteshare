import Vue from 'vue'
import Router from 'vue-router'

import ViewDefault from '@/views/main/Default.vue'
import ViewLogin from '@/views/main/Login.vue'
import ViewSubviewFile from '@/views/main/subviews/File.vue'
import ViewSubviewHome from '@/views/main/subviews/Home.vue'
import ViewSubviewMe from '@/views/main/subviews/Me.vue'
import ViewSubviewShop from '@/views/main/subviews/Shop.vue'

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
