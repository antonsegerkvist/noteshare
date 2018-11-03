import { ServiceApiV1AccountMeGet } from '@/service/api/v1/account/me/get'

const actions = {

  fetchAccount ({ state }) {
    ServiceApiV1AccountMeGet()
      .then(response => {
        switch (response._status) {
          case 200:
            state.account = response.account
            break
          default:
            state.account = {}
        }
      })
      .catch(() => {
        state.account = {}
      })
  }

}

export default actions
