<template>
  <div class="tinymce">
    <form @submit.prevent="submit">
      <textarea ref="tinymce">
      </textarea>
    </form>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  data () {
    return {
      content: ''
    }
  },

  mounted () {
    /* global tinymce */
    tinymce.init({
      target: this.$refs.tinymce,
      plugins: [
        'advlist autolink lists link image charmap print preview anchor',
        'searchreplace visualblocks code fullscreen',
        'insertdatetime media table contextmenu paste'
      ],
      menubar: false,
      toolbar: 'undo redo | fontselect | fontsizeselect | bold italic | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link',
      setup: (editor) => {
        editor.on('keyup', () => {
          this.content = editor.getContent()
        })
      }
    })
  },

  methods: {

    submit () {
      console.log(this.content)
      this.$emit('save', this.content)
    }

  }

})
</script>

<style lang="scss" scoped>
.tinymce {
  box-sizing: border-box;
  height: 100%;
  width: 100%;
}
</style>
