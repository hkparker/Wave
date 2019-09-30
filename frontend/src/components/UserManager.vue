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
        </tr>
      </thead>
      <tbody id="usersTable">
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
      }
    },
    methods: {
      populateUsers: function () {
        axios({url: '/users', method: 'GET', crossdomain: true, withCredentials: true })
          .then((resp) => {
            var userData = resp.data
            var i;
            for (i = 0; i < userData.length; i++) {
              var currentUser = userData[i]
              // create row
              var tableRow = document.createElement("tr")
              // username column
              var username = document.createElement("td")
              username.innerHTML = currentUser.username
              // admin column
              var admin = document.createElement("td")
              admin.innerHTML = currentUser.admin
              // sessions column
              var sessions = document.createElement("td")
              sessions.innerHTML = currentUser.sessions
              // password set column
              var password = document.createElement("td")
              var passwordField = document.createElement("div")
              passwordField.className = "field has-addons"
              var passwordInputControl = document.createElement("div")
              passwordInputControl.className = "control"
              var passwordInput = document.createElement("input")
              passwordInput.className = "input"
              passwordInput.setAttribute("type", "password")
              passwordInputControl.appendChild(passwordInput)
              var passwordButtonControl = document.createElement("div")
              passwordButtonControl.className = "control"
              var setPasswordButton = document.createElement("a")
              setPasswordButton.className = "button is-danger"
              setPasswordButton.innerHTML = "Set Password"
              passwordButtonControl.appendChild(setPasswordButton)
              password.appendChild(passwordField)
              passwordField.appendChild(passwordInputControl)
              passwordField.appendChild(passwordButtonControl)
              // add columns to row
              tableRow.appendChild(username)
              tableRow.appendChild(admin)
              tableRow.appendChild(sessions)
              tableRow.appendChild(password)
              document.getElementById("usersTable").appendChild(tableRow)
            }
          })
          .catch(() => {
            // error getting users
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
</style>
