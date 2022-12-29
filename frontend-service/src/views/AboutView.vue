<script setup lang="ts">
import StandardContainer from "@/components/layout/StandardContainer.vue";
import { LogoGithub, BeerOutline } from "@vicons/ionicons5";
</script>

<template>
  <StandardContainer>
    <article class="prose lg:prose-xl">
      <h2 class="pt-6">About Validity.Red</h2>
      <p>
        👋 Hi! My name is Sam, creator of the Validity.Red
        <a
          href="https://github.com/samgozman/validity.red"
          rel="noopener"
          target="_blank"
        >
          open-source
        </a>
        web app.
      </p>
      <p>
        I believe that keeping track of the validity periods for important
        documents can be a hassle. That's why I've created a simple, easy-to-use
        tool that helps you manage these expiration dates.
      </p>
      <p>
        I know how important it is to
        <strong>keep track of expiration dates</strong> for things like
        passports, driver's licenses, and insurance policies, and I want to
        make it as easy as possible for users to stay organized. With this app,
        you can easily set up notifications for when a document is set to
        expire. These notifications can be
        <strong>synced with your calendar app</strong>, so you never miss an
        important deadline.
      </p>
      <p>
        Thank you for considering my project, and I hope you will give it a try
        ❤️
      </p>
      <p>
        <strong>P.S.</strong> I believe in the power of open-source software.
        That's why I've made my project available on GitHub for anyone to
        contribute to. If you have ideas for new features or ways to improve our
        app, I encourage you to share them on our
        <a
          href="https://github.com/samgozman/validity.red"
          rel="noopener"
          target="_blank"
        >
          GitHub page</a
        >.
      </p>
    </article>
    <div class="pt-3 text-center sm:text-left">
      <a
        class="btn mr-2 w-full sm:w-auto mb-3"
        href="https://github.com/samgozman/validity.red"
        rel="noopener"
        target="_blank"
      >
        <LogoGithub class="w-8 pr-2" />
        <span>Star on Github</span>
      </a>
      <a
        class="btn btn-warning w-full sm:w-auto mb-3"
        href="https://ko-fi.com/samgozman"
        rel="noopener"
        target="_blank"
      >
        <BeerOutline class="w-8 pr-2" />
        <span>By me a cider!</span>
      </a>
    </div>
  </StandardContainer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  created() {
    document.title = "About | Validity.Red";
  },
});
</script>
