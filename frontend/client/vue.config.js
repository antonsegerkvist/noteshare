module.exports = {
  baseUrl: process.env.NODE_ENV === 'production' ? '/frontend' : '',
  devServer: { proxy: 'http://localhost:8081' }
}