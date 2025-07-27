<script setup>
import { ref } from 'vue'
import { faHome , faThumbsUp , faFile } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import { RouterLink, RouterView } from 'vue-router'

const isMobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

</script>

<template>
  <header>
    <nav class="border-gray-200 px-4 lg:px-6 py-2.5">
      <div class="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl">
        <a href="#" class="flex items-center">
          <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">Fiber Vue Todo</span>
        </a>
        <div class="flex items-center lg:order-2">
          <RouterLink to="/login" class="text-gray-800 dark:text-white font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2  focus:outline-none dark:focus:ring-gray-800">
            Log in
          </RouterLink>
          <RouterLink to="/signup" class="text-white focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-primary-600 dark:hover:bg-primary-700 focus:outline-none dark:focus:ring-primary-800">
            Register
          </RouterLink>
          <button @click="toggleMobileMenu" type="button" class="inline-flex items-center p-2 ml-1 text-sm text-gray-500 rounded-lg lg:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
            <span class="sr-only">Open main menu</span>
            <svg v-if="!isMobileMenuOpen" class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd"
                    d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                    clip-rule="evenodd"></path>
            </svg>
            <svg v-else class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd"
                    d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                    clip-rule="evenodd"></path>
            </svg>
          </button>
        </div>

        <div :class="{'hidden': !isMobileMenuOpen, 'block': isMobileMenuOpen}" class="w-full lg:flex lg:w-auto lg:order-1">
          <ul class="flex flex-col mt-4 font-medium lg:flex-row lg:space-x-8 lg:mt-0">
            <li>
              <RouterLink to="/" class="block py-2 pr-4 pl-3 text-white rounded bg-primary-700 lg:bg-transparent lg:text-primary-700 lg:p-0 dark:text-white" aria-current="page">
                Home <FontAwesomeIcon :icon="faHome" />
              </RouterLink>
            </li>
            <li>
              <a href="/#features" class="block py-2 pr-4 pl-3 lg:p-0 text-gray-400">
                Features <FontAwesomeIcon :icon="faThumbsUp" />
              </a>
            </li>
            <li>
              <a href="https://github.com/hosseinyn/fiber-vue-todo" class="block py-2 pr-4 pl-3 lg:hover:bg-transparent lg:hover:text-primary-700 lg:p-0 text-gray-400">
                About the project <FontAwesomeIcon :icon="faFile" />
              </a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>
</template>
