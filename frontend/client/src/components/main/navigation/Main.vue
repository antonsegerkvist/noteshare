<template>
  <div class="navigation-main">
    <main-navigation-button
      @click="() => { $router.push({ path: `/${$route.params.lang}/me` }) }"
      :style="getMeButtonStyle"
      :is-active="getActiveNavigationState === 1"
      :title="getMeButtonTitle"
      :icon="'account_circle'"/>
    <div v-if="getActiveNavigationState === 1" class="navigation-me">
      <main-navigation-me/>
    </div>
    <main-navigation-button
      @click="() => { $router.push({ path: `/${$route.params.lang}/file` }) }"
      :style="getFileButtonStyle"
      :is-active="getActiveNavigationState === 2"
      :title="getFileButtonTitle"
      :icon="'cloud'"/>
    <div v-if="getActiveNavigationState === 2" class="navigation-file">
      <main-navigation-file/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import MainNavigationButton from '@/components/main/navigation/Button.vue'
import MainNavigationMe from '@/components/main/navigation/Me.vue'
import MainNavigationFile from '@/components/main/navigation/File.vue'
export default Vue.extend({

  components: {
    'main-navigation-button': MainNavigationButton,
    'main-navigation-me': MainNavigationMe,
    'main-navigation-file': MainNavigationFile
  },

  computed: {

    getActiveNavigationState () {
      if (this.$route.path.startsWith(`/${this.$route.params.lang}/me`)) {
        return 1
      } else if (this.$route.path.startsWith(`/${this.$route.params.lang}/file`)) {
        return 2
      } else {
        return 0
      }
    },

    getMeButtonStyle () {
      return { top: '0px' }
    },

    getFileButtonStyle () {
      if (this.getActiveNavigationState < 2 && this.getActiveNavigationState !== 0) {
        return { bottom: '0px' }
      }
      return { top: '45px' }
    },

    getMeButtonTitle () {
      return 'Username'
    },

    getFileButtonTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Filer'
        case 'en':
        default:
          return 'Files'
      }
    }

  }

})
</script>

<style lang="scss" scoped>
.navigation-main {
  background-color: #333;
  height: 100%;
  position: relative;
  width: 100%;

  & > .navigation-me {
    background-color: #333;
    bottom: 45px;
    left: 0;
    position: absolute;
    right: 0;
    top: 45px;
  }

  & > .navigation-file {
    background-color: #333;
    bottom: 0px;
    left: 0;
    position: absolute;
    right: 0;
    top: 90px;
  }

}
</style>
