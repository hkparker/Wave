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
   }
}
}
</script>

<style>
</style>
