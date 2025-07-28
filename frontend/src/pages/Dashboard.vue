<script setup>
import '../styles/dashboard.css';
import { ref , onMounted , reactive } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

import Cookie from 'js-cookie'

import { toast } from "vue3-toastify";
import "vue3-toastify/dist/index.css";
import Swal from 'sweetalert2';

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faTrash , faPencil , faCheck } from "@fortawesome/free-solid-svg-icons";

const router = useRouter()

const todoTitle = ref("")
const Todos = ref([])
const todosDetails = reactive([{
    id : 1,
    title : '',
    done : false,
}])
const isTodoChecked = ref(false);

const jwt_token = Cookie.get("jwt");

const listTodo = () => {
    axios.get('http://127.0.0.1:4000/todos/get' , {
        headers : {
            'Authorization' : `Bearer ${jwt_token}`
        }
    })
    .then(function (response) {
        Todos.value = response.data.todos;
        let id_var = 1
        Todos.map((todo) => {
            todosDetails[id_var].id = id_var;
            todosDetails[id_var].title = todo.title;
            todosDetails[id_var].done = todo.title;
            id_var++
        });
    })
    .catch(function (error) {
        toast(error.response.data.error, {
            "theme": "dark",
            "type": "error",
            "dangerouslyHTMLString": true
        })
    })
}

const handleAddTodo = () => {
    axios.post('http://127.0.0.1:4000/todos/add', {
    title : todoTitle.value,
    done : false,
    } , {
        headers : {
            'Authorization' : `Bearer ${jwt_token}`
        }
    })
    .then(function (response) {
        toast("Todo created successfuly" , {
            "theme" : "dark",
            "type" : "success",
            "dangerouslyHTMLString" : true
        })
        listTodo()
    })
    .catch(function (error) {
        toast(error.response.data.error, {
            "theme": "dark",
            "type": "error",
            "dangerouslyHTMLString": true
        })
    });
}

const handleDeleteTodo = (id) => {
    axios.delete(`http://127.0.0.1:4000/todos/delete/${id}` , {
        headers : {
            'Authorization' : `Bearer ${jwt_token}`
        }
    })
    .then(function (response) {
        toast("Todo deleted successfuly" , {
            "theme" : "dark",
            "type" : "success",
            "dangerouslyHTMLString" : true
        })
        listTodo()
    })
    .catch(function (error) {
        toast(error.response.data.error, {
            "theme": "dark",
            "type": "error",
            "dangerouslyHTMLString": true
        })
    })
}

const handleUpdateTodo = (id , done) => {
    Swal.fire({
        title: "Update Todo",
        input: "text",
        inputAttributes: {
            autocapitalize: "off"
        },
        inputPlaceholder: "New todo title...",
        color : "#fff",
        background : "#363636",
        showCancelButton: true,
        confirmButtonText: "Update",
        showLoaderOnConfirm: true,
        preConfirm: async (new_title) => {
            axios.post(`http://127.0.0.1:4000/todos/update/` , {
                id : id,
                title : new_title,
                done : done,
            }, {
                headers : {
                    'Authorization' : `Bearer ${jwt_token}`
                }
            })
            .then(function (response) {
                toast("Todo updated successfuly" , {
                    "theme" : "dark",
                    "type" : "success",
                    "dangerouslyHTMLString" : true
                })
                listTodo()
            })
            .catch(function (error) {
                toast(error.response.data.error, {
                    "theme": "dark",
                    "type": "error",
                    "dangerouslyHTMLString": true
                })
            })
        },
        allowOutsideClick: () => !Swal.isLoading()
        });
}

const handleCheckTodo = (id , content) => {
    axios.post(`http://127.0.0.1:4000/todos/update/` , {
                id : id,
                title : content,
                done : !isTodoChecked.value,
            }, {
                headers : {
                    'Authorization' : `Bearer ${jwt_token}`
                }
            })
            .then(function (response) {
                if (isTodoChecked.value == false){
                    toast("Todo checked successfuly" , {
                        "theme" : "dark",
                        "type" : "success",
                        "dangerouslyHTMLString" : true
                    })
                } else if (isTodoChecked.value == true){
                    toast("Todo unchecked successfuly" , {
                        "theme" : "dark",
                        "type" : "success",
                        "dangerouslyHTMLString" : true
                    })
                }
                listTodo()
                isTodoChecked.value = !isTodoChecked.value
            })
            .catch(function (error) {
                toast(error.response.data.error, {
                    "theme": "dark",
                    "type": "error",
                    "dangerouslyHTMLString": true
                })
            })
}

const handleDeleteAccount = () => {
    Swal.fire({
            title: "Are you sure?",
            text: "You won't be able to revert your account.",
            icon: "warning",
            showCancelButton: true,
            color : "#fff",
            background : "#363636",
            confirmButtonText: "Yes, delete my account."
            }).then((result) => {
            if (result.isConfirmed) {
                axios.delete("http://127.0.0.1:4000/todos/delete-account" , {
                    headers : {
                        'Authorization' : `Bearer ${jwt_token}`
                    }
                })
                .then(function (response) {
                    router.push("/")
                })
                .catch(function (error) {
                    toast(error.response.data.error, {
                        "theme": "dark",
                        "type": "error",
                        "dangerouslyHTMLString": true
                    })
                })
            }
        });
    }

const handleChangePassword = () => {
    router.push("/change-password")
}

onMounted(() => {
  listTodo()
});

</script>

<template>
    <div class="flex flex-col items-center mt-10 gap-5">
        <h1 class="text-white text-3xl">Todos Dashboard</h1>

        <div class="flex gap-3">
            <button class="account-operation-btn" style="background-color: #ff0000;" @click="handleDeleteAccount">Delete Account</button>
            <button class="account-operation-btn" @click="handleChangePassword">Change Password</button>
        </div>

        <form class="flex flex-col gap-2 mt-3 add-todo-form" action="#" @submit.prevent="handleAddTodo">
            <label for="todo-title">Add Todo : </label>
            <div class="flex gap-2">
                <input type="text" placeholder="Add your todo title..." id="todo-title" name="todo-title" v-model="todoTitle">
                <button type="submit" class="add-todo-btn">Add Todo</button>
            </div>
        </form>

        <div class="flex flex-col gap-5">
            <ul >
                <li v-for="[key, value] in Todos.entries(Todos)" :key="key" :class="{ 'checked-todo': value.done }">
                    <div class="flex text-white todo-box text-center gap-3 justify-center items-center mt-4">
                        {{ key + 1 }} : {{ value.title }}

                        <button style="background-color: #ff0000;" @click="() => handleDeleteTodo(value.id)">
                            <FontAwesomeIcon :icon="faTrash" style="color: #fff; font-size: 17px;" />
                        </button>

                        <button style="background-color: #1a49c8;" @click="() => handleUpdateTodo(value.id , value.done)">
                            <FontAwesomeIcon :icon="faPencil" style="color: #fff; font-size: 17px;" />
                        </button>

                        <button style="background-color: #f6bf0b;" @click="handleCheckTodo(value.id , value.title)">
                            <FontAwesomeIcon :icon="faCheck" style="color: #fff; font-size: 17px;" />
                        </button>

                </div>
                </li>
            </ul>
        </div>
    </div>
</template>