import { ServiceApiV1FileGet } from '@/service/api/v1/file/get'
import { ServiceApiV1FolderGet } from '@/service/api/v1/folder/get'

const actions = {

  fetchFolderChildren ({ state }, folderID) {
    ServiceApiV1FolderGet(folderID)
      .then(response => {})
      .catch(() => {})
  },

  fetchFolderFiles ({ state }, folderID) {
    ServiceApiV1FileGet(folderID)
      .then(response => {})
      .catch(() => {})
  }

}

export default actions
