<template>
  <div class="container">
    <h1 class="title">Account Setting</h1>
    <h2 class="subtitle">Change Password</h2>
    <div v-if="passwordMismatchAlert" class="notification is-danger">
      <button class="delete" v-on:click="passwordMismatchAlert=false"></button>
        New passwords do not match
    </div>
    <div v-if="passwordIncorrectAlert" class="notification is-danger">
      <button class="delete" v-on:click="passwordIncorrectAlert=false"></button>
        Old password is incorrect
    </div>
    <div v-if="passwordSameAlert" class="notification is-danger">
      <button class="delete" v-on:click="passwordSameAlert=false"></button>
        Old password and new password are the same
    </div>
    <div v-if="passwordUpdatedAlert" class="notification is-success">
      <button class="delete" v-on:click="passwordUpdatedAlert=false"></button>
        Password Updated!
    </div>
    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">Old Password</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control is-expanded">
            <input class="input" type="password" v-model="oldPassword">
          </p>
        </div>
      </div>
    </div>
    <div class="field is-horizontal">
      <div class="field-label is-normal">
        <label class="label">New Password</label>
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
      <div class="field-label is-normal">
        <label class="label">Confirm Password</label>
      </div>
      <div class="field-body">
        <div class="field">
          <p class="control is-expanded">
            <input class="input" type="password" v-model="newPasswordConfirm">
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
            <button class="button is-info" v-on:click="updatePassword">
              Submit
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
  name: 'AccountSettings',
  data: function() {
    return {
      oldPassword: "",
      newPassword: "",
      newPasswordConfirm: "",
      passwordMismatchAlert: false,
      passwordIncorrectAlert: false,
      passwordSameAlert: false,
      passwordUpdatedAlert: false
    }
  },
  methods: {
    updatePassword: function () {
      if (this.newPassword != this.newPasswordConfirm) {
        this.passwordMismatchAlert = true
        return
      }
      if (this.newPassword == this.oldPassword) {
        this.passwordSameAlert = true
        return
      }
      // make sure password isn't being set to the same password
      var update = {
        "old_password": this.oldPassword,
        "new_password": this.newPassword
      }
      axios({url: '/users/password', data: update, method: 'POST', crossdomain: true, withCredentials: true })
        .then(() => {
          this.passwordUpdatedAlert = true
        })
        .catch(() => {
          // catch this more specifically
          this.passwordIncorrectAlert = true
      })
    }
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
    min-width: 140px;
    text-align: justify;
  }
  input {
    width: auto;
  }
</style>
