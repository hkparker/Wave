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
    loggedIn: true,
    version: 1,
    currentUser: ""
  },
  getters: {
    loggedIn: state => state.loggedIn,
    authenticationState: state => state.authenticationState,
    version: state => state.version,
    currentUser: state => state.currentUser
  },
  actions: {
    authenticate ({commit}, credentials) {
     return axios({url: 'http://localhost:8081/sessions/create', data: credentials, method: 'POST', crossdomain: true, withCredentials: true })
      .then(() => {
        commit('authSuccess', credentials.username)
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
    }
  },
  mutations: {
    setVersion: (state, newVersion) => {
      state.version = newVersion
    },
    authRequest (state) {
      state.authenticationState = "loading"
    },
    setCurrentUser: (state, username) => {
      state.currentUser = username
    },
    authSuccess (state, username) {
      state.authenticationState = "success"
      state.loggedIn = true
      state.currentUser = username
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
  } else {
    next('/login')
  }
}

const requireLoggedOut = (to, from, next) => {
  if (!store.getters.loggedIn) {
    next()
  } else {
    next('/')
  }
}

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: LoggedIn, beforeEnter: requireLogin },
    { path: '/login', component: Login, beforeEnter: requireLoggedOut }
  ]
})

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
