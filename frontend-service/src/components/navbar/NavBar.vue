<script setup lang="ts">
import { RouterLink } from "vue-router";
import LogoText from "../../components/LogoText.vue";
import NavBarItems from "./NavBarItems.vue";
import NavBarUserItems from "./NavBarUserItems.vue";
</script>

<template>
  <div class="navbar bg-base-100">
    <div class="navbar-start">
      <div class="flex-none md:hidden">
        <label for="left-sidebar" class="btn btn-square btn-ghost">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="inline-block w-6 h-6 stroke-current"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            ></path>
          </svg>
        </label>
      </div>
      <a class="btn btn-ghost normal-case text-xl"><LogoText /></a>
      <ul class="hidden md:flex menu menu-horizontal p-0">
        <NavBarItems />
      </ul>
    </div>
    <div class="navbar-center"></div>
    <div class="navbar-end">
      <RouterLink class="btn mr-4 rounded-full" to="/documents/create">
        <ion-icon name="add-outline" class="text-xl mr-1"></ion-icon> Add new
      </RouterLink>
      <div class="dropdown dropdown-end mr-4">
        <label tabindex="0" class="btn btn-neutral btn-circle">
          <ion-icon name="person-circle-outline" class="text-4xl"></ion-icon>
        </label>
        <ul
          tabindex="0"
          class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-48"
        >
          <NavBarUserItems />
        </ul>
      </div>
    </div>
  </div>
</template>
