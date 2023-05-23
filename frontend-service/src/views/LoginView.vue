<script setup lang="ts">
import LoginForm from "../components/auth/LoginForm.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
import HeroCard from "@/components/layout/HeroCard.vue";
</script>

<template>
  <StandardContainer>
    <HeroCard>
      <h1 class="text-3xl font-bold">Login</h1>
      <LoginForm />
    </HeroCard>
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "Login | Validity";
  },
});
</script>
