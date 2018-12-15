<template>
  <div class="navigation-main">
    <navigation-button
      @click="() => { $router.push({ path: `/${$route.params.lang}/me` }) }"
      :style="getMeButtonStyle"
      :is-active="getActiveNavigationState === 1"
      :title="getMeButtonTitle"
      :icon="'account_circle'"/>
    <div v-if="getActiveNavigationState === 1" class="navigation-me">
    </div>
    <navigation-button
      @click="() => { $router.push({ path: `/${$route.params.lang}/file` }) }"
      :style="getFileButtonStyle"
      :is-active="getActiveNavigationState === 2"
      :title="getFileButtonTitle"
      :icon="'cloud'"/>
    <div v-if="getActiveNavigationState === 2" class="navigation-file">
    </div>
    <navigation-button
      @click="() => { $router.push({ path: `/${$route.params.lang}/shop` }) }"
      :style="getShopButtonStyle"
      :is-active="getActiveNavigationState === 3"
      :title="getShopButtonTitle"
      :icon="'shop'"/>
    <div v-if="getActiveNavigationState === 3" class="navigation-shop">
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import NavigationButton from '@/components/main/navigation/Button.vue'
import NavigationMe from '@/components/main/navigation/Me.vue'
import NavigationFile from '@/components/main/navigation/File.vue'
import NavigationShop from '@/components/main/navigation/Shop.vue'
export default Vue.extend({

  components: {
    'navigation-button': NavigationButton,
    'navigation-me': NavigationMe,
    'navigation-file': NavigationFile,
    'navigation-shop': NavigationShop
  },

  computed: {

    getActiveNavigationState () {
      if (this.$route.path.startsWith(`/${this.$route.params.lang}/me`)) {
        return 1
      } else if (this.$route.path.startsWith(`/${this.$route.params.lang}/file`)) {
        return 2
      } else if (this.$route.path.startsWith(`/${this.$route.params.lang}/shop`)) {
        return 3
      } else {
        return 0
      }
    },

    getMeButtonStyle () {
      return { top: '0px' }
    },

    getFileButtonStyle () {
      if (this.getActiveNavigationState < 2 && this.getActiveNavigationState !== 0) {
        return { bottom: '45px' }
      }
      return { top: '45px' }
    },

    getShopButtonStyle () {
      if (this.getActiveNavigationState < 3 && this.getActiveNavigationState !== 0) {
        return { bottom: '0px' }
      }
      return { top: '90px' }
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
    },

    getShopButtonTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Affär'
        case 'en':
        default:
          return 'Shop'
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
    bottom: 90px;
    left: 0;
    position: absolute;
    right: 0;
    top: 45px;
  }

  & > .navigation-file {
    background-color: #333;
    bottom: 45px;
    left: 0;
    position: absolute;
    right: 0;
    top: 90px;
  }

  & > .navigation-shop {
    background-color: #333;
    bottom: 0;
    left: 0;
    position: absolute;
    right: 0;
    top: 135px;
  }

}
</style>
