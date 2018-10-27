<template>
  <div class="view-login">
    <div class="form">
      <h1 class="title">NoteShare</h1>
      <form @submit.prevent="handleFormSubmit" class="login-form">
        <input v-model="email" type="text" name="email" :placeholder="getEmailPlaceholder">
        <input v-model="password" type="password" name="password" :placeholder="getPasswordPlaceholder">
        <input type="submit" :value="getSubmitPlaceholder">
      </form>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'
import { ServiceApiV1LoginPost } from '@/service/api/v1/login/post'

const loginFailReasons = {
  shortEmail: 0,
  longEmail: 1,
  shortPassword: 2,
  userNotFound: 3,
  unknown: 4
}

export default Vue.extend({

  data () {
    return {
      email: '',
      password: '',
      loginFailed: false,
      loginFailedReason: loginFailReasons.unknown
    }
  },

  methods: {

    handleFormSubmit () {
      const body = {
        email: this.email,
        password: this.password
      }
      ServiceApiV1LoginPost(body)
        .then((response) => {
          const lang = this.$route.params.lang
          switch (response._status) {
            case 200:
              this.$router.push({ path: `/${lang}` })
          }
        })
        .catch(() => {
          this.loginFailed = true
          this.loginFailedReason = loginFailReasons.unknown
        })
    }

  },

  computed: {

    getEmailPlaceholder () {
      const lang = this.$route.params.lang
      switch (lang) {
        case 'se':
        case 'en':
        default:
          return 'Email'
      }
    },

    getPasswordPlaceholder () {
      const lang = this.$route.params.lang
      switch (lang) {
        case 'se':
          return 'Lösenord'
        case 'en':
        default:
          return 'Password'
      }
    },
    getSubmitPlaceholder () {
      const lang = this.$route.params.lang
      switch (lang) {
        case 'se':
          return 'Logga in'
        case 'en':
        default:
          return 'Login'
      }
    }
  }
})
</script>

<style lang="scss" scoped>
.view-login {
  height: 100%;
  position: relative;
  width: 100%;

  & > .form {
    display: inline-block;
    left: 50%;
    min-height: 1px;
    max-width: 300px;
    position: absolute;
    text-align: center;
    top: 50%;
    transform: translate(-50%, -50%);
    width: 100%;

    & > .title {
      font-size: 20px;
    }

    & > .login-form {
      width: 100%;

      & > input[type="text"], input[type="password"] {
        appearance: none;
        border: 1px solid #aaa;
        box-sizing: border-box;
        font-size: 14px;
        height: 40px;
        margin-top: 15px;
        outline: none;
        padding: 0 5px;
        width: 100%;
      }

      & > input[type="submit"] {
        appearance: none;
        background-color: #eee;
        border: 1px solid #aaa;
        box-sizing: border-box;
        cursor: pointer;
        font-size: 14px;
        height: 40px;
        margin-top: 15px;
        outline: none;
        padding: 0 5px;
        width: 100%;
      }

    }

  }

}
</style>
