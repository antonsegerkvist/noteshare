<template>
  <div class="view-subview-home">
    <div class="header">
      <h5 class="title">{{ getWelcomeMessage }}</h5>
      <h5 class="date">{{ getFriendlyDate }}</h5>
    </div>
    <div class="body">
      <home-layout
        :edit-mode="false"
        :data="[ { type: 0 }, { type: 1 }, { type: 0 } ]"/>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  computed: {

    getWelcomeMessage () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Välkommen'
        case 'en':
        default:
          return 'Welcome'
      }
    },

    getFriendlyDate () {
      let d = new Date()
      let n = d.getDay()
      let l = d.getDate()
      let m = d.getMonth()
      let y = d.getFullYear()

      d = new Date(Date.UTC(d.getFullYear(), d.getMonth(), d.getDate()))
      let dayNum = d.getUTCDay() || 7
      d.setUTCDate(d.getUTCDate() + 4 - dayNum)
      let yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1))
      let week = Math.ceil((((d - yearStart) / 86400000) + 1) / 7)

      const seDayArray = ['Söndag', 'Måndag', 'Tisdag', 'Onsdag', 'Torsdag', 'Fredag', 'Lördag']
      const seMonthArray = ['Januari', 'Februari', 'Mars', 'April', 'Maj', 'Juni', 'Juli', 'Augusti', 'September', 'Oktober', 'November', 'December']
      const enDayArray = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']
      const enMonthArray = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']

      switch (this.$route.params.lang) {
        case 'en':
          return `${enDayArray[n]} the ${l} ${enMonthArray[m].toLowerCase()} ${y}, week ${week}`
        case 'se':
        default:
          return `${seDayArray[n]} den ${l} ${seMonthArray[m].toLowerCase()} ${y}, vecka ${week}`
      }
    }

  }

})
</script>

<style lang="scss" scoped>
.view-subview-home {
  background-color: #eee;
  height: 100%;
  overflow-x: hidden;
  overflow-y: auto;
  width: 100%;

  & > .header {
    box-sizing: border-box;
    display: block;
    height: 45px;
    line-height: 45px;
    overflow: hidden;
    padding: 0 10px;
    width: 100%;

    & > h5.title {
      float: left;
      font-family: Arial, Helvetica, sans-serif;
      font-size: 24px;
    }

    & > h5.date {
      float: right;
      font-family: Arial, Helvetica, sans-serif;
      font-size: 20px;
    }

  }

  & > .body {
    display: block;
    width: 100%;
  }

}
</style>
