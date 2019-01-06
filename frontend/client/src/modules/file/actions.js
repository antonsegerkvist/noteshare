import { ServiceApiV1FolderIDGet } from '@/service/api/v1/folder/id/get'

const actions = {

  fetchFolderChildren ({ state }, folderID) {
    ServiceApiV1FolderIDGet(folderID)
      .then(response => {
        state.folders[folderID] = response
      })
      .catch(() => {
        state.folders[folderID] = []
      })
  }

}

export default actions
