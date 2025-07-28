import { createApp } from 'vue'
import App from './App.vue'
import './index.css';

import { createRouter, createWebHistory } from 'vue-router';
import Home from './pages/Home.vue';
import Login from './pages/Login.vue';
import Signup from './pages/Signup.vue';
import Dashboard from './pages/Dashboard.vue';

import axios from 'axios';
import Cookie from 'js-cookie'

const routes = [
  { path: '/', component: Home },
  { path: '/login' , component: Login },
  { path: '/signup' , component: Signup},
  { path : '/dashboard' , component: Dashboard , meta: {requiresAuth : true}}
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to, from, next) => {
  if (!to.meta.requiresAuth) {
    return next()
  }

  const jwt_token = Cookie.get("jwt");
  if (!jwt_token) return next("/login")

  try {
    const res = await axios.get("http://127.0.0.1:4000/todos/check-token", {
      headers: {
        Authorization: `Bearer ${jwt_token}`
      }
    })

    if (res.status === 200) {
      next()
    } else {
      next("/login")
    }
  } catch (err) {
    next("/login")
  }
})

createApp(App).use(router).mount('#app')