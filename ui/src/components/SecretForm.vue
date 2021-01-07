<template>
  <div>
    <ErrorNotice v-if="isError">{{ errMsg }}</ErrorNotice>
    <div v-elseif>{{ data.Name }}</div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import ErrorNotice from './ErrorNotice.vue'

export default {
  components: {
    ErrorNotice,
  },
  data () {
    return {
      data: {},
      isError: false,
      errorMsg: '',
    }
  },
  computed: mapState({
    selectedARN: state => state.secrets.selectedARN
  }),
  mounted() {
    this.loadSecret(this.selectedARN);
  }, 
  methods: {
    loadSecret(arn) {
      let serverHost = "";
      if(process.env.VUE_APP_SERVER_HOST) {
        serverHost = process.env.VUE_APP_SERVER_HOST;
      }
      const baseURI =  `${serverHost}/api/secrets/detail?arn=${arn}`;
      window.axios.get(baseURI).then((result) => {
        this.data = result.data
        this.isError = false;
        this.errMsg = '';
      })
      .catch(error => {
        this.isError = true;
        this.errMsg = `Load secrets error: ${error.message}`
        console.log(error);
        window.setTimeout(() => this.loadSecret(arn), 10000);
      });
    },
  },
} </script>
