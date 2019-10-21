<template>
  <div class="container">
    <h1 class="title">User Management</h1>
    <h2 class="subtitle">Current Users</h2>
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
            <button class="button is-danger">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <h2 class="subtitle">Add User</h2>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'UserManager',
    data: function() {
      return {
        users: [],
      }
    },
    methods: {
      populateUsers: function () {
        axios({url: '/users', method: 'GET', crossdomain: true, withCredentials: true })
          .then((resp) => {
              this.users = resp.data
            })
          .catch(() => {
            // error getting users
        })
      },
      setPassword: function (username) {
        var update = {
          "username": username,
          "password": document.querySelector("input#"+username).value
        }
        console.log(update)
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
</style>
