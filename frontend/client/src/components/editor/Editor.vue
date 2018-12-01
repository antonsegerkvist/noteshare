<template>
  <div class="editor-editor">
    <div
      :style="{
        height: getRulerAreaHeight + 'px',
        'min-width': getRulerAreaWidth + 'px'
      }"
      class="ruler-area">
      <div class="top-ruler">
        <div
          :style="{ width: document.metadata.width + 'px' }"
          class="ruler">
        </div>
      </div>
      <div class="left-ruler">
        <div
          :style="{ height: document.metadata.height + 'px' }"
          class="ruler">
        </div>
      </div>
      <div class="corner-ruler">
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
            'margin-top': calculateTop(index) + 'px',
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
const documentSpacing = 10
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
      if (index === 0) {
        return documentBackgroundBottom
      } else {
        return documentSpacing
      }
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
        documentSpacing * Math.max(this.document.pages.length - 1, 0) +
        documentBackgroundBottom
    },

    getRulerAreaWidth () {
      return rulerAreaWidth + this.document.metadata.width
    },

    getDocumentBackgroundHeight () {
      return documentBackgroundTop +
        this.document.metadata.height * this.document.pages.length +
        documentSpacing * Math.max(this.document.pages.length - 1, 0) +
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
  top: 87px;

  & > .ruler-area {
    background-color: rgb(238, 238, 238);
    bottom: 0;
    left: 0;
    position: absolute;
    top: 0;
    width: 100%;

    & > .top-ruler {
      display: inline-block;
      height: 27px;
      left: 27px;
      position: absolute;
      right: 0;
      top: 0;

      & > .ruler {
        background-color: #fff;
        height: 20px;
        left: 50%;
        position: absolute;
        top: 50%;
        transform: translate(-50%, -50%);
      }

    }

    & > .left-ruler {
      bottom: 0;
      display: inline-block;
      left: 0;
      position: absolute;
      top: 27px;
      width: 27px;

      & > .ruler {
        background-color: #fff;
        left: 50%;
        position: absolute;
        top: 50px;
        transform: translateX(-50%);
        width: 20px;
      }

    }

    & > .corner-ruler {
      display: inline-block;
      left: 0;
      height: 27px;
      width: 27px;
      top: 0;
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
        box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
        box-sizing: border-box;
        font-size: 14px;
        margin-left: auto;
        margin-right: auto;
        padding: 20px;

        &:focus {
          outline: 0px solid transparent;
        }

      }

    }

  }

}
</style>
