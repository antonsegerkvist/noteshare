import Vue from 'vue'
import Editor from '@/apps/Editor.vue'

Vue.config.productionTip = false

new Vue({
  render: h => h(Editor)
}).$mount('#app')
