<template>
  <div class="container">
    <h1 class="title">Collector Management</h1>
    <h2 class="subtitle">Current Collectors</h2>
    <div v-if="getCollectorsAlert" class="notification is-success">
      <button class="delete" v-on:click="getCollectorsAlert=false"></button>
      Error getting collectors: {{ getCollectorsError }}
    </div>
    <div v-if="collectorDeletedErrorAlert" class="notification is-danger">
      <button class="delete" v-on:click="collectorDeletedErrorAlert"></button>
      Error deleting collector: {{ collectorDeletedError }}
    </div>
    <div v-if="collectorDeletedAlert" class="notification is-success">
      <button class="delete" v-on:click="collectorDeletedAlert=false"></button>
      Collector Deleted
    </div>
    <table class="table is-striped">
      <thead>
        <tr>
          <th>Name</th>
          <th>Certificate</th>
          <th>Key</th>
          <th>Server Certificate</th>
          <th>Delete</th>
        </tr>
      </thead>
      <tbody id="usersTable">
        <tr v-for="collector in collectors" v-bind:key="collector">
          <td>{{ collector }}</td>
          <td><a v-on:click="downloadCertificate(collector)">Download</a></td>
          <td><a v-on:click="downloadKey(collector)">Download</a></td>
          <td><a v-on:click="downloadServerCertificate()">Download</a></td>
          <td>
            <button class="button is-danger" v-on:click="deleteCollector(collector)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <h2 class="subtitle">Add Collector</h2>
    <div v-if="collectorCreatedAlert" class="notification is-success">
      <button class="delete" v-on:click="collectorCreatedAlert=false"></button>
        Collector created
    </div>
    <div v-if="errorCreatingCollectorAlert" class="notification is-danger">
      <button class="delete" v-on:click="errorCreatingCollectorAlert=false"></button>
        Error creating collector: {{ this.collectorCreatedError }}
    </div>
    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Name</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control is-expanded">
            <input class="input" v-model="newCollectorName">
          </p>
        </div>
      </div>
    </div>
    <div class="field is-horizontal">
      <div class="field-label">
        <!-- Left empty for spacing -->
      </div>
      <div class="field-body">
        <div class="field">
          <div class="control">
            <button class="button is-info" v-on:click="addCollector">
              Create
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  const FileDownload = require('js-file-download');

  export default {
    name: 'CollectorManager',
    data: function() {
      return {
        newCollectorName: "",
        collectorCreatedError: "",
        collectorDeletedError: "",
        getCollectorsError: "",
        collectorCreatedAlert: false,
        collectorDeletedErrorAlert: false,
        collectorDeletedAlert: false,
        errorCreatingCollectorAlert: false,
        errorDeletingCollectorAlert: false,
        getCollectorsAlert: false,
        collectors: []
      }
    },
    methods: {
      populateCollectors: function () {
        axios({url: '/collectors', method: 'GET', crossdomain: true, withCredentials: true })
          .then((resp) => {
              this.collectors = resp.data
            })
          .catch((err) => {
            this.getCollectorsAlert = true
            this.getCollectorsError = err.response.data.error
        })
      },
      addCollector: function() {
        var update = {
          "name": this.newCollectorName
        }
        axios({url: '/collectors/create', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then(() => {
            this.collectorCreatedAlert = true
            this.populateCollectors() // just add to the data object?
          })
          .catch((err) => {
            this.collectorCreatedError = err.response.data.error
            this.errorCreatingCollectorAlert = true
        })
      },
      downloadCertificate: function(name) {
        var update = {
          "name": name
        }
        axios({url: '/collector/certificate', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then((resp) => {
            FileDownload(resp.data, name + ".pem")
          })
          .catch(() => {
            // catch this more specifically
        })
      },
      downloadKey: function(name) {
        var update = {
          "name": name
        }
        axios({url: '/collector/key', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then((resp) => {
            FileDownload(resp.data, name + ".key")
          })
          .catch(() => {
            // catch this more specifically
        })
      },
      downloadServerCertificate: function() {
        axios({url: '/collector/server_certificate', method: 'POST', crossdomain: true, withCredentials: true })
          .then((resp) => {
            FileDownload(resp.data, "wave.pem")
          })
          .catch(() => {
            // catch this more specifically
        })
      },
      deleteCollector: function(name) {
        var update = {
          "name": name
        }
        axios({url: '/collectors/delete', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then(() => {
            this.collectorDeletedAlert = true
            this.populateCollectors() // just remove from the data object?
          })
          .catch((err) => {
            this.collectorDeletedError = err.response.data.error
            this.collectorDeletedErrorAlert = true
        })
      },
    },
    beforeMount(){
      this.populateCollectors()
    }
  }
</script>

<style scoped>
  h1 {
    text-align: center;
    padding-top: 30px;
  }
  .field-label {
    max-width: fit-content;
    min-width: 60px;
    text-align: justify;
  }
  input {
    width: auto;
  }
</style>
