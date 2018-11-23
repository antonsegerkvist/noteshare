<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'

import { ServiceApiV1LoginCheckPost } from '@/service/api/v1/login/check/post'
import { ServiceApiV1LoginRenewPost } from '@/service/api/v1/login/renew/post'

export default Vue.extend({

  created () {
    const lang = this.$route.params.lang
    const $router = this.$router
    axios.interceptors.response.use(function (response) {
      return response
    }, function (error) {
      const response = error.response
      if (response.status === 401 && response.config && !response.config.__isRetryRequest && !response.config.__isRenewRequest) {
        return new Promise((resolve, reject) => {
          ServiceApiV1LoginRenewPost()
            .then((renewResponse) => {
              error.config.__isRetryRequest = true
              resolve(axios(error.config))
            })
            .catch((renewError) => {
              window.location = (process.env.NODE_ENV === 'production' ? `/login/#/${lang}` : `//localhost:8082/#/${lang}`)
              reject(renewError)
            })
        })
      } else {
        return Promise.reject(error)
      }
    })
    ServiceApiV1LoginCheckPost()
      .then(() => {})
      .catch(() => {})
  }

})
</script>

<style lang="scss">
html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
b, u, i, center,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td,
article, aside, canvas, details, embed,
figure, figcaption, footer, header, hgroup,
menu, nav, output, ruby, section, summary,
time, mark, audio, video {
  margin: 0;
  padding: 0;
  border: 0;
  font-size: 100%;
  font: inherit;
  vertical-align: baseline;
}
article, aside, details, figcaption, figure,
footer, header, hgroup, menu, nav, section {
  display: block;
}
body, html {
  height: 100%;
  line-height: 1;
  overflow: hidden;
  width: 100%;
}
ol, ul {
  list-style: none;
}
blockquote, q {
  quotes: none;
}
blockquote:before, blockquote:after,
q:before, q:after {
  content: '';
  content: none;
}
table {
  border-collapse: collapse;
  border-spacing: 0;
}
#app {
  display: inline-block;
  height: 100%;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  user-select: none;
  width: 100%;
}
</style>
