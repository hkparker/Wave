import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from './Login.vue'
import Dashboard from './components/Dashboard.vue'
import AccountSettings from './components/AccountSettings.vue'
import store from "./store.js"

Vue.use(VueRouter)

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

export default new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Dashboard, beforeEnter: requireLogin },
    { path: '/settings', component: AccountSettings, beforeEnter: requireLogin },
    { path: '/login', component: Login, beforeEnter: requireLoggedOut }
  ]
})
