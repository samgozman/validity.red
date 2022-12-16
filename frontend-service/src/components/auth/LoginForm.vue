<script setup lang="ts">
import { RouterLink } from "vue-router";
import InputLabel from "../elements/InputLabel.vue";
</script>

<template>
  <form @submit.prevent="login">
    <div class="form-control">
      <InputLabel label="Email" />
      <input
        type="email"
        v-model="email"
        placeholder="email"
        class="input input-bordered"
      />
    </div>
    <div class="form-control">
      <InputLabel label="Password" />
      <input
        type="password"
        v-model="password"
        placeholder="password"
        class="input input-bordered"
      />
      <label class="label">
        <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
      </label>
    </div>
    <div v-show="errorMsg" class="badge badge-error badge-outline w-full">
      {{ errorMsg }}
    </div>
    <div class="form-control mt-6">
      <button
        class="btn btn-primary"
        :disabled="password.length < 6 || email.length < 4"
        type="submit"
      >
        Login
      </button>
    </div>
  </form>
  <div class="divider">Don't have an account yet?</div>
  <RouterLink class="btn btn-secondary" to="/register">Sign Up</RouterLink>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { AuthService } from "./AuthService";
import { ErrorDecoder } from "@/services/ErrorDecoder";
import { setIsAuthenticated } from "@/state";

export default defineComponent({
  data() {
    return {
      email: "",
      password: "",
      errorMsg: "",
    };
  },
  methods: {
    async login() {
      try {
        await AuthService.userLogin({
          email: this.email,
          password: this.password,
        });
        setIsAuthenticated(true);
        this.$router.push(
          (this.$route.query.redirect as string) || "/dashboard"
        );
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
});
</script>
