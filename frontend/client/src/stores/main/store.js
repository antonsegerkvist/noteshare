import Vue from 'vue'
import Vuex from 'vuex'
import account from '@/modules/account'
import file from '@/modules/file'
import user from '@/modules/user'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    account,
    file,
    user
  }
})
