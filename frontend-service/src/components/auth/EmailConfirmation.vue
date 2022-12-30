<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <div class="hero min-h-screen min-h-screen-safe bg-base-200">
    <div class="hero-content flex-col lg:flex-row-reverse min-w-full">
      <div class="card flex-shrink-0 w-full max-w-md shadow-2xl bg-base-100">
        <div class="card-body">
          <div v-if="verified">
            <p>Email address is confirmed!</p>
            <p>
              You can now
              <RouterLink to="/login">login</RouterLink>
            </p>
          </div>
          <div v-else>
            <p>Verifying email address...</p>
          </div>
          <div v-if="isError">
            <p class="text-red-500">
              Error: this link was already used or expired.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { AuthService } from "./AuthService";

export default defineComponent({
  data() {
    return {
      verified: false,
      isError: false,
    };
  },
  methods: {
    async verifyEmail() {
      try {
        await AuthService.userVerifyEmail(this.$route.query.token as string);
        this.verified = true;
        setTimeout(() => {
          this.$router.push("/login");
        }, 7000);
      } catch (error) {
        this.verified = false;
        this.isError = true;
      }
    },
  },
  mounted() {
    this.verifyEmail();
  },
});
</script>
