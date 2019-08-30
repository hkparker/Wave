<template>
  <div class="hello">
    <h1>version: {{ this.$store.getters.currentUser.version }}</h1><br />
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
