import { ServiceApiV1UserGet } from '@/service/api/v1/user/get'

const actions = {

  fetchUserMe ({ state }) {
    ServiceApiV1UserGet(`me`)
      .then(response => {
        switch (response._status) {
          case 200:
            delete response._status
            state.user = response
            break
          default:
            state.user = {}
        }
      })
      .catch(() => {
        state.user = {}
      })
  }

}

export default actions
