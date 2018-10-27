<template>
  <div class="tinymce">
    <form>
      <textarea ref="tinymce">
      </textarea>
    </form>
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({

  mounted () {
    /* global tinymce */
    tinymce.init({
      target: this.$refs.tinymce,
      plugins: [
        'advlist autolink lists link image charmap print preview anchor',
        'searchreplace visualblocks code fullscreen',
        'insertdatetime media table contextmenu paste'
      ],
      menubar: true,
      toolbar: 'undo redo | fontselect | fontsizeselect | bold italic | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link image'
    })
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
