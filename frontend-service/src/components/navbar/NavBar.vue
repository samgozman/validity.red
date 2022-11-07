<script setup lang="ts">
import { RouterLink } from "vue-router";
import {
  navbarItems,
  navbarItemsUser,
  navbarItemsUserAuth,
} from "@/components/navbar/items";
import LogoText from "@/components/LogoText.vue";
import NavItem from "./NavItem.vue";
import { state } from "@/state";
</script>

<template>
  <div class="navbar bg-base-100">
    <div class="navbar-start">
      <div class="flex-none md:hidden">
        <label for="left-sidebar" class="btn btn-square btn-ghost">
          <ion-icon name="menu-outline" size="large"></ion-icon>
        </label>
      </div>
      <a class="btn btn-ghost normal-case text-xl"><LogoText /></a>
      <ul class="hidden md:flex menu menu-horizontal p-0">
        <NavItem
          v-for="item in navbarItems"
          v-bind:key="item.name"
          :item="item"
        />
      </ul>
    </div>
    <div class="navbar-center"></div>
    <div class="navbar-end">
      <RouterLink
        class="btn btn-circle mr-2 lg:mr-4 lg:rounded-full lg:w-max lg:px-4"
        to="/documents/create"
      >
        <ion-icon name="add-outline" class="text-xl lg:mr-1"></ion-icon>
        <span class="hidden lg:block">Add new</span>
      </RouterLink>
      <div class="dropdown dropdown-end lg:mr-4">
        <label tabindex="0" class="btn btn-neutral btn-circle">
          <ion-icon name="person-circle-outline" class="text-4xl"></ion-icon>
        </label>
        <ul
          tabindex="0"
          class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-48"
        >
          <NavItem
            v-for="item in state.user.isAuthenticated
              ? navbarItemsUserAuth
              : navbarItemsUser"
            v-bind:key="item.name"
            :item="item"
          />
        </ul>
      </div>
    </div>
  </div>
</template>
