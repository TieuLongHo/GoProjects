<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <header>
    <h1>Welcome to the Todo APP</h1>
  </header>

  <RouterView />
</template>

<style scoped>
h1 {
  font-size: 2.6rem;
  font-weight: 500;
  text-align: center;
}

</style>
