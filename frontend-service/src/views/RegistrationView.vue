<script setup lang="ts">
import RegistrationFormVue from "../components/auth/RegistrationForm.vue";
</script>

<template>
  <main>
    <div class="hero min-h-screen bg-base-200">
      <div class="hero-content flex-col lg:flex-row-reverse min-w-full">
        <div class="card flex-shrink-0 w-full max-w-md shadow-2xl bg-base-100">
          <div class="card-body">
            <h1 class="text-5xl font-bold">Register</h1>
            <RegistrationFormVue />
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
