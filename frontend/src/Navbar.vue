<template>
  <div id="navbar">
    <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <div class="navbar-item">
          <img src="wave.svg" width="112" height="28">
        </div>
        <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
      <div id="navbar-menu" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item" v-on:click="dashboard">
            Dashboard
          </a>
          <a class="navbar-item" v-on:click="ids">
            IDS
          </a>
          <a class="navbar-item" v-on:click="rules">
            Rules
          </a>
          <a class="navbar-item" v-on:click="visualization">
            Visualization
          </a>
        </div>
        <div class="navbar-end">
          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">
              {{ this.$store.getters.currentUser.username }}
            </a>
            <div class="navbar-dropdown">
              <a class="navbar-item" v-on:click="settings">
                Account Settings
              </a>
              <a class="navbar-item" v-if="this.$store.getters.currentUser.admin" v-on:click="userManager">
                User Management
              </a>
              <a class="navbar-item" v-if="this.$store.getters.currentUser.admin" v-on:click="collectorManager">
                Collector Management
              </a>
            </div>
          </div>
          <div class="navbar-item">
            <div class="buttons">
              <a class="button is-danger is-outlined" v-on:click="logout">
                Log out
              </a>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'app',
  methods: {
    logout: function () {
      this.$store.dispatch('logout')
    },
    settings: function () {
      this.$store.dispatch('settings')
    },
    dashboard: function () {
      this.$store.dispatch('dashboard')
    },
    userManager: function () {
      this.$store.dispatch('userManager')
    },
    collectorManager: function () {
      this.$store.dispatch('collectorManager')
    },
    ids: function () {
      this.$store.dispatch('ids')
    },
    rules: function () {
      this.$store.dispatch('rules')
    },
    visualization: function () {
      this.$store.dispatch('visualization')
    }
  }
}

document.addEventListener('DOMContentLoaded', () => {

  // Get all "navbar-burger" elements
  const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

  // Check if there are any navbar burgers
  if ($navbarBurgers.length > 0) {

    // Add a click event on each of them
    $navbarBurgers.forEach( el => {
      el.addEventListener('click', () => {

        // Get the target from the "data-target" attribute
        const target = el.dataset.target;
        const $target = document.getElementById(target);

        // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
        el.classList.toggle('is-active');
        $target.classList.toggle('is-active');

      });
    });
  }

});
</script>

<style>
</style>
