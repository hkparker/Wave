<template>
  <div>
    <form class="login" @submit.prevent="login">
      <h1>Sign in</h1>
      <label>User name</label>
      <input required v-model="username" type="text" placeholder="Hayden"/>
      <br />
      <label>Password</label>
      <input required v-model="password" type="password" placeholder="Password"/>
      <hr/>
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script>

export default {
  name: 'app',
  data: function() {
    return {
      username: "",
      password: ""
    }
  },
  methods: {
    login: function () {
      const { username, password } = this
      this.$store.dispatch('authenticate', { username, password })
      //.then(() => { this.$router.push('/') })
      //debugger; // go do stuff like kick off a func to check this and either set cookies and variables or pop a warning
      //const { username, password } = this
      //this.$store.dispatch(AUTH_REQUEST, { username, password }).then(() => {
      //  this.$router.push('/')
      //})
   }
}
}
</script>

<style>
</style>
