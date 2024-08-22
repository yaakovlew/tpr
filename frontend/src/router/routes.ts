import { RouteRecordRaw } from 'vue-router';

const MainLayout = () => import('layouts/MainLayout.vue');
const ErrorPage = () => import('pages/ErrorNotFound.vue');

import MainPageVue from 'src/pages/main/MainPage.vue';
import AuthPageVue from 'src/pages/auth/AuthPage.vue';
// import CreateDisciplineFormVue from 'src/pages/lecturer/CreateDisciplineForm.vue';
import { studentRoutes } from './student/student';
import { lectruerRoutes } from './lecturer/lecturer';
import { seminarianRoutes } from './seminarian/seminarian';
import AuthRecoverVue from 'src/pages/auth/components/AuthRecover.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: MainLayout,
    children: [{ path: '', component: MainPageVue }],
  },
  {
    path: '/reset-password',
    component: MainLayout,
    children: [{ path: '', component: AuthRecoverVue }],
  },
  {
    path: '/auth',
    component: MainLayout,
    children: [{ path: '', component: AuthPageVue }],
  },
  {
    path: '/student',
    component: MainLayout,
    children: [studentRoutes],
  },
  {
    path: '/lecturer',
    component: MainLayout,
    children: [lectruerRoutes],
  },
  {
    path: '/seminarian',
    component: MainLayout,
    children: [seminarianRoutes],
  },
  {
    path: '/:catchAll(.*)*',
    component: ErrorPage,
  },
];

export default routes;
