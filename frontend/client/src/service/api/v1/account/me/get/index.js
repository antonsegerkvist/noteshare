import axios from 'axios'

export function ServiceApiV1AccountMeGet () {
  return new Promise((resolve, reject) => {
    axios({
      method: 'GET',
      url: '/service/api/v1/account/me',
      validateStatus (status) {
        const validStatusCodes = [200]
        return validStatusCodes.indexOf(status) !== -1
      }
    })
      .then(response => {
        response.data._status = response.status
        resolve(response.data)
      })
      .catch(error => {
        reject(error)
      })
  })
}
