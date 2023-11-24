import { createRouter, createWebHistory } from 'vue-router'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/container',
      name: 'container',
      component: () => import('../components/PageContainer.vue')
    },
  ]
})

export default router
