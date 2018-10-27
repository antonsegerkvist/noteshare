<template>
  <div class="default">
    <div class="header">
      <div class="icon">
        <h5 @click="() => { pushRoute('') }">NoteShare</h5>
      </div>
      <div class="settings">
        <i ref="settings" :title="getSettingsTitle" class="material-icons">
          settings
        </i>
      </div>
      <div ref="settings-dropdown" :class="getSettingsDropdownClass">
        <div class="entry">
          <div class="icon">
            <i class="material-icons">
              account_circle
            </i>
          </div>
          <div class="title">
            <p>{{ getProfileTitle }}</p>
          </div>
        </div>
        <div class="entry">
          <div class="icon">
            <i class="material-icons">
              settings_applications
            </i>
          </div>
          <div class="title">
            <p>{{ getAdministrateTitle }}</p>
          </div>
        </div>
        <div @click="() => { pushRoute('/login') }" class="entry">
          <div class="icon">
            <i class="material-icons">
              exit_to_app
            </i>
          </div>
          <div class="title">
            <p>{{ getLogoutTitle }}</p>
          </div>
        </div>
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

  data () {
    return {
      settingsOpen: false
    }
  },

  mounted () {
    window.addEventListener('click', this.click)
  },

  beforeDestroy () {
    window.removeEventListener('click', this.click)
  },

  methods: {

    pushRoute (path) {
      this.$router.push({ path: `/${this.$route.params.lang}${path}` })
    },

    click (event) {
      if (this.$refs.settings === event.target) {
        this.settingsOpen = !this.settingsOpen
      } else if (this.$refs['settings-dropdown'] === event.target || this.$refs['settings-dropdown'].contains(event.target)) {
        // Do nothing.
      } else {
        this.settingsOpen = false
      }
    }

  },

  computed: {

    getSettingsTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Inställningar'
        case 'en':
        default:
          return 'Settings'
      }
    },

    getSettingsDropdownClass () {
      if (this.settingsOpen) {
        return 'settings-dropdown open'
      } else {
        return 'settings-dropdown'
      }
    },

    getProfileTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Profil'
        case 'en':
        default:
          return 'Profile'
      }
    },

    getAdministrateTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Administrera'
        case 'en':
        default:
          return 'Administrate'
      }
    },

    getLogoutTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Logga ut'
        case 'en':
        default:
          return 'Logout'
      }
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

        &:hover {
          color: #fff;
        }

      }

    }

    & > .settings-dropdown {
      background-color: #333;
      border-bottom-left-radius: 3px;
      height: 0;
      overflow: hidden;
      position: absolute;
      right: 0;
      top: $header-height;
      transition: height 267ms ease-in-out;
      width: 200px;
      z-index: 1000;

      & > .entry {
        cursor: pointer;
        height: 40px;
        position: relative;
        width: 100%;

        &:hover {
          background-color: #444;
        }

        & > .icon {
          bottom: 0;
          left: 0;
          position: absolute;
          top: 0;
          width: 40px;

          & > i {
            color: #eee;
            left: 50%;
            position: absolute;
            top: 50%;
            transform: translate(-50%, -50%);
          }

        }

        & > .title {
          bottom: 0;
          left: 40px;
          line-height: 40px;
          position: absolute;
          right: 0;
          top: 0;

          & > p {
            color: #eee;
            font-family: Arial, Helvetica, sans-serif;
            font-size: 14px;
            font-weight: bold;
          }

        }

      }

    }

    & > .open {
      height: 120px;
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
