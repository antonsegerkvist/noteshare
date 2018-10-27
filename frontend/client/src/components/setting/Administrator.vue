<template>
  <div ref="background" @click="backgroundClick" class="setting-administrator">
    <div class="window">
      <div class="header">
        <div class="title">
          <h5>{{ getHeaderTitle }}</h5>
        </div>
        <div class="close">
          <i @click="$emit('close')" class="material-icons">
            close
          </i>
        </div>
      </div>
      <div class="body">
        <div class="navigation">
          <div class="entry">
            <div class="icon">
              <i class="material-icons">
                account_circle
              </i>
            </div>
            <div class="title">
              <p>{{ getUsersTitle }}</p>
            </div>
          </div>
          <div class="entry">
            <div class="icon">
              <i class="material-icons">
                group
              </i>
            </div>
            <div class="title">
              <p>{{ getGroupsTitle }}</p>
            </div>
          </div>
          <div class="entry">
            <div class="icon">
              <i class="material-icons">
                bar_chart
              </i>
            </div>
            <div class="title">
              <p>{{ getStatisticsTitle }}</p>
            </div>
          </div>
        </div>
        <div class="container">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  methods: {

    backgroundClick (event) {
      if (event.target === this.$refs.background) {
        this.$emit('close')
      }
    }

  },

  computed: {

    getHeaderTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Administrera'
        case 'en':
        default:
          return 'Administrate'
      }
    },

    getUsersTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Användare'
        case 'en':
        default:
          return 'Users'
      }
    },

    getGroupsTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Grupper'
        case 'en':
        default:
          return 'Groups'
      }
    },

    getStatisticsTitle () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Statistik'
        case 'en':
        default:
          return 'Statistics'
      }
    }

  }

})
</script>

<style lang="scss" scoped>
.setting-administrator {
  background-color: rgba(0,0,0,0.5);
  bottom: 0;
  left: 0;
  position: fixed;
  right: 0;
  top: 0;
  z-index: 1000000;

  & > .window {
    background-color: #eee;
    height: 100%;
    left: 50%;
    max-height: 600px;
    max-width: 700px;
    position: absolute;
    top: 50%;
    transform: translate(-50%,-50%);
    width: 100%;

    & > .header {
      background-color: #333;
      box-sizing: border-box;
      height: 40px;
      left: 0;
      line-height: 40px;
      position: absolute;
      right: 0;
      top: 0;

      & > .title {
        bottom: 0;
        left: 0;
        position: absolute;
        text-align: center;
        top: 0;
        width: 200px;

        & > h5 {
          color: #eee;
          font-size: 18px;
        }

      }

      & > .close {
        bottom: 0;
        position: absolute;
        right: 0;
        top: 0;
        width: 40px;

        & > i {
          color: #eee;
          cursor: pointer;
          font-size: 18px;
          left: 50%;
          position: absolute;
          top: 50%;
          transform: translate(-50%,-50%);
        }

      }

    }

    & > .body {
      bottom: 0;
      left: 0;
      position: absolute;
      right: 0;
      top: 40px;

      & > .navigation {
        background-color: #333;
        bottom: 0;
        left: 0;
        position: absolute;
        top: 0;
        width: 200px;

        & > .entry {
          cursor: pointer;
          height: 30px;
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
            width: 30px;

            & > i {
              color: #eee;
              left: 50%;
              font-size: 20px;
              position: absolute;
              top: 50%;
              transform: translate(-50%,-50%);
            }

          }

          & > .title {
            bottom: 0;
            left: 30px;
            line-height: 30px;
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

      & > .container {
        bottom: 0;
        left: 200px;
        position: absolute;
        right: 0;
        top: 0;
      }

    }

  }

}
</style>
