import Vue from 'vue'
import router from '@/routers/editor/router.js'
import Editor from '@/apps/Editor.vue'

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(Editor)
}).$mount('#app')
