<script setup lang="ts">
import InputLabel from "../elements/InputLabel.vue";
</script>

<template>
  <form @submit="register" v-if="showForm">
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
    </div>
    <div v-show="error" class="badge badge-error badge-outline w-full">
      {{ errorMsg }}
    </div>
    <div class="form-control mt-6">
      <button
        class="btn btn-primary"
        :disabled="password.length < 6 || email.length < 4"
        type="submit"
      >
        Sign up
      </button>
    </div>
  </form>
  <span v-else>
    <p>You have successfully registered!</p>
    <p>Please click on the confirmation link sent to your email address.</p>
  </span>
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
      showForm: true,
    };
  },
  methods: {
    async register(e: Event) {
      e.preventDefault();

      const payload = JSON.stringify({
        action: "UserRegister",
        register: {
          email: this.email,
          password: this.password,
        },
      });

      try {
        const res = await axios.post<AuthResponse>(
          `http://localhost:8080/handle`,
          payload,
          {
            // TODO: Handle duplicate error
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

        this.showForm = false;

        setTimeout(() => {
          this.$router.push("/");
        }, 7000);
      } catch (error) {
        this.error = true;
        this.errorMsg = "An error occurred, please try again";
        // TODO: Push error to Sentry
      }
    },
  },
});
</script>