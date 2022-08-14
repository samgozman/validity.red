<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <li>
    <RouterLink to="/login">
      <ion-icon name="log-in-outline"></ion-icon> Login
    </RouterLink>
  </li>
  <li>
    <RouterLink to="/register">
      <ion-icon name="person-add-outline"></ion-icon> Register
    </RouterLink>
  </li>
</template>
