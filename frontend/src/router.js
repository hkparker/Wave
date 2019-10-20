import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from './Login.vue'
import Dashboard from './components/Dashboard.vue'
import AccountSettings from './components/AccountSettings.vue'
import UserManager from './components/UserManager.vue'
import CollectorManager from './components/CollectorManager.vue'
import IDS from './components/IDS.vue'
import Rules from './components/Rules.vue'
import Visualization from './components/Visualization.vue'
import Unauthorized from './components/Unauthorized.vue'
import store from "./store.js"

Vue.use(VueRouter)

const requireLogin = (to, from, next) => {
  if (store.getters.loggedIn) {
    next()
  } else {
    next('/login')
  }
}

const requireLoginAdmin = (to, from, next) => {
  if (store.getters.loggedIn && store.getters.admin) {
    next()
  } else {
    next('/unauthorized')
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
    { path: '/users', component: UserManager, beforeEnter: requireLogin }, // require admin once works on refresh
    { path: '/collectors', component: CollectorManager, beforeEnter: requireLogin }, // require admin once works on refresh
    { path: '/login', component: Login, beforeEnter: requireLoggedOut },
    { path: '/ids', component: IDS, beforeEnter: requireLogin },
    { path: '/rules', component: Rules, beforeEnter: requireLogin },
    { path: '/visualization', component: Visualization, beforeEnter: requireLogin },
    { path: '/unauthorized', component: Unauthorized, beforeEnter: requireLogin },
  ]
})
