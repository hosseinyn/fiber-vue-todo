<script setup>
import { ref , onMounted } from 'vue';
import axios from 'axios';
import Cookies from 'js-cookie';
import { useRouter } from 'vue-router';

import { toast } from "vue3-toastify";
import "vue3-toastify/dist/index.css";

const router = useRouter()


const usernameValue = ref("")
const passwordValue = ref("")

const handleUser = async () => {
    const jwt_token = Cookies.get("jwt");
    const res = await axios.get("http://127.0.0.1:4000/todos/check-token", {
      headers: {
        Authorization: `Bearer ${jwt_token}`
      }
    })

    if (res.status === 200) {
      router.push("/dashboard")
    }
}

const handleLogin = async () => {
    axios.post('http://127.0.0.1:4000/auth/login', {
    username: usernameValue.value,
    password: passwordValue.value
    })
    .then(function (response) {
        Cookies.set("jwt" , response.data.token);
        router.push("/dashboard")
    })
    .catch(function (error) {
        toast(error.response.data.error, {
            "theme": "dark",
            "type": "error",
            "dangerouslyHTMLString": true
        })
    });
}

onMounted(() => {
    handleUser()
})

</script>

<template>
    <div class="flex flex-col items-center text-white">
        <section>
            <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <a href="#" class="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
                    Fiber Vue Todo App
                </a>
                <div
                    class="bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700" style="width: 400px !important;">
                    <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                        <h1
                            class="text-xl text-center font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                            Log In
                        </h1>
                        <form class="space-y-4 md:space-y-6" action="#" @submit.prevent="handleLogin">
                            <div>
                                <label for="username"
                                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Username : </label>
                                <input type="text" name="username" id="username"
                                    class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                    placeholder="Enter your username..." required v-model="usernameValue" oninvalid="this.setCustomValidity('Username is required')"
                                    oninput="this.setCustomValidity('')">
                            </div>
                            <div>
                                <label for="password"
                                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                                <input type="password" name="password" id="password" placeholder="Enter your password..."
                                    class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                    required v-model="passwordValue" oninvalid="this.setCustomValidity('Password is required')"
                                    oninput="this.setCustomValidity('')">
                            </div>
                            <button type="submit"
                                class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign
                                in</button>
                            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
                                Don’t have an account yet? <RouterLink to="/signup"
                                    class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign
                                    up</RouterLink>
                            </p>
                        </form>
                    </div>
                </div>
            </div>
        </section>
    </div>
</template>