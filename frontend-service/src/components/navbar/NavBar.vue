<script setup lang="ts">
import { RouterLink } from "vue-router";
import {
  MenuOutline,
  AddOutline,
  PersonCircleOutline,
  LogInOutline,
} from "@vicons/ionicons5";
import { navbarItems, navbarItemsUserAuth } from "@/components/navbar/items";
import LogoText from "@/components/LogoText.vue";
import NavItem from "./NavItem.vue";
import { state } from "@/state";
</script>

<template>
  <div class="navbar bg-base-100">
    <div class="navbar-start md:w-[90%]">
      <div class="flex-none md:hidden">
        <label for="left-sidebar" class="btn btn-square btn-ghost">
          <MenuOutline class="w-8" />
        </label>
      </div>
      <RouterLink
        class="btn btn-ghost normal-case text-xl"
        :to="state.user.isAuthenticated ? '/dashboard' : '/'"
      >
        <LogoText />
      </RouterLink>
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
        class="btn mr-2 lg:mr-4 rounded-full w-max px-4"
        :class="{ hidden: state.user.isAuthenticated }"
        to="/login"
      >
        <LogInOutline class="w-6 mr-1" />
        <span>Sign in</span>
      </RouterLink>
      <RouterLink
        class="btn btn-circle mr-2 lg:mr-4 lg:rounded-full lg:w-max lg:px-4"
        :class="{ hidden: !state.user.isAuthenticated }"
        to="/documents/create"
      >
        <AddOutline class="w-6 lg:mr-1" />
        <span class="hidden lg:block">Add new</span>
      </RouterLink>
      <div :class="{ hidden: !state.user.isAuthenticated }">
        <div class="hidden md:block dropdown dropdown-end lg:mr-4">
          <label tabindex="0" class="btn btn-neutral btn-circle">
            <PersonCircleOutline class="w-10" />
          </label>
          <ul
            tabindex="0"
            class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-48"
          >
            <NavItem
              v-for="item in navbarItemsUserAuth"
              v-bind:key="item.name"
              :item="item"
            />
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
