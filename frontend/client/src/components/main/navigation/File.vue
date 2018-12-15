<template>
  <div class="main-navigation-file">
    <div class="container">
    </div>
    <div class="options">
      <div :title="getAddFolderTitle" class="add-folder">
        <i class="material-icons">add</i>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import { ServiceApiV1FolderGet } from '@/service/api/v1/folder/get'
export default Vue.extend({

  data () {
    return {
      folderTree: {},
      folderTreeError: undefined
    }
  },

  created () {
    ServiceApiV1FolderGet(0)
      .then(response => {
        console.log(response)
        if (response && response.length) {
          let obj = {}
          obj[0] = response
          this.folderTree = obj
        }
        this.folderTreeError = undefined
      })
      .catch(() => {})
  },

  computed: {

    getAddFolderTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Lägg till en mapp i den markerade mappen.'
        case 'en':
        default:
          return 'Add a folder to the marked folder.'
      }
    }

  }

})
</script>

<style lang="scss" scoped>
.main-navigation-file {
  bottom: 0;
  left: 0;
  position: absolute;
  right: 0;
  top: 0;

  & > .container {
    bottom: 25px;
    left: 0;
    position: absolute;
    right: 0;
    top: 0;
  }

  & > .options {
    bottom: 0;
    height: 25px;
    left: 0;
    position: absolute;
    right: 0;

    & > .add-folder {
      bottom: 0;
      cursor: pointer;
      position: absolute;
      right: 0;
      top: 0;
      width: 25px;

      & > i {
        color: #fff;
        left: 50%;
        position: absolute;
        top: 50%;
        transform: translate(-50%, -50%);
      }

    }

  }

}
</style>
