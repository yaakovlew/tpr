import PersonalAccountPage from 'src/pages/personal_account/PersonalAccountPage.vue';
import DisciplineInfo from 'src/pages/personal_account/components/DisciplineInfo.vue';
import PersonalTest from 'src/pages/personal_account/modules/PersonalTest.vue';
import { RouteRecordRaw, RouterView } from 'vue-router';

export const studentRoutes: RouteRecordRaw = {
  path: '/profile',
  component: RouterView,
  children: [
    {
      path: '',
      name: 'student-profile',
      component: PersonalAccountPage,
    },
    {
      path: '/discipline/:disciplineId',
      name: 'student-discipline',
      component: DisciplineInfo,
    },
    {
      path: '/test',
      name: 'student-test',
      component: PersonalTest,
    },
  ],
};
