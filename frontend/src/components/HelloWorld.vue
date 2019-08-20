<template>
  <div class="hello">
    <h1>version: {{ this.$store.getters.version }}</h1><br />
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  //data: function () {
  //  return {
  //    version: 1,
  //  }
  //},
  //methods: {
  //  updateFoo: function (str) {
  //    this.foo = str
  //  }
  //},
  mounted () {
    var $this = this
    axios({url: 'http://localhost:8081/version', method: 'GET', crossdomain: true, withCredentials: true })
      .then((response) => {
        //$this.$store.version = response.data.version
        $this.$store.commit('setVersion', response.data.version)
      })
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
