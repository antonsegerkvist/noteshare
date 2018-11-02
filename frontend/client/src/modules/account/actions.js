import { ServiceApiV1AccountLayoutGet } from '@/service/api/v1/account/layout/get'

const actions = {

  fetchAccountLayout ({ state }) {
    ServiceApiV1AccountLayoutGet()
      .then(response => { console.log(response) })
      .catch(() => {})
  }

}

export default actions
