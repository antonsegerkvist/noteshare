<template>
  <div class="default">
    <div class="header">
      <div class="icon">
        <h5 @click="() => { pushRoute('') }">NoteShare</h5>
      </div>
      <div class="settings">
        <i class="material-icons">
          settings
        </i>
      </div>
    </div>
    <div class="navigation">
      <div class="container">
        <navigation-main/>
      </div>
      <div class="size-box"></div>
    </div>
    <div class="body">
      <router-view :key="$route.fullPath"/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import NavigationMain from '@/components/navigation/Main.vue'
export default Vue.extend({

  components: {
    'navigation-main': NavigationMain
  },

  methods: {

    pushRoute (path) {
      this.$router.push({ path: `/${this.$route.params.lang}${path}` })
    }

  }

})
</script>

<style lang="scss" scoped>
$header-height: 37px;
$navigation-width: 250px;

.default {
  height: 100%;
  position: relative;
  width: 100%;

  & > .header {
    background-color: #333;
    box-sizing: border-box;
    left: 0;
    position: absolute;
    height: $header-height;
    right: 0;
    top: 0;

    & > .icon {
      float: left;
      height: 100%;
      position: relative;
      width: 250px;

      & > h5 {
        color: #eee;
        cursor: pointer;
        font-size: 20px;
        left: 50%;
        position: absolute;
        top: 50%;
        transform: translate(-50%, -50%);
      }

    }

    & > .settings {
      float: right;
      height: 100%;
      position: relative;
      width: 50px;

      & > i {
        color: #eee;
        cursor: pointer;
        font-size: 24px;
        left: 50%;
        position: absolute;
        top: 50%;
        transform: translate(-50%, -50%);
      }

    }

  }

  & > .navigation {
    background-color: #333;
    bottom: 0;
    left: 0;
    position: absolute;
    top: $header-height;
    width: $navigation-width;

    & > .container {
      bottom: 0;
      left: 0;
      position: absolute;
      right: 4px;
      top: 0;
    }

    & > .size-box {
      bottom: 0;
      cursor: col-resize;
      position: absolute;
      right: 0;
      top: 0;
      width: 4px;
    }

  }

  & > .body {
    bottom: 0;
    left: $navigation-width;
    position: absolute;
    right: 0;
    top: $header-height;
  }

}
</style>
