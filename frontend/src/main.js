import Vue from 'vue'
import App from './App.vue'
import './../node_modules/bulma/css/bulma.css';
import store from "./store.js"
import router from "./router.js"

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
  store,
  beforeCreate () {
    this.$store.dispatch("setEnvironment")
  }
}).$mount('#app')
