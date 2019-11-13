module.exports = {
  devServer: {
    proxy: {
      '/streams/': {
        target: 'ws://localhost:8081',
        ws: true,
      },
      '/sessions': {
        target: 'http://localhost:8081',
      },
      '/users': {
        target: 'http://localhost:8081',
      },
      '/status': {
        target: 'http://localhost:8081',
      },
      '/collector': {
        target: 'http://localhost:8081',
      },
      '/collectors': {
        target: 'http://localhost:8081',
      },
      '/tls': {
        target: 'http://localhost:8081',
      },
      '/favicon.ico': {
        target: 'http://localhost:8081'
      },
      '/wave.svg': {
        target: 'http://localhost:8081'
      }
    }
  },
  runtimeCompiler: true
}
