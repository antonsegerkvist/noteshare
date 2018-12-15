import Vue from 'vue'
import Router from 'vue-router'

import ViewDefault from '@/views/main/Default.vue'
import ViewSubviewFile from '@/views/main/subviews/File.vue'
import ViewSubviewHome from '@/views/main/subviews/Home.vue'
import ViewSubviewMe from '@/views/main/subviews/Me.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/en'
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
        }
      ]
    }
  ]
})
