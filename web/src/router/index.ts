import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../views/home/home.vue'


const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomePage,
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
