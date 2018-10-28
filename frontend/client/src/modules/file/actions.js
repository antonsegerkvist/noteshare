import ServiceApiV1FolderGet from '@/service/api/v1/folder/get'

const actions = {

  fetchFolderChildren ({ state }, folderID) {
    ServiceApiV1FolderGet(folderID)
      .then(response => {})
      .catch(() => {})
  },

  fetchFolderFiles (folderID) {

  }

}

export default actions
