<template>
  <b-dropdown v-bind:text="selectedRegion || 'Select region'">
    <b-dropdown-item v-bind:key="region" v-for="region in regions" v-bind:active="selectedRegion === region" @click="$emit('onChange', region)">
      {{ region }}
    </b-dropdown-item>
  </b-dropdown>
</template>

<script>

export default {
  props: ['selectedRegion'],
  data () {
    return {
      regions: [],
    }
  },
  mounted() {
    this.loadRegions();
  }, 
  methods: {
    loadRegions() {
      let serverHost = "";
      if(process.env.VUE_APP_SERVER_HOST) {
        serverHost = process.env.VUE_APP_SERVER_HOST;
      }

      const baseURI =  `${serverHost}/api/aws/regions`;
      window.axios.get(baseURI).then((result) => {
        this.regions = result.data.sort();
      })
      .catch(error => {
        console.log(error);
      });
    },
  },
} </script>
