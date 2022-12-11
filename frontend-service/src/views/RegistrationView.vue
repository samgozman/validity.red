<script setup lang="ts">
import RegistrationFormVue from "../components/auth/RegistrationForm.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
import HeroCard from "@/components/layout/HeroCard.vue";
</script>

<template>
  <StandardContainer>
    <HeroCard>
      <h1 class="text-3xl font-bold">Register</h1>
      <RegistrationFormVue />
    </HeroCard>
  </StandardContainer>
</template>
