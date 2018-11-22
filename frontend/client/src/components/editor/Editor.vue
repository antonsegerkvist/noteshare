<template>
  <div class="editor-editor">
    <div class="ruler-area">
      <div class="document-background">
        <div contenteditable="true" class="document">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  props: {

    pages: {
      type: Array,
      default () {
        return []
      }
    }

  },

  created () {
    document.execCommand('defaultParagraphSeparator', false, 'p')
  }

})
</script>

<style lang="scss" scoped>
.editor-editor {
  bottom: 0;
  left: 0;
  overflow: auto;
  position: absolute;
  right: 0;
  top: 127px;

  & > .ruler-area {
    background-color: rgb(238, 238, 238);
    bottom: 0;
    left: 0;
    min-width: 700px;
    position: absolute;
    right: 0;
    top: 0;

    & > .document-background {
      background-color: rgb(161, 161, 161);
      border: 1px solid #000;
      bottom: 0;
      box-sizing: border-box;
      left: 27px;
      min-width: 673px;
      position: absolute;
      right: 0;
      top: 27px;

      & > .document {
        background-color: #fff;
        border: 1px solid #000;
        box-sizing: border-box;
        left: 50%;
        padding: 20px;
        position: absolute;
        top: 50%;
        transform: translate(-50%, -50%);
        height: 800px;
        width: 673px;

        &:focus {
          box-shadow: none;
          outline: 0px solid transparent;
        }

      }

    }

  }

}
</style>
