
// import Register from '@/components/register.vue'
import Register from '@/auth/Register.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    name: 'register',
    path: '/',
    component: Register,
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  linkActiveClass: "text-yellow-500"
})

export default router
