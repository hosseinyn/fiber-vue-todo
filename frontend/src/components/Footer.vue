<script setup>
import { faHome , faRightToBracket , faUser , faFile } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
</script>

<template>


    <footer class="rounded-lg shadow-sm m-4">
        <div class="max-w-screen-xl mx-auto p-4 md:py-8">
            <div class="sm:flex sm:items-center sm:justify-between">
                <a href="/" class="flex items-center mb-4 sm:mb-0 space-x-3 rtl:space-x-reverse">
                    <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Fiber Vue Todo</span>
                </a>
                <ul
                    class="flex flex-wrap items-center mb-6 text-sm font-medium text-gray-500 sm:mb-0 dark:text-gray-400">
                    <li>
                        <a href="/" class="hover:underline me-4 md:me-6">Home <FontAwesomeIcon :icon="faHome" /></a>
                    </li>
                    <li>
                        <a href="/login" class="hover:underline me-4 md:me-6">Login <FontAwesomeIcon :icon="faRightToBracket" /></a>
                    </li>
                    <li>
                        <a href="/signup" class="hover:underline me-4 md:me-6">Register <FontAwesomeIcon :icon="faUser" /></a>
                    </li>
                    <li>
                        <a href="https://github.com/hosseinyn/fiber-vue-todo" class="hover:underline">About the project <FontAwesomeIcon :icon="faFile" /></a>
                    </li>
                </ul>
            </div>
            <hr class="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
            <span class="block text-sm text-gray-500 sm:text-center dark:text-gray-400">© 2025 <a
                    href="https://github.com/hosseinyn" class="hover:underline">Hosseinyn</a>. All Rights Reserved.</span>
        </div>
    </footer>
</template>