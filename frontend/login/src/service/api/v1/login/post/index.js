import axios from 'axios'

export function ServiceApiV1LoginPost (bodyData) {
  return new Promise((resolve, reject) => {
    axios({
      method: 'POST',
      url: '/service/api/v1/login',
      headers: {
        'Content-Type': 'application/json'
      },
      data: bodyData,
      validateStatus (status) {
        const validStatusCodes = [200, 404, 412]
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
