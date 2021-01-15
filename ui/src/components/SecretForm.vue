<template>
  <div>
    <b-breadcrumb :items="breadcrumb"></b-breadcrumb>
    <b-container fluid>
      <b-row>
        <b-col></b-col>
        <b-col cols="8">
          <b-card no-body>
            <b-tabs card>
              <b-tab title="Info" active>
                <ErrorNotice v-if="isError">{{ errMsg }}</ErrorNotice>

                <template v-if="!isError">
                  <b-card-text>
                    <div><label class="font-weight-bold">ARN</label></div>
                    <div>{{ data.ARN }}</div>
                  </b-card-text>

                  <b-card-text>
                    <div><label class="font-weight-bold">Description</label></div>
                    <div><b-form-textarea id="textarea-description" v-model="data.Description" rows="1" max-rows="10" plaintext readonly></b-form-textarea></div>
                  </b-card-text>

                  <b-card-text>
                    <div><label class="font-weight-bold">Owning service</label></div>
                    <div>{{ data.OwningService }}</div>
                  </b-card-text>

                  <b-card-text>
                    <div><label class="font-weight-bold">Accessed at</label></div>
                    <div>{{ data.LastAccessedDate }}</div>
                  </b-card-text>

                  <b-card-text>
                    <div><label class="font-weight-bold">Changed at</label></div>
                    <div>{{ data.LastChangedDate }}</div>
                  </b-card-text>

                  <b-card-text>
                    <div><label class="font-weight-bold">Rotation enabled:</label> {{ data.RotationEnabled ? '✓' : 'no' }}</div>
                    <div v-if="data.RotationEnabled">Last rotation at: {{ data.LastRotatedDate }}</div>
                  </b-card-text>

                  <b-card-text v-if="data.DeletedDate">
                    <div><label class="font-weight-bold">Deleted at</label></div>
                    <div>{{ data.DeletedDate }}</div>
                  </b-card-text>

                  <b-card-text>
                    <div><label class="font-weight-bold">Tags</label></div>
                    <b-table striped hover small :items="data.Tags"></b-table>
                  </b-card-text>

                  <b-button href="#" variant="primary">Back</b-button>
                </template>
              </b-tab>
              <b-tab title="Data">
                <SecretValue :secretName="data.Name"/>
              </b-tab>
            </b-tabs>
          </b-card>

        </b-col>
        <b-col></b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import ErrorNotice from './ErrorNotice.vue'
import SecretValue from './SecretValue.vue'

export default {
  components: {
    ErrorNotice,
    SecretValue,
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
        this.errMsg = `Load secret info error: ${error.message}`
        console.log(error);
        window.setTimeout(() => this.loadSecret(arn), 10000);
      });
    },
  },
} </script>
