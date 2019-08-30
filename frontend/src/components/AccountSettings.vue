<template>
  <div class="container">
    <h1 class="title">Account Setting</h1>
    <h2 class="subtitle">Change Password</h2>
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
        <label class="label">Confirm</label>
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
export default {
  name: 'AccountSettings',
  data: function() {
    return {
      oldPassword: "",
      newPassword: "",
      newPasswordConfirm: ""
    }
  },
  methods: {
    updatePassword: function () {
      //debugger;
    }
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
  padding-top: 30px;
}
</style>
