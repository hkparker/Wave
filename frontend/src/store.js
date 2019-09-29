import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import router from "./router.js"

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    authenticationState: "unauthenticated",
    loggedIn: true,
    version: 1,
    currentUser: {}
  },
  getters: {
    loggedIn: state => state.loggedIn,
    authenticationState: state => state.authenticationState,
    version: state => state.version,
    currentUser: state => state.currentUser,
    admin: state => state.currentUser.admin
  },
  actions: {
    authenticate ({commit}, credentials) {
     return axios({url: '/sessions/create', data: credentials, method: 'POST', crossdomain: true, withCredentials: true })
      .then((resp) => {
        commit("setCurrentUser", resp.data)
        commit('authSuccess', credentials.username)
      })
      .catch(err => {
        commit('authFailed', err)
      })
    },
    logout ({commit}) {
     return axios({url: '/sessions/destroy', method: 'POST', crossdomain: true, withCredentials: true })
      .then(() => {
        commit('logout')
      })
    },
    settings ({commit}) {
      commit("settings")
    },
    dashboard ({commit}) {
      commit("dashboard")
    },
    userManager ({commit}) {
      commit("userManager")
    },
    collectorManager ({commit}) {
      commit("collectorManager")
    },
    ids ({commit}) {
      commit("ids")
    },
    rules ({commit}) {
      commit("rules")
    },
    visualization ({commit}) {
      commit("visualization")
    },
    setEnvironment ({commit}) {
      axios({url: '/status', method: 'GET', crossdomain: true, withCredentials: true })
        .then((resp) => {
          commit("setCurrentUser", resp.data)
        })
        .catch(() => {
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
    setCurrentUser: (state, user) => {
      state.currentUser = user
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
    },
    settings () {
      router.push("/settings")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    },
    dashboard () {
      router.push("/")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    },
    userManager () {
      router.push("/users")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    },
    collectorManager () {
      router.push("/collectors")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    },
    ids () {
      router.push("/ids")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    },
    rules () {
      router.push("/rules")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    },
    visualization () {
      router.push("/visualization")
        .catch((e) => {
          if (e.Name == "NavigationDuplicated") {
            // user clicked the page they were already on, no need to log eror
          }
        })
    }
  }
})
