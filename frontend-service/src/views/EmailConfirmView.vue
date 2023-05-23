<script setup lang="ts">
import EmailConfirmation from "../components/auth/EmailConfirmation.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
</script>

<template>
  <StandardContainer>
    <EmailConfirmation />
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "Email confirmation | Validity";
  },
});
</script>
