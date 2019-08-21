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
    loggedIn: false,
    version: 1
  },
  getters: {
    loggedIn: state => state.loggedIn,
    authenticationState: state => state.authenticationState,
    version: state => state.version,
  },
  actions: {
    authenticate ({commit}, credentials) {
     return axios({url: 'http://localhost:8081/sessions/create', data: credentials, method: 'POST', crossdomain: true, withCredentials: true })
      .then(() => {
        commit('authSuccess')
      })
      .catch(err => {
        commit('authFailed', err)
      })
    },
    logout ({commit}) {
     return axios({url: 'http://localhost:8081/sessions/destroy', method: 'POST', crossdomain: true, withCredentials: true })
      .then(() => {
        commit('logout')
      })
      .catch(err => {
        console.log(err)
        //commit('authFailed', err)
      })
    }
  },
  mutations: {
    setVersion: (state, newVersion) => {
      state.version = newVersion
    },
    authRequest (state) {
      state.authenticationState = "loading"
    },
    authSuccess (state) {
      state.authenticationState = "success"
      state.loggedIn = true
      router.push('/')
    },
    authFailed (state) {
      state.authenticationState = "failed"
    },
    logout (state) {
      state.loggedIn = false
      state.authenticationState = "logged_out"
      router.push("/login")
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


axios({url: 'http://localhost:8081/version', method: 'GET', crossdomain: true, withCredentials: true })
  .then(() => {
    store.commit('authSuccess')
})

new Vue({
  router,
  render: h => h(App),
  store,
}).$mount('#app')
