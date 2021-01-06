<template>
  <SecretTable msg="Welcome to Your Vue.js App" />
</template>

<script>
import SecretTable from './components/SecretTable.vue'

export default {
  name: 'App',
  components: {
    SecretTable,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
