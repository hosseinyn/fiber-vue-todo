<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter , RouterLink } from 'vue-router';
import Cookie from 'js-cookie'

import { toast } from "vue3-toastify";
import "vue3-toastify/dist/index.css";

const router = useRouter()


const currentPasswordValue = ref("")
const newPasswordValue = ref("")
const confirmNewPasswordValue = ref("")

const changePasswordError = ref("")

const jwt_token = Cookie.get("jwt");

const handleChangePassword = async () => {
    if (newPasswordValue.value != confirmNewPasswordValue.value){
        changePasswordError.value = "Passwords are diffrent.";
        return
    }
    axios.post('http://127.0.0.1:4000/todos/change-password', {
        current_password: currentPasswordValue.value,
        new_password: newPasswordValue.value
    } , {
        headers : {
            'Authorization' : `Bearer ${jwt_token}`
        }
    })
    .then(function (response) {
        router.push("/dashboard");
    })
    .catch(function (error) {
        toast(error.response.data.error, {
            "theme": "dark",
            "type": "error",
            "dangerouslyHTMLString": true
        })
    });
}


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
                            Change Password
                        </h1>
                        <form class="space-y-4 md:space-y-6" action="#" @submit.prevent="handleChangePassword">
                            <div>
                                <label for="current_password"
                                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Current Password : </label>
                                <input type="password" name="current_password" id="current_password"
                                    class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                    placeholder="Enter your current password..." required v-model="currentPasswordValue" oninvalid="this.setCustomValidity('Current password is required')"
                                    oninput="this.setCustomValidity('')">
                            </div>
                            <div>
                                <label for="new_password"
                                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">New Password : </label>
                                <input type="password" name="new_password" id="new_password" placeholder="Enter your new password..."
                                    class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                    required v-model="newPasswordValue" oninvalid="this.setCustomValidity('New password is required')"
                                    oninput="this.setCustomValidity('')">
                            </div>
                            <div>
                                <label for="confirm_new_password"
                                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Confirm New Password : </label>
                                <input type="password" name="confirm_new_password" id="confirm_new_password" placeholder="Repeat your new password..."
                                    class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                    required v-model="confirmNewPasswordValue" oninvalid="this.setCustomValidity('Confirm password is required')"
                                    oninput="this.setCustomValidity('')">
                            </div>
                            <button type="submit"
                                class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Change Passsword
                                </button>

                                <p class="text-center" style="color: #ff0000;">{{changePasswordError}}</p>
                        </form>
                    </div>
                </div>
            </div>
        </section>
    </div>
</template>