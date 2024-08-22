<template>
  <div class="seminar-name fixed-width text-primary">
    {{ lesson.lesson_name }}
  </div>
  <div class="seminar-name text-primary">
    {{ date }}
  </div>
  <div class="flex g-m">
    <q-icon
      name="person"
      color="primary"
      size="20px"
      class="cursor-pointer"
      @click="openGroupAttendanceModal"
    />
    <q-icon
      name="edit"
      color="primary"
      size="20px"
      class="cursor-pointer"
      @click="openEditModal"
    />
    <q-icon
      name="delete"
      color="primary"
      size="20px"
      class="cursor-pointer"
      @click="deleteSeminar"
    />
  </div>
  <q-dialog v-model="editModal">
    <div class="flex column g-m modal">
      <div class="flex g-m">
        <q-date
          v-model="dateForInput"
          :mask="mask"
          color="purple"
          :locale="myLocale"
        />
        <q-time v-model="dateForInput" :mask="mask" color="purple" format24h />
        <q-btn label="Изменить" @click="changeLesson" />
      </div>
    </div>
  </q-dialog>
  <q-dialog v-model="groupAttendanceModal">
    <div class="flex column g-m modal">
      <discipline-group-lesson-attendance
        class="flex g-m"
        v-for="student in groupStudents"
        :key="student.student_id"
        :student="student"
        :lesson-id="lesson.lesson_id"
      />
    </div>
  </q-dialog>
</template>

<script lang="ts" setup>
import { IAttendance } from 'src/models/attendance/attendance';
import { computed, ref } from 'vue';
import { useAttendanceStore } from 'src/stores/attendance';
import { useGroupsStore } from 'src/stores/groups';
import DisciplineGroupLessonAttendance from './DisciplineGroupLessonAttendance.vue';

const props = defineProps<{
  lesson: IAttendance.LessonDate;
  disciplineId: number;
  groupId: number;
}>();

const store = useAttendanceStore();

const groupStore = useGroupsStore();

const allGroupStudents = computed(() => groupStore.seminarianGroupStudnets);

const groupStudents = computed(() => store.lessonsVisiting);

const groupAttendanceModal = ref(false);

const openGroupAttendanceModal = async () => {
  await store.getLessonsVisitingSeminarian(props.groupId, Number(props.lesson.lesson_id));
  if (groupStudents.value === null && allGroupStudents.value !== null) {
    await Promise.all(
      allGroupStudents.value.map(async (student) => {
        console.log(student);
        await store.addLessonVisiting(
          true,
          Number(props.lesson.lesson_id),
          Number(student.student_id)
        );
      })
    );
    await store.getLessonsVisitingSeminarian(
      props.groupId,
      Number(props.lesson.lesson_id)
    );
  }
  groupAttendanceModal.value = true;
};

const date = computed(() => {
  return getDate(props.lesson.date);
});

const getDate = (date: number) => {
  const newDate = new Date(date * 1000);
  newDate.setSeconds(0, 0);
  return newDate.toLocaleString('ru');
};

const myLocale = {
  /* starting with Sunday */
  days: 'Воскресенье_Понедельник_Вторник_Среда_Четверг_Пятница_Суббота'.split(
    '_'
  ),
  daysShort: 'Вскр_Пон_Втор_Сред_Четв_Пятн_Суб'.split('_'),
  months:
    'Январь_Февраль_Март_Апрель_Май_Июнь_Июль_Август_Сентябрь_Октябрь_Ноябрь_Декабрь'.split(
      '_'
    ),
  monthsShort: 'Янв_Фев_Март_Апр_Май_Июнь_Июль_Авг_Сан_Окт_Нояб_Дек'.split('_'),
  firstDayOfWeek: 1, // 0-6, 0 - Sunday, 1 Monday, ...
  format24h: true,
  pluralDay: 'дни',
};

const mask = 'DD.MM.YYYY, HH:mm:ss';

const dateForInput = ref(date.value);

const changeLesson = async () => {
  const dateParse = new Date();
  dateParse.setDate(Number(dateForInput.value[0] + dateForInput.value[1]));
  dateParse.setMonth(Number(dateForInput.value[3] + dateForInput.value[4]) - 1);
  dateParse.setFullYear(
    Number(
      dateForInput.value[6] +
        dateForInput.value[7] +
        dateForInput.value[8] +
        dateForInput.value[9]
    )
  );
  dateParse.setHours(
    Number(dateForInput.value[12] + dateForInput.value[13]),
    Number(dateForInput.value[15] + dateForInput.value[16]),
    Number(dateForInput.value[18] + dateForInput.value[19])
  );
  await store.changeLessonDate(
    Math.floor(dateParse.getTime() / 1000),
    Number(props.groupId),
    Number(props.lesson.lesson_id)
  );
  await store.getGroupLessons(props.groupId, props.disciplineId);
  editModal.value = false;
};

const editModal = ref(false);

const openEditModal = () => {
  editModal.value = true;
};

const deleteSeminar = async () => {
  await store.deleteLessonDate(Number(props.lesson.lesson_id), props.groupId);
  await store.getGroupLessons(props.groupId, props.disciplineId);
};
</script>

<style lang="scss" scoped>
.seminar-name {
  font-weight: 500;
  font-size: 18px;
}

.modal {
  background-color: white;
  min-width: 650px;
  padding: 10px 20px;
}

.fixed-width {
  width: 150px;
}
</style>
