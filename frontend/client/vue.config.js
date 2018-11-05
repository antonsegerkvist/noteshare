module.exports = {

  baseUrl: process.env.NODE_ENV === 'production' ? '/frontend' : '',

  devServer: { proxy: 'http://localhost:8081' },

  pages: {
    
    index: {

      entry: 'src/entries/main.js',

      template: 'public/index.html',

      filename: 'index.html',

      title: 'NoteShare'
    
    },

    editor: {

      entry: 'src/entries/editor.js',

      template: 'public/editor.html',

      filename: 'editor.html',

      title: 'NoteShare - Editor'

    }

  }
}