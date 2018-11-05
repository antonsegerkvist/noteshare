import Vue from 'vue'
import Router from 'vue-router'

import ViewDefault from '@/views/editor/Default.vue'
import ViewSubviewTinyMCE from '@/views/editor/subviews/TinyMCE.vue'
import ViewSubviewEditHTML from '@/views/editor/subviews/EditHTML.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/en'
    },
    {
      path: '/:lang',
      component: ViewDefault,
      children: [
        {
          path: '',
          name: 'tinymce',
          component: ViewSubviewTinyMCE
        },
        {
          path: 'html',
          name: 'html',
          component: ViewSubviewEditHTML
        }
      ]
    }
  ]
})
