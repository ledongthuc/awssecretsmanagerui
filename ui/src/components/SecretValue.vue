<template>
  <div>
    <ErrorNotice v-if="isError">{{ errMsg }}</ErrorNotice>
    <template v-if="!isError">
      <b-card-text>
        <b-form-group label="Type" class-label="font-weight-bold" v-slot="{ ariaDescribedby }">
          <b-form-radio-group
            v-model="selectedType"
            :options="typeOptions"
            :aria-describedby="ariaDescribedby"
            name="radio-inline"
          ></b-form-radio-group>
        </b-form-group>
      </b-card-text>

      <b-card-text v-if="selectedType === 'string'">
        <b-form-textarea v-model="secretString" rows="3" max-rows="10"></b-form-textarea>
      </b-card-text>

      <b-button href="#" variant="primary">Save</b-button>
    </template>
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
      selectedType: 'string',
      typeOptions: [
        { text: 'String', value: 'string' },
        { text: 'Binary', value: 'binary' },
      ],
      value: {},
      isError: false,
      errorMsg: '',
    }
  },
  computed: {
    secretString: {
      get: function() {
        return this.beautifyJSON(this.value.SecretString)
      },
      set: function(v) {
        this.value.SecretString = v;
      },
    },
    ...mapState({
      selectedARN: state => state.secrets.selectedARN
    }),
  },
  mounted() {
    this.loadSecretValue(this.selectedARN);
  }, 
  methods: {
    beautifyJSON(raw) {
      try {
        return JSON.stringify(JSON.parse(raw), null, 2);
      } catch(e) {
        return raw;
      }
    },
    loadSecretValue(arn) {
      let serverHost = "";
      if(process.env.VUE_APP_SERVER_HOST) {
        serverHost = process.env.VUE_APP_SERVER_HOST;
      }
      const baseURI =  `${serverHost}/api/secrets/value?arn=${arn}`;
      window.axios.get(baseURI).then((result) => {
        this.value = result.data
        this.selectedType = (this.value?.SecretBinary !== null && this.value?.SecretBinary !== undefined) ? 'binary' : 'string';
        this.isError = false;
        this.errMsg = '';
      })
      .catch(error => {
        this.isError = true;
        this.errMsg = `Load secret value error: ${error.message}`
        console.log(error);
      });
    },
  },
} </script>
