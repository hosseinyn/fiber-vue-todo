<script setup>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
</script>

<template>

  <Header />
  <router-link to="/"></router-link>
  
  <router-view />

  <Footer />
  
</template>