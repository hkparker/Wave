import Vue from 'vue'
import App from './App.vue'
import Login from './Login.vue'
import LoggedIn from './LoggedIn.vue'
import './../node_modules/bulma/css/bulma.css';
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import axios from 'axios'

Vue.use(VueRouter)
Vue.use(Vuex)

Vue.config.productionTip = false

const store = new Vuex.Store({
  state: {
    authenticationState: "unauthenticated",
    loggedIn: false
  },
  getters: {
    loggedIn: state => state.loggedIn,
    authenticationState: state => state.authenticationState,
  },
  actions: {
    authenticate ({commit}, credentials) {
     return axios({url: 'http://localhost:8081/sessions/create', data: credentials, method: 'POST', crossdomain: true })
      .then(() => {
        commit('authSuccess')
        router.push('/')
      })
      .catch(err => {
        commit('authFailed', err)
      })
    },
    logout ({commit}) {
      commit('logout')
    }
  },
  mutations: {
    authRequest (state) {
      state.authenticationState = "loading"
    },
    authSuccess (state) {
      state.authenticationState = "success"
      state.loggedIn = true
    },
    authFailed (state) {
      state.authenticationState = "failed"
    },
    logout (state) {
      state.loggedIn = false
      state.authenticationState = "logged_out"
    }
  }
})

const requireLogin = (to, from, next) => {
  if (store.getters.loggedIn) {
    next()
    return
  }
  next('/login')
}

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: LoggedIn, beforeEnter: requireLogin },
    { path: '/login', component: Login }
  ]
})

new Vue({
  router,
  render: h => h(App),
  store,
}).$mount('#app')
