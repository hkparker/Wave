<template>
  <div id="app">
    <Navbar v-if="this.$store.getters.loggedIn" />
    <router-view></router-view>
  </div>
</template>

<script>
import Navbar from './Navbar.vue'

export default {
  name: 'app',
  components: {
    Navbar,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
