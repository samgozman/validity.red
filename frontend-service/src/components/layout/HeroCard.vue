<template>
  <div class="hero md:min-h-screen md:min-h-screen-safe bg-base-200">
    <div class="hero-content flex-col lg:flex-row-reverse min-w-full px-2 py-4">
      <div
        class="card card-compact md:card-normal flex-shrink-0 w-full max-w-md shadow-2xl bg-base-100"
      >
        <div class="card-body">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>
