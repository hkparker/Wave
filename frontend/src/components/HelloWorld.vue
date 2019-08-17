<template>
  <div class="hello">
    <h1>msg variable: {{ foo }}</h1>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data: function () {
    return {
      foo: "bar"
    }
  },
  methods: {
    updateFoo: function (str) {
      this.foo = str
    }
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
}
</style>
