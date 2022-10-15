<script setup lang="ts">
import { RouterLink } from "vue-router";
import { navbarItems, navbarItemsUser } from "@/components/navbar/items";
import LogoText from "@/components/LogoText.vue";
import NavBarItems from "./NavBarItems.vue";
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
        <NavBarItems :items="navbarItems" />
      </ul>
    </div>
    <div class="navbar-center"></div>
    <div class="navbar-end">
      <RouterLink
        class="btn btn-circle mr-2 sm:mr-4 sm:rounded-full sm:w-max sm:px-4"
        to="/documents/create"
      >
        <ion-icon name="add-outline" class="text-xl sm:mr-1"></ion-icon>
        <span class="hidden sm:block">Add new</span>
      </RouterLink>
      <div class="dropdown dropdown-end sm:mr-4">
        <label tabindex="0" class="btn btn-neutral btn-circle">
          <ion-icon name="person-circle-outline" class="text-4xl"></ion-icon>
        </label>
        <ul
          tabindex="0"
          class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-48"
        >
          <NavBarItems :items="navbarItemsUser" />
        </ul>
      </div>
    </div>
  </div>
</template>
