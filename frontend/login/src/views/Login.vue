<template>
  <div class="view-login">
    <div class="left">
      <ul class="features">
        <li class="feature"><i class="material-icons">create</i>{{ getCreateDocuments }}</li>
        <li class="feature"><i class="material-icons">storage</i>{{ getStoreDocuments }}</li>
        <li class="feature"><i class="material-icons">share</i>{{ getShareDocuments }}</li>
      </ul>
    </div>
    <div class="right">
      <div class="form">
        <h1 class="title">NoteShare</h1>
        <form @submit.prevent="handleFormSubmit" class="login-form">
          <span v-if="loginFailed && loginFailedReason === 0" class="error">
            {{ getShortEmailError }}
          </span>
          <span v-else-if="loginFailed && loginFailedReason === 1" class="error">
            {{ getLongEmailError }}
          </span>
          <span v-else-if="loginFailed && loginFailedReason === 2" class="error">
            {{ getShortPasswordError }}
          </span>
          <span v-else-if="loginFailed && loginFailedReason === 3" class="error">
            {{ getBadCredentialsError }}
          </span>
          <span v-else-if="loginFailed && loginFailedReason === 4" class="error">
            {{ getUserNotFoundError }}
          </span>
          <span v-else-if="loginFailed && loginFailedReason === 5" class="error">
            {{ getUnknownError }}
          </span>
          <input v-model="email" type="text" name="email" :placeholder="getEmailPlaceholder">
          <input v-model="password" type="password" name="password" :placeholder="getPasswordPlaceholder">
          <input type="submit" :value="getSubmitPlaceholder">
        </form>
      </div>
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
  badCredentials: 3,
  userNotFound: 4,
  unknown: 5
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
      if (this.email.length < 3) {
        this.loginFailed = true
        this.loginFailedReason = loginFailReasons.shortEmail
        return
      } else if (this.email.length > 320) {
        this.loginFailed = true
        this.loginFailedReason = loginFailReasons.longEmail
        return
      } else if (this.password.length < 5) {
        this.loginFailed = true
        this.loginFailedReason = loginFailReasons.shortPassword
        return
      }
      const body = {
        email: this.email,
        password: this.password
      }
      ServiceApiV1LoginPost(body)
        .then((response) => {
          const lang = this.$route.params.lang
          switch (response._status) {
            case 200:
              window.location.href = (process.env.NODE_ENV === 'production' ? `/frontend/#/${lang}` : `//localhost:8083/#/${lang}`)
              break
            case 404:
              this.loginFailed = true
              this.loginFailedReason = loginFailReasons.userNotFound
              break
            case 412:
              this.loginFailed = true
              this.loginFailedReason = loginFailReasons.badCredentials
              break
            default:
              this.loginFailed = true
              this.loginFailedReason = loginFailReasons.unknown
          }
        })
        .catch(() => {
          this.loginFailed = true
          this.loginFailedReason = loginFailReasons.unknown
        })
    }

  },

  computed: {

    getCreateDocuments () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Skapa dokument'
        case 'en':
        default:
          return 'Create documents'
      }
    },

    getStoreDocuments () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Lagra dokument'
        case 'en':
        default:
          return 'Store documents'
      }
    },

    getShareDocuments () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Dela dokument'
        case 'en':
        default:
          return 'Share documents'
      }
    },

    getShortEmailError () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Email för kort'
        case 'en':
        default:
          return 'Email too short'
      }
    },

    getLongEmailError () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Email för lång'
        case 'en':
        default:
          return 'Email too long'
      }
    },

    getShortPasswordError () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Lösenord för kort'
        case 'en':
        default:
          return 'Password too short'
      }
    },

    getBadCredentialsError () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Email eller lösenordet har inte rätt längd'
        case 'en':
        default:
          return 'Email or password has a bad length'
      }
    },

    getUserNotFoundError () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Fel användarnamn eller lösenord'
        case 'en':
        default:
          return 'Wrong username or password'
      }
    },

    getUnknownError () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Något oväntat gick fel'
        case 'en':
        default:
          return 'Something unexpected happened'
      }
    },

    getEmailPlaceholder () {
      switch (this.$route.params.lang) {
        case 'se':
        case 'en':
        default:
          return 'Email'
      }
    },

    getPasswordPlaceholder () {
      switch (this.$route.params.lang) {
        case 'se':
          return 'Lösenord'
        case 'en':
        default:
          return 'Password'
      }
    },
    getSubmitPlaceholder () {
      switch (this.$route.params.lang) {
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

  & > .left {
    background-color: #0081c9;
    bottom: 0;
    left: 0;
    position: absolute;
    top: 0;
    width: 50%;

    @media (max-width: 758px) {
      display: none;
    }

    & > .features {
      left: 50%;
      position: absolute;
      top: 50%;
      transform: translate(-50%, -50%);

      & > .feature {
        box-sizing: border-box;
        color: #fff;
        height: 60px;
        line-height: 60px;
        font-size: 20px;
        white-space: nowrap;

        & > i {
          font-size: 30px;
          margin-right: 10px;
          position: relative;
          top: 7px;
        }

      }

    }

  }

  & > .right {
    bottom: 0;
    position: absolute;
    right: 0;
    top: 0;

    @media (max-width: 758px) {
      left: 0;
    }

    @media (min-width: 759px) {
      width: 50%;
    }

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

        & > .error {
          background-color: #f8d7da;
          border-radius: .25rem;
          box-sizing: border-box;
          color: #721c24;
          display: inline-block;
          margin-top: 10px;
          padding: 10px;
          text-align: left;
          width: 100%;
        }

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

}
</style>
