<template>
  <div id="app">
    <LoggedIn v-if="this.$store.getters.loggedIn" />
    <router-view></router-view>
  </div>
</template>

<script>
import LoggedIn from './LoggedIn.vue'

export default {
  name: 'app',
  components: {
    LoggedIn,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
