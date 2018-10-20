
import axios from 'axios'

export function ServiceApiV1LoginRenewPost () {
  return new Promise((resolve, reject) => {
    const requestData = {
      __isRenewRequest: true,
      method: 'POST',
      url: '/service/api/v1/login/renew',
      headers: {
        'Content-Type': 'application/json'
      },
      data: {},
      validateStatus (status) {
        const validStatusList = [200]
        return validStatusList.indexOf(status) !== -1
      }
    }

    axios(requestData)
      .then(response => {
        response.data._status = response.status
        resolve(response.data)
      })
      .catch(error => {
        reject(error)
      })
  })
}
