<template>
  <div class="home-layout">
    <div v-for="(child, index) in getSafeData" :key="index" class="element">
      <home-sublayout-two-column
        v-if="child.type === 0"
        :edit-mode="editMode"
        :data="child"/>
      <home-sublayout-three-column
        v-else-if="child.type === 1"
        :edit-mode="editMode"
        :data="child"/>
      <home-object-text
        v-else-if="child.type === 10000"
        :edit-mode="editMode"
        :data="child"/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'

export default Vue.extend({

  name: 'home-layout',

  props: {

    editMode: {
      type: Boolean,
      default () { return false }
    },

    data: {
      type: Array,
      default () { return [] }
    }

  },

  computed: {

    getSafeData () {
      return this.data && this.data.length ? this.data : []
    }

  }

})
</script>

<style lang="scss" scoped>
.home-layout {
  float: left;
  width: 100%;

  & > .element {
    float: left;
    width: 100%;
  }

}
</style>
