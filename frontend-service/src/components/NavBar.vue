<script setup lang="ts">
import BigButton from "./buttons/BigButton.vue";
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
      <a class="btn btn-ghost normal-case text-xl">Valitidy.RED</a>
      <ul class="hidden md:flex menu menu-horizontal p-0">
        <NavBarItems />
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
          <NavBarUserItems />
        </ul>
      </div>
    </div>
  </div>
</template>
