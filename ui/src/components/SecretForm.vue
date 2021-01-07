<template>
  <b-container fluid>
    <b-breadcrumb :items="breadcrumb"></b-breadcrumb>
    <b-row>
      <b-col></b-col>
      <b-col  cols="8">
        <b-form>

          <b-form-group id="input-group-description" label="Description:" label-for="textarea-description">
            <b-form-textarea id="textarea-description" v-model="data.Description" rows="3" max-rows="8"></b-form-textarea>
          </b-form-group>

          <b-form-group id="input-group-owning-service" label="Owning service:" label-for="input-owning-service" v-if="data.OwningService">
            <b-form-input id="input-owning-service" v-model="data.OwningService" plaintext readonly></b-form-input>
          </b-form-group>

          <b-form-group id="input-group-accessed-date" label="Accessed at:" label-for="input-accessed-date">
            <b-form-input id="input-accessed-date" v-bind:value="data.LastAccessedDate" plaintext readonly></b-form-input>
          </b-form-group>

          <b-form-group id="input-group-changed-date" label="Changed at:" label-for="input-changed-date">
            <b-form-input id="input-changed-date" v-bind:value="data.LastChangedDate" plaintext readonly></b-form-input>
          </b-form-group>

          <b-form-group id="input-group-created-date" label="Created at:" label-for="input-created-date">
            <b-form-input id="input-created-date" v-bind:value="data.CreatedDate" plaintext readonly></b-form-input>
          </b-form-group>

          <b-form-group id="input-group-rotation-enabled" label="Rotation enabled:" label-for="input-last-rotation-enabled" :description="lastRotationDescription">
            <b-form-checkbox id="checkbox-rotation-enabled" v-bind:checked="data.RotationEnabled" readonly disabled></b-form-checkbox>
          </b-form-group>

          <b-form-group id="input-group-deleted-date" label="Deleted at:" label-for="input-deleted-date" v-if="data.DeletedDate">
            <b-form-input id="input-deleted-date" v-bind:value="data.DeletedDate" plaintext readonly></b-form-input>
          </b-form-group>


          <b-button type="submit" variant="primary">Submit</b-button>
          <b-button type="reset" variant="danger">Reset</b-button>
        </b-form>
      </b-col>
      <b-col></b-col>
    </b-row>
  </b-container>
</template>

<script>
import { mapState } from 'vuex'
// import ErrorNotice from './ErrorNotice.vue'

export default {
  components: {
    // ErrorNotice,
  },
  data () {
    return {
      breadcrumb: [
        {
          text: 'secrets',
          href: '#'
        },
      ],
      data: {},
      isError: false,
      errorMsg: '',
    }
  },
  computed: {
    lastRotationDescription() {
      if(this.data.LastRotatedDate) {
        return `Last roration: ${this.data.LastRotatedDate}`;
      }
      return '';
    },
    ...mapState({
      selectedARN: state => state.secrets.selectedARN
    }),
  },
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
        this.breadcrumb.push({
          text: result.data.Name,
          active: true,
        });
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
