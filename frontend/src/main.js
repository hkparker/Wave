import Vue from 'vue'
import App from './App.vue'
import './../node_modules/bulma/css/bulma.css';
import axios from 'axios'
import store from "./store.js"
import router from "./router.js"

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
  store,
  beforeCreate () {
    var $this = this
    axios({url: 'http://localhost:8081/status', method: 'GET', crossdomain: true, withCredentials: true })
      .then((resp) => {
        $this.$store.commit("setCurrentUser", resp.data.user)
      })
      .catch(() => {
        $this.$store.commit('logout')
      })
  }
}).$mount('#app')
