import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import GlobalEditorTinyMCE from '@/components/_global/editor/TinyMCE.vue'

Vue.config.productionTip = false

Vue.component('global-editor-tinymce', GlobalEditorTinyMCE)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
