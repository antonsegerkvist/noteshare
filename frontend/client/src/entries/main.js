import Vue from 'vue'
import App from '@/apps/App.vue'
import router from '@/routers/main/router'
import store from '@/stores/main/store'
import '@/components/_global'
import '@/components/home'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
