import Vue from 'vue'
import Editor from '@/apps/Editor.vue'
import '@/components/_global'
import '@/components/home'

Vue.config.productionTip = false

new Vue({
  render: h => h(Editor)
}).$mount('#app')
