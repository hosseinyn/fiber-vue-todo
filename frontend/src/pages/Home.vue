<script setup>
import "../styles/home.css";

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUser , faGauge , faDatabase} from "@fortawesome/free-solid-svg-icons";
</script>

<template>
    <section>
        <div class="grid max-w-screen-xl px-4 py-8 mx-auto lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12">
            <div class="mr-auto place-self-center lg:col-span-7">
                <h1
                    class="max-w-2xl mb-4 text-4xl font-extrabold tracking-tight leading-none md:text-5xl xl:text-6xl dark:text-white">
                    Todo application</h1>
                <p class="max-w-2xl mb-6 font-light text-gray-500 lg:mb-8 md:text-lg lg:text-xl dark:text-gray-400">
                    A Full-Stack Todo List App for Learning and Practice
                    This project is a full-stack practice application centered around a Todo List system. The frontend
                    is built with Vue.js, TailwindCSS, and SCSS, while the backend uses Golang with the Fiber framework.
                    It includes JWT-based authentication and persistent todo storage using a SQLite database through a
                    secure user account system.
                    This app is created as an educational project to explore and practice full-stack development with
                    modern technologies.</p>
                <a href="/login"
                    class="inline-flex items-center justify-center px-5 py-3 mr-3 text-base font-medium text-center text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:focus:ring-primary-900">
                    Login
                    <svg class="w-5 h-5 ml-2 -mr-1" fill="currentColor" viewBox="0 0 20 20"
                        xmlns="http://www.w3.org/2000/svg">
                        <path fill-rule="evenodd"
                            d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
                            clip-rule="evenodd"></path>
                    </svg>
                </a>
                <a href="/register">
                    <button
                        class="inline-flex items-center justify-center px-5 py-3 text-base font-medium text-center rounded-lg">
                        Sign Up
                    </button>
                </a>
            </div>
            <div class="hidden lg:mt-0 lg:col-span-5 lg:flex">
                <img src="../assets/images/hero-section.png" alt="mockup">
            </div>
        </div>
    </section>

    <section id="features">
        <div class="py-8 px-4 max-w-screen-xl sm:py-16 lg:px-6">
            <div class="max-w-screen-md mb-8 lg:mb-16">
                <h2 class="mb-4 text-4xl tracking-tight font-extrabold text-gray-900 dark:text-white">Todo Application Features</h2>
                <p class="text-gray-500 sm:text-xl dark:text-gray-400">Several features of this project set it apart from typical educational or practice-oriented todo list projects, enhancing its realism and aligning it more closely with real-world applications : </p>
            </div>
            <div class="space-y-8 md:grid md:grid-cols-2 lg:grid-cols-3 md:gap-12 md:space-y-0">
                <div>
                    <div
                        class="flex justify-center items-center mb-4 w-10 h-10 rounded-full bg-primary-100 lg:h-12 lg:w-12 dark:bg-primary-900">
                        <FontAwesomeIcon :icon="faUser" style="color: #fff; font-size: 30px;" />
                    </div>
                    <h3 class="mb-2 text-xl font-bold dark:text-white">User Account</h3>
                        <p class="text-gray-500 dark:text-gray-400">This project uses a JWT-based authentication system on both the frontend and backend. The todo-related APIs are protected and require authentication, ensuring that todos are categorized and personalized per user. Additionally, the website includes features such as account deletion and password change.</p>
                </div>
                <div>
                    <div
                        class="flex justify-center items-center mb-4 w-10 h-10 rounded-full bg-primary-100 lg:h-12 lg:w-12 dark:bg-primary-900">
                        <FontAwesomeIcon :icon="faGauge" style="color: #fff; font-size: 30px;" />
                    </div>
                    <h3 class="mb-2 text-xl font-bold dark:text-white">Speed</h3>
                    <p class="text-gray-500 dark:text-gray-400">While speed may not be critical in simple projects, one of the key features of this project is its high-performance communication, powered by a Golang backend using the Fiber framework, which delivers fast and efficient APIs.</p>
                </div>
                <div>
                    <div
                        class="flex justify-center items-center mb-4 w-10 h-10 rounded-full bg-primary-100 lg:h-12 lg:w-12 dark:bg-primary-900">
                        <FontAwesomeIcon :icon="faDatabase" style="color: #fff; font-size: 30px;" />
                    </div>
                    <h3 class="mb-2 text-xl font-bold dark:text-white">Database / Storage</h3>
                    <p class="text-gray-500 dark:text-gray-400">This project goes several steps beyond typical frontend-only todo list applications that store data in cookies or local storage. All todos are properly organized and stored in a database, making them accessible from anywhere through a user account. The project uses a lightweight and simple file-based database — SQLite.</p>
                </div>
            </div>
        </div>
    </section>
</template>