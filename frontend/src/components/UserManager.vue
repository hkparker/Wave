<template>
  <div class="container">
    <h1 class="title">User Management</h1>
    <h2 class="subtitle">Current Users</h2>
    <div v-if="getUsersAlert" class="notification is-success">
      <button class="delete" v-on:click="getUsersAlert=false"></button>
        Error getting users: {{ getUsersError }}
    </div>
    <div v-if="passwordSetAlert" class="notification is-success">
      <button class="delete" v-on:click="passwordSetAlert=false"></button>
        New password has been assigned
    </div>
    <div v-if="userDeletedAlert" class="notification is-success">
      <button class="delete" v-on:click="userDeletedAlert=false"></button>
        User has been deleted
    </div>
    <div v-if="errorAssigningPasswordAlert" class="notification is-danger">
      <button class="delete" v-on:click="errorAssigningPasswordAlert=false"></button>
        Error assigning password: {{ passwordSetError }}
    </div>
    <div v-if="errorDeletingUserAlert" class="notification is-danger">
      <button class="delete" v-on:click="errorDeletingUserAlert=false"></button>
        Error deleting user: {{ userDeleteError }}
    </div>
    <table class="table is-striped">
      <thead>
        <tr>
          <th>Username</th>
          <th>Admin</th>
          <th>Active Sessions</th>
          <th>Password Reset</th>
          <th>Delete User</th>
        </tr>
      </thead>
      <tbody id="usersTable">
        <tr v-for="user in users" v-bind:key="user.username">
          <td>{{ user.username }}</td>
          <td>{{ user.admin }}</td>
          <td>{{ user.sessions}}</td>
          <td>
            <div class="field has-addons">
              <div class="control">
                <input :id="user.username" class="input" type="password">
              </div>
              <div class="control">
                <button class="button is-danger" v-on:click="setPassword(user.username)">Set Password</button>
              </div>
            </div>
          </td>
          <td>
            <button class="button is-danger" v-on:click="deleteUser(user.username)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <h2 class="subtitle">Add User</h2>
    <div v-if="userCreatedAlert" class="notification is-success">
      <button class="delete" v-on:click="userCreatedAlert=false"></button>
        User created
    </div>
    <div v-if="errorCreatingUserAlert" class="notification is-danger">
      <button class="delete" v-on:click="errorCreatingUserAlert=false"></button>
        Error creating user: {{ this.userCreatedError }}
    </div>
    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Username</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control is-expanded">
            <input class="input" v-model="newUsername">
          </p>
        </div>
      </div>
    </div>
    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Password</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control is-expanded">
            <input class="input" type="password" v-model="newPassword">
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
            <label class="newUserAdmin">
              <input type="checkbox" v-model="newAdmin">
                Admin
            </label>
          </div>
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
            <button class="button is-info" v-on:click="addUser">
              Create User
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'UserManager',
    data: function() {
      return {
        passwordSetAlert: false,
        userDeletedAlert: false,
        errorDeletingUserAlert: false,
        errorAssigningPasswordAlert: false,
        getUsersAlert: false,
        userCreatedAlert: false,
        errorCreatingUserAlert: false,
        userDeleteError: "",
        userCreatedError: "",
        passwordSetError: "",
        getUsersError: "",
        newUsername: "",
        newPassword: "",
        newAdmin: false,
        users: [],
      }
    },
    methods: {
      populateUsers: function () {
        axios({url: '/users', method: 'GET', crossdomain: true, withCredentials: true })
          .then((resp) => {
              this.users = resp.data
            })
          .catch((err) => {
            this.getUsersAlert = true
            this.getUsersError = err.response.data.error
        })
      },
      setPassword: function (username) {
        var update = {
          "username": username,
          "password": document.querySelector("input#"+username).value
        }
        axios({url: '/users/assign-password', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then(() => {
            this.passwordSetAlert = true
            // if the user sets their own password this way, all their sessions
            // will be destroyed and they will need to be logged out
            this.$store.dispatch("setEnvironment")
          })
          .catch((err) => {
            this.passwordSetAlert = true
            this.passwordSetError = err.response.data.error
        })
      },
      deleteUser: function (username) {
        var update = {
          "username": username
        }
        axios({url: '/users/delete', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then(() => {
            this.userDeletedAlert = true
            this.populateUsers() // just delete from data object?
          })
          .catch((err) => {
            this.errorDeletingUserAlert = true
            this.userDeletedError = err.response.data.error
        })
      },
      addUser: function () {
        var update = {
          "username": this.newUsername,
          "password": this.newPassword,
          "admin": this.newAdmin.toString()
        }
        axios({url: '/users/create', data: update, method: 'POST', crossdomain: true, withCredentials: true })
          .then(() => {
            this.userCreatedAlert = true
            this.populateUsers() // just add to the data object?
          })
          .catch((err) => {
            this.errorCreatingUserAlert = true
            this.userCreatedError = err.response.data.error
        })
      }
    },
    beforeMount(){
      this.populateUsers()
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
    min-width: 80px;
    text-align: justify;
  }
  input {
    width: auto;
  }
</style>
