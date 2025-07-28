<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="flex flex-col gap-6 items-center mt-7 text-white">
        <img class="h-100 rounded-2xl" src="../assets/images/404-image.png" alt="404 not found">
        <h1 class="text-4xl">Page Not Found</h1>
        <button style="width: 160px;
    height: 34px;"><RouterLink to="/">Go Back To Home</RouterLink></button>
    </div>
</template>