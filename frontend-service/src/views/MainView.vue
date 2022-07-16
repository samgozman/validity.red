<script setup lang="ts">
import BigButton from "../components/buttons/BigButton.vue";
import LogoText from "../components/LogoText.vue";
</script>

<template>
  <main>
    <div class="hero min-h-screen bg-base-100 m-6 rounded-box">
      <div class="hero-content text-center">
        <div class="max-w-xl">
          <article class="prose lg:prose-xl mb-10">
            <h1><LogoText /></h1>
            <p>
              Lorem ipsum dolor sit, amet consectetur adipisicing elit. Pariatur
              sunt aliquam amet quaerat provident possimus eius, ex quos odit
              earum debitis molestias eligendi accusamus adipisci consequatur
              reiciendis quo aliquid nemo?
            </p>
          </article>
          <BigButton msg="Some button" />
        </div>
      </div>
    </div>
  </main>
</template>
