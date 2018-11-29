<template>
  <div class="editor-editor">
    <div
      :style="{
        height: getRulerAreaHeight + 'px',
        'min-width': getRulerAreaWidth + 'px'
      }"
      class="ruler-area">
      <div class="top-ruler">
      </div>
      <div class="left-ruler">
      </div>
      <div
        :style="{ height: getDocumentBackgroundHeight + 'px' }"
        class="document-background">
        <div
          v-for="(page, index) in document.pages"
          @keyup="keyup"
          :key="index"
          :style="{
            height: document.metadata.height + 'px',
            top: calculateTop(index) + 'px',
            width: document.metadata.width + 'px',
          }"
          contenteditable="true"
          class="document">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'

const documentBackgroundBottom = 50
const documentBackgroundTop = 50
const rulerAreaHeight = 27
const rulerAreaWidth = 27

export default Vue.extend({

  props: {

    document: {
      type: Object,
      default () {
        return {
          metadata: {
            height: 842,
            width: 595
          },
          pages: [
            {
              content: []
            },
            {
              content: []
            }
          ]
        }
      }
    }

  },

  created () {
    document.execCommand('defaultParagraphSeparator', false, 'p')
  },

  methods: {

    calculateTop (index) {
      return documentBackgroundTop + index * this.document.metadata.height
    },

    keyup (event) {
      console.log(event)
    }

  },

  computed: {

    getRulerAreaHeight () {
      return rulerAreaHeight +
        documentBackgroundTop +
        this.document.metadata.height * this.document.pages.length +
        documentBackgroundBottom
    },

    getRulerAreaWidth () {
      return rulerAreaWidth + this.document.metadata.width
    },

    getDocumentBackgroundHeight () {
      return documentBackgroundTop +
        this.document.metadata.height * this.document.pages.length +
        documentBackgroundBottom
    }

  },

  watch: {

    pages: {
      immediate: true,
      handler (newPages) {
        console.log(newPages)
      }
    }

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
  top: 77px;

  & > .ruler-area {
    background-color: rgb(238, 238, 238);
    bottom: 0;
    left: 0;
    position: absolute;
    top: 0;
    width: 100%;

    & > .top-ruler {
      background-color: red;
      height: 27px;
      left: 27px;
      position: absolute;
      right: 0;
      top: 0;
    }

    & > .left-ruler {
      background-color: blue;
      bottom: 0;
      left: 0;
      position: absolute;
      top: 27px;
      width: 27px;
    }

    & > .document-background {
      background-color: rgb(161, 161, 161);
      border: 1px solid #000;
      box-sizing: border-box;
      left: 27px;
      position: absolute;
      right: 0;
      top: 27px;

      & > .document {
        background-color: #fff;
        border: 1px solid #000;
        box-sizing: border-box;
        font-size: 14px;
        left: 50%;
        padding: 20px;
        position: absolute;
        top: 50px;
        transform: translateX(-50%);

        &:focus {
          box-shadow: none;
          outline: 0px solid transparent;
        }

      }

    }

  }

}
</style>
