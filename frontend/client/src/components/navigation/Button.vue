<template>
  <div :class="getClass" @click="() => { $emit('click') }">
    <div class="icon">
      <i class="material-icons">{{ icon }}</i>
    </div>
    <div class="title">
      <p :title="title">{{ title }}</p>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  props: {

    isActive: {
      type: Boolean,
      default () { return false }
    },

    title: {
      type: String,
      default () { return '' }
    },

    icon: {
      type: String,
      default () { return '' }
    }

  },

  computed: {

    getClass () {
      return this.isActive ? 'navigation-button active' : 'navigation-button'
    }

  }

})
</script>

<style lang="scss" scoped>
.navigation-button {
  background-color: #444;
  cursor: pointer;
  height: 45px;
  left: 0;
  position: absolute;
  width: 100%;

  & > .icon {
    bottom: 0;
    left: 0;
    position: absolute;
    top: 0;
    width: 45px;

    & > i {
      color: #eee;
      font-size: 30px;
      left: 50%;
      position: absolute;
      top: 50%;
      transform: translate(-50%, -50%);
    }

  }

  & > .title {
    box-sizing: border-box;
    height: 100%;
    line-height: 45px;
    padding-left: 45px;
    width: 100%;

    & > p {
      color: #eee;
      font-family: Arial, Helvetica, sans-serif;
      font-size: 14px;
      overflow: hidden;
      text-overflow: ellipsis;
      width: 100%;
      white-space: nowrap;
    }

  }

}
.active > .title > p,
.active > .icon > i { color: #49aad0; }
</style>
