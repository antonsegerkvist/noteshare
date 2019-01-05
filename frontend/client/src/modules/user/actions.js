import { ServiceApiV1UserMeGet } from '@/service/api/v1/user/me/get'

const actions = {

  fetchUserMe ({ state }) {
    ServiceApiV1UserMeGet()
      .then(response => {
        switch (response._status) {
          case 200:
            delete response._status
            state.me = response
            break
          default:
            state.me = {}
        }
      })
      .catch(() => {
        state.me = {}
      })
  }

}

export default actions
