import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/email',
      component: () => import('@/layout/DashboardLayout.vue'),
      children: [
        {
          name: 'email',
          path: '/email',
          component: () => import('../views/EmailView.vue'),
        },
      ],
    },
  ],
})

export default router
