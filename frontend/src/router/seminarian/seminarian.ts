import PersonalAccountPageVue from 'src/pages/seminarian/personal-account/PersonalAccountPage.vue';
import { RouteRecordRaw, RouterView } from 'vue-router';
import DisciplineGroup from 'src/pages/seminarian/personal-account/modules/DisciplineGroup.vue';
import DicsiplineGroupMarks from 'src/pages/seminarian/personal-account/modules/DicsiplineGroupMarks.vue';

export const seminarianRoutes: RouteRecordRaw = {
  path: '/profile',
  component: RouterView,
  children: [
    {
      path: '',
      name: 'seminarian-profile',
      component: PersonalAccountPageVue,
    },
    {
      path: '/seminarian-group-discipline/:disciplineId/:groupId',
      name: 'seminarian-group-discipline',
      component: DisciplineGroup,
    },
    {
      path: '/seminarian-group-discipline/:disciplineId/:groupId/marks',
      name: 'seminarian-group-discipline-marks',
      component: DicsiplineGroupMarks,
    },
  ],
};
