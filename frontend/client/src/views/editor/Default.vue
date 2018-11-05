<template>
  <div class="default">
    <div class="header">
      <router-link class="button visual" :to="`/${$route.params.lang}`">
        {{ getVisualTitle }}
      </router-link>
      <router-link class="button html" :to="`/${$route.params.lang}/html`">
        {{ getHTMLTitle }}
      </router-link>
    </div>
    <div class="body">
      <router-view :key="$route.fullPath"/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  computed: {

    getVisualTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'VISUEL'
        case 'en':
        default:
          return 'VISUAL'
      }
    },

    getHTMLTitle () {
      switch (this.$route.params.lang) {
        case 'se':
        case 'en':
        default:
          return 'HTML'
      }
    }

  }

})
</script>

<style lang="scss" scoped>
.default {
  height: 100%;
  position: relative;
  width: 100%;

  & > .header {
    height: 30px;
    left: 0;
    position: absolute;
    right: 0;
    top: 0;
    z-index: 1;

    & > .button {
      background-color: #eee;
      border-bottom: 1px solid #ccc;
      border-left: 1px solid #ccc;
      border-right: 1px solid #ccc;
      border-top: 1px solid #ccc;
      bottom: 0;
      color: #333;
      line-height: 30px;
      font-family: Arial, Helvetica, sans-serif;
      font-size: 13px;
      position: absolute;
      text-align: center;
      text-decoration: none;
      top: 0;
      width: 60px;
    }

    & > .visual { right: 65px; }
    & > .html { right: 0; }

    & > .router-link-exact-active {
      background-color: #f5f5f5;
      border-bottom: none;
    }

  }

  & > .body {
    border-top: 1px solid #ccc;
    bottom: 0;
    left: 0;
    position: absolute;
    right: 0;
    top: 29px;
    z-index: 0;
  }

}
</style>
