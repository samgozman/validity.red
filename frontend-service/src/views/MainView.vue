<script setup lang="ts">
import LogoText from "../components/LogoText.vue";
import FullScreenHeroVue from "../components/layout/FullScreenHero.vue";
</script>

<template>
  <main>
    <FullScreenHeroVue>
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
        <button class="btn">Some button</button>
      </div>
    </FullScreenHeroVue>
  </main>
</template>
