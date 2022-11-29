<script setup lang="ts">
import { RouterLink } from "vue-router";
import LoginForm from "../components/auth/LoginForm.vue";
import StandardContainer from "@/components/layout/StandardContainer.vue";
</script>

<template>
  <StandardContainer>
    <div class="hero min-h-screen bg-base-200">
      <div class="hero-content flex-col lg:flex-row-reverse p-0">
        <div class="text-center lg:text-left">
          <h1 class="text-4xl font-bold">Login now!</h1>
          <p class="py-3 min-w-[30vw] lg:py-6 lg:max-w-[40vw]">
            Login to your account to access all the features of the app. If you
            don't have an account yet, you can
            <RouterLink to="/register" class="underline">sign up</RouterLink>
            for free.
          </p>
        </div>
        <div
          class="card-compact card xl:card-normal flex-shrink-0 w-full max-w-md shadow-2xl bg-base-100"
        >
          <div class="card-body">
            <LoginForm />
          </div>
        </div>
      </div>
    </div>
  </StandardContainer>
</template>
