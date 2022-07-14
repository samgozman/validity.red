<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <div class="navbar bg-base-100">
    <div class="navbar-start">
      <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
    </div>
    <div class="navbar-center">
      <ul class="menu menu-horizontal p-0">
        <li><RouterLink to="/">Home</RouterLink></li>
        <li><RouterLink to="/about">About</RouterLink></li>
      </ul>
    </div>
    <div class="navbar-end">
      <label for="right-sidebar" class="drawer-button btn btn-primary">
        Get started
      </label>
    </div>
  </div>
</template>
