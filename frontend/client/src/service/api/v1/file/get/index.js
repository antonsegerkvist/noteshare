import axios from 'axios'

export function ServiceApiV1FileGet (folderID) {
  return new Promise((resolve, reject) => {
    axios({
      method: 'POST',
      url: `/service/api/v1/file/${folderID}`,
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
