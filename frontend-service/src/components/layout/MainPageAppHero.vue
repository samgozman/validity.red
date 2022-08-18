<template>
  <div class="min-h-screen w-full grid grid-cols-1 content-between">
    <div class="flex flex-wrap items-center py-10 lg:pt-24 2xl:pt-32 px-4">
      <div class="text-center w-full lg:w-[60%] xl:w-[50%] min-h-full mx-auto">
        <slot />
      </div>
    </div>
    <div class="flex justify-center px-4">
      <div
        class="mockup-window border bg-base-300 w-full lg:w-[60%] xl:w-[50%] shadow-xl"
      >
        <div class="flex justify-center bg-base-100">
          <div class="carousel w-full">
            <div id="slide1" class="carousel-item relative w-full">
              <img src="/img/example.png" class="w-full" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
