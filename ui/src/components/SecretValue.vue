<template>
  <div>
    <ErrorNotice v-if="isError">{{ errMsg }}</ErrorNotice>
    <template>
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
        <b-form-textarea v-model="secretString" rows="3" max-rows="20"></b-form-textarea>
        <b-button v-on:click="saveSecretValue" variant="primary" v-bind:disabled="saving">Save</b-button>
      </b-card-text>

      <b-card-text v-if="selectedType === 'binary'">
        <p v-if="value.SecretBinary">
          <a v-bind:download="secretName" v-bind:href="'data:application/octet-stream;base64,' + value.SecretBinary">Download '{{secretName}}'</a>
        </p>
        <div>
          <input type="file" id="file" ref="file" v-on:change="handleFileUpload()"/>
          <b-button v-on:click="uploadSecretValue" variant="primary" v-bind:disabled="saving">Upload</b-button>
        </div>
      </b-card-text>

    </template>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import ErrorNotice from './ErrorNotice.vue'

export default {
  props: ['secretName'],
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
      fileUpload: '',
      isError: false,
      errMsg: '',
      saving: false,
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
    serverHost: function() {
      let serverHost = "";
      if(process.env.VUE_APP_SERVER_HOST) {
        serverHost = process.env.VUE_APP_SERVER_HOST;
      }
      return serverHost;
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
    saveSecretValue() {
      this.isError = false;
      this.errMsg = '';
      this.saving = true;

      const baseURI =  `${this.serverHost}/api/secrets/value/update`;
      window.axios.post(baseURI, {
        SecretId: this.value.ARN,
        SecretBinary: this.value.SecretBinary,
        SecretString: this.value.SecretString,
      }).then((result) => {
        this.value = result.data
        this.selectedType = (this.value?.SecretBinary !== null && this.value?.SecretBinary !== undefined) ? 'binary' : 'string';
        this.isError = false;
        this.errMsg = '';
        this.toast('Update successful', 'success');
      })
      .catch(error => {
        this.isError = true;
        this.errMsg = `Save secret value error: ${error?.response?.data?.message ? error.response.data.message : error.message}`;
        console.log(error);
      }).finally(() => {
        this.saving = false;
      });
    },
    loadSecretValue(arn) {
      const baseURI =  `${this.serverHost}/api/secrets/value?arn=${arn}`;
      window.axios.get(baseURI).then((result) => {
        this.value = result.data
        this.selectedType = (this.value?.SecretBinary !== null && this.value?.SecretBinary !== undefined) ? 'binary' : 'string';
        this.isError = false;
        this.errMsg = '';
      })
      .catch(error => {
        this.isError = true;
        this.errMsg = `Load secret value error: ${error?.response?.data?.message ? error.response.data.message : error.message}`
        console.log(error);
      });
    },
    toast(message, variant = null) {
      this.$bvToast.toast(message, {
        title: 'Update',
        autoHideDelay: 5000,
        appendToast: true, 
        variant: variant,
      })
    },
    handleFileUpload() {
      this.fileUpload = this.$refs.file.files[0];
    },
    uploadSecretValue() {
      if(!this.fileUpload) {
        this.isError = true;
        this.errMsg = 'Please choose file';
        return;
      } 

      this.isError = false;
      this.errMsg = '';
      this.saving = true;

      let formData = new FormData();
      formData.append('file', this.fileUpload);
      const baseURI =  `${this.serverHost}/api/secrets/value/upload?arn=${this.selectedARN}`;
      window.axios.post( baseURI, formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          },
        },
      ).then(result => {
        this.value = result.data
        this.selectedType = (this.value?.SecretBinary !== null && this.value?.SecretBinary !== undefined) ? 'binary' : 'string';
        this.isError = false;
        this.errMsg = '';
        this.fileUpload = '';
        this.toast('Update successful', 'success');
      }).catch(error => {
        this.isError = true;
        this.errMsg = `upload secret error: ${error?.response?.data?.message ? error.response.data.message : error.message}`
        console.log(error);
      }).finally(() => {
        this.saving = false;
      });
    },
  },
} </script>
