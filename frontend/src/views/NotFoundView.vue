<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
  <div class="not-found">
    <h2>The page you requested can't be found!</h2>
    <p>
      Return to
      <RouterLink to="/">Home</RouterLink>
    </p>
  </div>
</template>
