<template>
  <div>
    <ErrorNotice v-if="isError">{{ errMsg }}</ErrorNotice>
    <b-breadcrumb :items="[{text: 'secrets',href: '#'}]"></b-breadcrumb>
    <b-container fluid>
      <b-row>
        <b-col></b-col>
        <b-col cols="8">
          <b-card>
              <BootstrapTable :columns="columns" :data="data" :options="options"></BootstrapTable>
          </b-card>
        </b-col>
        <b-col></b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
import ErrorNotice from './ErrorNotice.vue'
import { mapState, mapActions } from 'vuex'

export default {
  components: {
    ErrorNotice,
  },
  data () {
    return {
      columns: [
        {
          field: 'Name',
          title: 'Name',
          sortable: true,
        },
        {
          field: 'LastAccessedDate',
          title: 'Accessed at',
          sortable: true,
        },
        {
          field: 'CreatedDate',
          title: 'Created at',
          sortable: true,
        },
        {
          field: 'LastChangedDate',
          title: 'Updated at',
          sortable: true,
        },
        {
          field: 'RotationEnabled',
          title: 'Rotated',
          sortable: true,
          align: 'center',
        },
        {
          field: 'LastRotatedDate',
          title: 'Rotated at',
          sortable: true,
        },
        {
          field: 'DeletedDate',
          title: 'Deleted at',
          sortable: true,
        },
      ],
      data: [],
      options: {
        search: true,
        searchAlign: true,
        showColumns: true,
        onClickRow: (row) => {
          this.selectARN(row.ARN);
        },
      },
      isError: false,
      errorMsg: '',
    }
  },
  computed: mapState({
    selectedARN: state => state.secrets.selectedARN
  }),
  mounted() {
    this.loadSecrets();
  }, 
  methods: {
    loadSecrets() {
      let serverHost = "";
      if(process.env.VUE_APP_SERVER_HOST) {
        serverHost = process.env.VUE_APP_SERVER_HOST;
      }
      const baseURI =  `${serverHost}/api/secrets`;
      window.axios.get(baseURI).then((result) => {
        this.data = result.data.map(item => {
          item.CreatedDate = item.CreatedDate ? window.moment(item.CreatedDate).format('YYYY/MM/DD hh:mm:ss') : null;
          item.LastAccessedDate = item.LastAccessedDate ? window.moment(item.LastAccessedDate).format('YYYY/MM/DD hh:mm:ss') : null;
          item.LastChangedDate = item.LastChangedDate ? window.moment(item.LastChangedDate).format('YYYY/MM/DD hh:mm:ss') : null;
          item.LastRotatedDate = item.LastRotatedDate ? window.moment(item.LastRotatedDate).format('YYYY/MM/DD hh:mm:ss') : null;
          item.DeletedDate = item.DeletedDate ? window.moment(item.DeletedDate).format('YYYY/MM/DD hh:mm:ss') : null;
          item.RotationEnabled = item.RotationEnabled ? '✓' : '';
          return item;
        })
        this.isError = false;
        this.errMsg = '';
      })
      .catch(error => {
        this.isError = true;
        this.errMsg = `Load secrets error: ${error.message}`
        console.log(error);
        window.setTimeout(this.loadSecrets, 10000);
      });
    },
    ...mapActions('secrets', [
      'selectARN'
    ]),
  },
} </script>
