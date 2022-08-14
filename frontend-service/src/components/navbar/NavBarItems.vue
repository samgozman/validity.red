<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <li>
    <RouterLink to="/">
      <ion-icon name="planet-outline"></ion-icon> Home
    </RouterLink>
  </li>
  <li>
    <RouterLink to="/dashboard"
      ><ion-icon name="speedometer-outline"></ion-icon> Dashboard</RouterLink
    >
  </li>
  <li>
    <RouterLink to="/documents"
      ><ion-icon name="documents-outline"></ion-icon>Documents</RouterLink
    >
  </li>
  <li>
    <RouterLink to="/about"
      ><ion-icon name="cube-outline"></ion-icon>About</RouterLink
    >
  </li>
</template>
