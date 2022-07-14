<script setup lang="ts">
import { RouterLink } from "vue-router";
import BigButton from "./buttons/BigButton.vue";
</script>

<template>
  <div class="navbar bg-base-100">
    <div class="navbar-start">
      <a class="btn btn-ghost normal-case text-xl">daisyUI</a>
      <ul class="menu menu-horizontal p-0">
        <li><RouterLink to="/">Home</RouterLink></li>
        <li><RouterLink to="/dashboard">Dashboard</RouterLink></li>
        <li><RouterLink to="/about">About</RouterLink></li>
      </ul>
    </div>
    <div class="navbar-center"></div>
    <div class="navbar-end">
      <BigButton msg="+ Add new" class="mr-3" />
      <div class="dropdown dropdown-end">
        <label tabindex="0" class="btn btn-ghost btn-circle avatar">
          <div class="w-10 rounded-full">
            <img src="https://placeimg.com/80/80/people" />
          </div>
        </label>
        <ul
          tabindex="0"
          class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52"
        >
          <li><a>Settings</a></li>
          <li><a>Logout</a></li>
        </ul>
      </div>
    </div>
  </div>
</template>
