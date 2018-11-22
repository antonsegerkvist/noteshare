module.exports = {

  baseUrl: process.env.NODE_ENV === 'production' ? '/login' : '',

  devServer: { proxy: 'http://localhost:8081' }
}