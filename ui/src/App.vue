<template>
  <component v-bind:is="component">
  </component>
</template>

<script>
import { mapState } from 'vuex'
import SecretTable from './components/SecretTable.vue'
import SecretForm from './components/SecretForm.vue'

export default {
  name: 'App',
  components: {
    SecretTable,
    SecretForm,
  },
  computed: mapState({
    selectedARN: state => state.secrets.selectedARN,
    component() {
      if(!this.selectedARN) return "SecretTable";
      else return "SecretForm";
    },
  }),
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
