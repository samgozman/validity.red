<template>
  <form @submit="login">
    <div class="form-control">
      <label class="label">
        <span class="label-text">Email</span>
      </label>
      <input
        type="email"
        v-model="email"
        placeholder="email"
        class="input input-bordered"
      />
    </div>
    <div class="form-control">
      <label class="label">
        <span class="label-text">Password</span>
      </label>
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
    <div v-show="error" class="badge badge-error badge-outline w-full">
      {{ errorMsg }}
    </div>
    <div class="form-control mt-6">
      <button
        class="btn btn-primary"
        :disabled="password.length < 3"
        type="submit"
      >
        Login
      </button>
    </div>
  </form>
  <div class="divider">Don't have an account yet?</div>
  <button class="btn btn-secondary">Sign Up</button>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import axios from "axios";

interface AuthResponse {
  error: boolean;
  message: string;
}

export default defineComponent({
  data() {
    return {
      email: "",
      password: "",
      error: false,
      errorMsg: "",
    };
  },
  methods: {
    async login(e: Event) {
      e.preventDefault();

      const payload = JSON.stringify({
        action: "UserLogin",
        auth: {
          email: this.email,
          password: this.password,
        },
      });

      try {
        const res = await axios.post<AuthResponse>(
          `http://localhost:8080/handle`,
          payload,
          {
            // To pass Set-Cookie header
            withCredentials: true,
            // Handle 401 error like a normal situation
            validateStatus: (status) =>
              (status >= 200 && status < 300) || status === 401,
          }
        );

        const { error, message } = res.data;

        if (error) {
          this.error = true;
          this.errorMsg = message;
          return;
        }

        // window.localStorage.setItem("userData", JSON.stringify(user));

        this.$router.push("/");
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
      }
    },
  },
});
</script>
