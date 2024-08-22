import PersonalAccountPageVue from 'src/pages/lecturer/personal_account/PersonalAccountPage.vue';
import LecturerAttendanceVue from 'src/pages/lecturer/attendance/LecturerAttendance.vue';
import LecturerDiscipline from 'src/pages/lecturer/personal_account/modules/components/LecturerAddDiscipline.vue';
import LecturerEditDiscipline from 'src/pages/lecturer/personal_account/modules/components/components/LecturerEditDiscipline.vue';
import DisciplineGroup from 'src/pages/lecturer/personal_account/modules/components/components/DisciplineGroup.vue';
import LecturerQuestions from 'src/pages/lecturer/questions/LecturerQuestions.vue';
import LecturerThemes from 'src/pages/lecturer/themes/LecturerThemes.vue';
import DicsiplineGroupMarks from 'src/pages/lecturer/personal_account/modules/components/components/DicsiplineGroupMarks.vue';
import { RouteRecordRaw, RouterView } from 'vue-router';

export const lectruerRoutes: RouteRecordRaw = {
  path: '/profile',
  component: RouterView,
  children: [
    {
      path: '',
      name: 'lecturer-profile',
      component: PersonalAccountPageVue,
    },
    {
      path: '/attendance',
      name: 'lecturer-attendance',
      component: LecturerAttendanceVue,
    },
    {
      path: '/new-discipline',
      name: 'new-discipline',
      component: LecturerDiscipline,
    },
    {
      path: '/edit-discipline/:disciplineId',
      name: 'edit-discipline',
      component: LecturerEditDiscipline,
    },
    {
      path: '/group-discipline/:disciplineId/:groupId',
      name: 'group-discipline',
      component: DisciplineGroup,
    },
    {
      path: '/questions',
      name: 'questions',
      component: LecturerQuestions,
    },
    {
      path: '/themes',
      name: 'themes',
      component: LecturerThemes,
    },
    {
      path: '/group-discipline/:disciplineId/:groupId/marks',
      name: 'group-discipline-marks',
      component: DicsiplineGroupMarks,
    },
  ],
};
