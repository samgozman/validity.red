<script setup lang="ts">
import timezones from "timezones-list";
import InputLabel from "../elements/InputLabel.vue";
</script>

<template>
  <form @submit.prevent="register" v-if="showForm">
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
    <div class="form-control">
      <InputLabel label="Timezone" />
      <select class="select select-bordered w-full" v-model="selectedTz">
        <option v-for="(option, index) in timezones" v-bind:key="index">
          {{ option.label }}
        </option>
      </select>
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
import { AuthService } from "./AuthService";
import { ErrorDecoder } from "@/services/ErrorDecoder";

export default defineComponent({
  data() {
    return {
      email: "",
      password: "",
      errorMsg: "",
      selectedTz: "",
      showForm: true,
    };
  },
  methods: {
    async register() {
      const tz = timezones.find((t) => t.label === this.selectedTz);
      const tzCode = tz ? tz.tzCode : "Europe/London";
      try {
        await AuthService.userRegister({
          email: this.email,
          password: this.password,
          timezone: tzCode,
        });

        this.showForm = false;
        setTimeout(() => {
          this.$router.push("/");
        }, 7000);
      } catch (error) {
        this.errorMsg = await ErrorDecoder.decode(error);
      }
    },
  },
  beforeMount() {
    const userTz = Intl.DateTimeFormat().resolvedOptions().timeZone;
    const tz = timezones.find((t) => t.tzCode === userTz);
    this.selectedTz = tz ? tz.label : "Europe/London (GMT+00:00)";
  },
});
</script>
