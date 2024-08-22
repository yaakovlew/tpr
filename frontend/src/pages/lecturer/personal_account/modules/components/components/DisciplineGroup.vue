<template>
  <div class="q-px-md flex column g-m">
    <banner-component>
      <template #title>
        <div>
          {{ disciplineInfo?.discipline_name }}
        </div>
        <div>Группа: {{ currentGroupName }}</div>
      </template>
    </banner-component>
    <banner-component>
      <template #title> Лекции </template>
      <q-list separator>
        <q-item
          class="flex g-m items-center justify-between"
          v-for="lesson in groupLessons"
          :key="lesson.lesson_id"
        >
          <discipline-group-lesson
            :lesson="lesson"
            :discipline-id="disciplineIdNumber"
            :group-id="groupIdNumber"
          />
        </q-item>
      </q-list>
      <q-btn
        label="Добавить лекцию"
        color="primary"
        @click="openAddLessonModal"
        class="self-end"
      />
      <q-dialog v-model="addSeminarModal" @hide="closeAddSeminarModal">
        <div class="flex column g-m modal">
          <q-input v-model="seminarName" label="Название" />
          <div class="flex g-m">
            <q-date
              v-model="seminarDate"
              :mask="mask"
              color="purple"
              :locale="myLocale"
            />
            <q-time
              v-model="seminarDate"
              :mask="mask"
              color="purple"
              format24h
            />
            <q-btn
              label="Добавить"
              @click="createSeminar"
              :disable="seminarName === ''"
            />
          </div>
        </div>
      </q-dialog>
    </banner-component>
    <banner-component>
      <template #title> Семинары </template>
      <q-list separator>
        <q-item
          class="flex g-m items-center justify-between"
          v-for="seminar in sortedSeminars"
          :key="seminar.seminar_id"
        >
          <discipline-group-element
            :seminar="seminar"
            :discipline-id="disciplineIdNumber"
            :group-id="groupIdNumber"
          />
        </q-item>
      </q-list>
      <q-btn
        label="Добавить семинар"
        color="primary"
        @click="openAddSeminarModal"
        class="self-end"
      />
      <q-dialog v-model="addLessonModal" @hide="closeAddLessonModal">
        <div class="flex column g-m modal">
          <q-select
            v-model="lessonName"
            :options="avaliableToAddLessons"
            option-label="name"
            option-value="id"
          />
          <div class="flex g-m">
            <q-date
              v-model="lessonDate"
              :mask="mask"
              color="purple"
              :locale="myLocale"
            />
            <q-time
              v-model="lessonDate"
              :mask="mask"
              color="purple"
              format24h
            />
            <q-btn label="Добавить" @click="addLesson" :disable="!lessonName" />
          </div>
        </div>
      </q-dialog>
      <q-dialog v-model="addSeminarModal" @hide="closeAddSeminarModal">
        <div class="flex column g-m modal">
          <q-input v-model="seminarName" label="Название" />
          <div class="flex g-m">
            <q-date
              v-model="seminarDate"
              :mask="mask"
              color="purple"
              :locale="myLocale"
            />
            <q-time
              v-model="seminarDate"
              :mask="mask"
              color="purple"
              format24h
            />
            <q-btn
              label="Добавить"
              @click="createSeminar"
              :disable="seminarName === ''"
            />
          </div>
        </div>
      </q-dialog>
    </banner-component>
    <banner-component>
      <template #title> Разделы </template>
      <q-list separator class="text-primary">
        <q-item
          v-for="section in sections"
          :key="section.section_id"
          class="flex items-center justify-between q-px-none section-name"
        >
          <div>
            {{ section.name }}
          </div>
          <q-icon
            name="list"
            @click="getSectionTests(section.section_id)"
            size="26px"
            class="cursor-pointer"
          />
        </q-item>
      </q-list>
    </banner-component>
    <banner-component>
      <template #title>
        Оценки
        <q-icon
          name="forward"
          class="cursor-pointer"
          @click="goToMarks"
          color="primary"
          size="18px"
        />
      </template>
    </banner-component>
  </div>
  <q-dialog v-model="testsModal">
    <div class="add-group-modal flex column g-m q-pa-lg test-modal-tests">
      <div class="text-primary test-title text-bolder">Тесты</div>
      <q-list separator>
        <q-item
          v-for="test in tests"
          :key="test.test_id"
          class="text-primary flex column g-m q-px-none"
        >
          <div class="dialog-title">Название: {{ test.name }}</div>
          <div class="dialog-title">Описание: {{ test.task_description }}</div>
          <div class="dialog-title">Оценка: {{ test.default_mark }}</div>
          <div class="dialog-title">
            Длительность: {{ test.minutes_duration }}
          </div>
          <q-btn
            label="Открыть тест"
            color="primary"
            flat
            @click="openOpenTestModal(test.test_id)"
          />
        </q-item>
      </q-list>
    </div>
  </q-dialog>
  <q-dialog
    v-model="openTestModal"
    transition-show="scale"
    transition-hide="scale"
  >
    <div class="test-modal">
      <q-tabs
        v-model="tab"
        dense
        class="text-grey"
        active-color="primary"
        indicator-color="primary"
        align="justify"
        narrow-indicator
      >
        <q-tab name="open" label="Открыть" />
        <q-tab name="close" label="Закрыть" />
      </q-tabs>
      <q-tab-panels v-model="tab" animated>
        <q-tab-panel name="open">
          <div class="flex g-m justify-between">
            <q-list separator style="flex-grow: 100">
              <q-btn
                flat
                color="primary"
                label="Отметить всех студентов"
                @click="checkAllStudents(studentsClosedModel)"
                v-if="
                  studentsWithClosedTest && studentsWithClosedTest.length > 0
                "
              />
              <q-item
                v-for="student in studentsWithClosedTest"
                :key="student.student_id"
                class="q-px-none full-width"
              >
                <div
                  class="text-primary full-width flex items-center justify-between"
                >
                  {{ student.name + ' ' + student.surname }}
                  <q-checkbox
                    v-model="studentsClosedModel[Number(student.student_id)]"
                  />
                </div>
              </q-item>
            </q-list>
            <div class="flex column g-m">
              <div class="flex g-m">
                <q-date
                  v-model="testDate"
                  :mask="mask"
                  color="purple"
                  :locale="myLocale"
                />
                <q-time
                  v-model="testDate"
                  :mask="mask"
                  color="purple"
                  format24h
                />
              </div>
              <q-btn label="Открыть" @click="createTest" />
            </div>
          </div>
        </q-tab-panel>

        <q-tab-panel name="close">
          <div class="flex column g-m justify-between">
            <q-list separator>
              <q-btn
                v-if="
                  studentsWithOpenedTest && studentsWithOpenedTest.length > 0
                "
                flat
                color="primary"
                label="Отметить всех студентов"
                @click="checkAllStudents(studentsOpenedModel)"
              />
              <q-item
                v-for="student in studentsWithOpenedTest"
                :key="student.student_id"
                class="q-px-none full-width"
              >
                <div
                  class="text-primary full-width flex items-center justify-between"
                >
                  {{ student.name + ' ' + student.surname }}
                  <q-checkbox
                    v-model="studentsOpenedModel[Number(student.student_id)]"
                  />
                </div>
              </q-item>
            </q-list>

            <q-btn label="Закрыть" @click="closeTest" />
          </div>
        </q-tab-panel>
      </q-tab-panels>
    </div>
  </q-dialog>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from 'vue-router';
import {
  computed,
  watch,
  onMounted,
  ComputedRef,
  ref,
  Ref,
  reactive,
} from 'vue';
import { useDisciplinesStore } from '../../../../../../stores/disciplines';
import { useAttendanceStore } from 'src/stores/attendance';
import DisciplineGroupElement from './DisciplineGroupElement.vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import { ISeminar, ILesson } from 'src/models/attendance/attendance';
import { useReportStore } from '../../../../../../stores/report';
import { useGroupsStore } from 'src/stores/groups';
import DisciplineGroupLesson from './DisciplineGroupLesson.vue';
import { useSectionStore } from 'src/stores/section';
import { useTestsStore } from 'src/stores/test';

const route = useRoute();

const tab = ref('open');

const disciplineId = computed(() => route.params.disciplineId);
const groupId = computed(() => route.params.groupId);

const disciplineStore = useDisciplinesStore();
const groupStore = useGroupsStore();

const groups = computed(() => groupStore.groups);
const currentGroupName = ref('');

const disciplineInfo = computed(() => disciplineStore.discipline);

watch(disciplineId, () => {
  if (!Array.isArray(disciplineId.value)) {
    disciplineStore.getDisciplineInfo(disciplineId.value);
  }
});

watch(groupId, async () => {
  if (!Array.isArray(groupId.value)) {
    await groupStore.getGroups();
    const gr = groups.value?.find(
      (g) => String(g.group_id) === String(groupId.value)
    );
    if (gr) {
      currentGroupName.value = gr.name;
    }
  }
});

onMounted(async () => {
  if (!Array.isArray(groupId.value)) {
    await groupStore.getGroups();
    const gr = groups.value?.find(
      (g) => String(g.group_id) === String(groupId.value)
    );
    if (gr) {
      currentGroupName.value = gr.name;
    }
  }
});

if (!Array.isArray(disciplineId.value)) {
  disciplineStore.getDisciplineInfo(disciplineId.value);
}

const sectionStore = useSectionStore();
sectionStore.getSections(Number(disciplineId.value));

const sections = computed(() => sectionStore.sections);

const testsStore = useTestsStore();

const tests = computed(() => testsStore.sectionTests);

const testsModal = ref(false);

const selectedTest = ref(0);

const getSectionTests = async (id: number) => {
  testsModal.value = true;
  await testsStore.getSectionTests(id);
};

const openTestModal = ref(false);

const sortedSeminars: ComputedRef<ISeminar[]> = computed(() => {
  if (seminars.value) {
    return JSON.parse(JSON.stringify(seminars.value)).sort(
      (a: ISeminar, b: ISeminar) => a.date - b.date
    );
  }
  return [];
});

const disciplineIdNumber = computed(() => {
  if (!Array.isArray(route.params.disciplineId)) {
    return Number(route.params.disciplineId);
  }
  return 0;
});

const groupIdNumber = computed(() => {
  if (!Array.isArray(route.params.groupId)) {
    return Number(route.params.groupId);
  }
  return 0;
});

onMounted(async () => {
  await groupStore.getGroupStudents(String(groupIdNumber.value));
});

const attendanceStore = useAttendanceStore();

const seminars = computed(() => attendanceStore.seminars);

const allDisciplineLessons = computed(
  () => attendanceStore.allDisciplineLessons
);

const groupLessons = computed(() => {
  const copy = ref(attendanceStore.groupLessons);
  copy.value?.filter((groupLesson) =>
    allDisciplineLessons.value?.find(
      (lesson) => lesson.lesson_id === groupLesson.lesson_id
    )
  );
  return copy.value;
});

onMounted(async () => {
  if (!Array.isArray(disciplineId.value) && !Array.isArray(groupId.value)) {
    await disciplineStore.getDisciplineInfo(disciplineId.value);
    await attendanceStore.getDisciplineLessons(Number(disciplineId.value));
    await attendanceStore.getSeminars(
      Number(disciplineId.value),
      Number(groupId.value)
    );
    await attendanceStore.getGroupLessons(
      groupIdNumber.value,
      disciplineIdNumber.value
    );
  }
});

const avaliableToAddLessons = computed(() =>
  allDisciplineLessons.value?.filter(
    (lesson) =>
      !groupLessons.value?.find(
        (addedLesson) => lesson.lesson_id === addedLesson.lesson_id
      )
  )
);

const router = useRouter();

const goToMarks = () => {
  router.push({
    name: 'group-discipline-marks',
    params: { disciplineId: disciplineId.value, groupId: groupId.value },
  });
};

const addSeminarModal = ref(false);
const addLessonModal = ref(false);

const openAddSeminarModal = () => {
  addSeminarModal.value = true;
};

const openAddLessonModal = () => {
  addLessonModal.value = true;
};

// attendanceStore.addLessonDate({
//   group_id: groupIdNumber.value,
//   lesson_id: 1,
//   date: new Date().getTime(),
// });

const closeAddSeminarModal = () => {
  seminarName.value = '';
  seminarDate.value = getDate();
  addSeminarModal.value = false;
};

const closeAddLessonModal = () => {
  lessonName.value = null;
  lessonDate.value = getDate();
  addLessonModal.value = false;
};

const closeOpenTestModal = () => {
  openTestModal.value = false;
  testDate.value = getDate();
};

const getDate = () => {
  const newDate = new Date();
  newDate.setSeconds(0, 0);
  return newDate.toLocaleString('ru');
};

const seminarName = ref('');
const lessonName: Ref<ILesson | null> = ref(null);
const seminarDate = ref(getDate());
const lessonDate = ref(getDate());
const testDate = ref(getDate());

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

const addLesson = async () => {
  const dateParse = new Date();
  dateParse.setDate(Number(lessonDate.value[0] + lessonDate.value[1]));
  dateParse.setMonth(Number(lessonDate.value[3] + lessonDate.value[4]) - 1);
  dateParse.setFullYear(
    Number(
      lessonDate.value[6] +
        lessonDate.value[7] +
        lessonDate.value[8] +
        lessonDate.value[9]
    )
  );
  dateParse.setHours(
    Number(lessonDate.value[12] + lessonDate.value[13]),
    Number(lessonDate.value[15] + lessonDate.value[16]),
    Number(lessonDate.value[18] + lessonDate.value[19])
  );
  if (lessonName.value?.lesson_id) {
    await attendanceStore.addLessonDate({
      group_id: groupIdNumber.value,
      lesson_id: Number(lessonName.value?.lesson_id),
      date: Math.floor(dateParse.getTime() / 1000),
    });
  }
  await attendanceStore.getGroupLessons(
    groupIdNumber.value,
    disciplineIdNumber.value
  );
  closeAddLessonModal();
};

const createSeminar = async () => {
  const dateParse = new Date();
  dateParse.setDate(Number(seminarDate.value[0] + seminarDate.value[1]));
  dateParse.setMonth(Number(seminarDate.value[3] + seminarDate.value[4]) - 1);
  dateParse.setFullYear(
    Number(
      seminarDate.value[6] +
        seminarDate.value[7] +
        seminarDate.value[8] +
        seminarDate.value[9]
    )
  );
  dateParse.setHours(
    Number(seminarDate.value[12] + seminarDate.value[13]),
    Number(seminarDate.value[15] + seminarDate.value[16]),
    Number(seminarDate.value[18] + seminarDate.value[19])
  );
  await attendanceStore.createSeminar({
    date: Math.floor(dateParse.getTime() / 1000),
    discipline_id: disciplineIdNumber.value,
    group_id: groupIdNumber.value,
    name: seminarName.value,
  });
  await attendanceStore.getSeminars(
    Number(disciplineId.value),
    Number(groupId.value)
  );
  closeAddSeminarModal();
};

const studentsOpenedTest = computed(() => testsStore.studentsOpenedTest);
const groupStudents = computed(() => groupStore.groupStudents);

const studentsClosedModel: Record<number, boolean> = reactive({});
const studentsOpenedModel: Record<number, boolean> = reactive({});

const studentsWithClosedTest = computed(() => {
  const res = groupStudents.value?.filter(
    (student) =>
      !studentsOpenedTest.value?.find(
        (st) => Number(student.student_id) === st.student_id
      )
  );
  res?.forEach((student) => {
    studentsClosedModel[Number(student.student_id)] = false;
  });
  return res;
});

const studentsWithOpenedTest = computed(() => {
  const res = studentsOpenedTest.value?.filter((student) =>
    groupStudents.value?.find(
      (st) => Number(st.student_id) === student.student_id
    )
  );
  res?.forEach((student) => {
    studentsOpenedModel[Number(student.student_id)] = false;
  });
  return res;
});

const checkAllStudents = (students: Record<number, boolean>) => {
  Object.keys(students).forEach((key) => {
    students[Number(key)] = !students[Number(key)];
  });
};

const openOpenTestModal = async (testId: number) => {
  await testsStore.getStudentsOpenTest(testId);
  await groupStore.getGroupStudents(String(groupId.value));

  selectedTest.value = testId;
  openTestModal.value = true;
};

const createTest = async () => {
  const dateParse = new Date();
  dateParse.setDate(Number(testDate.value[0] + testDate.value[1]));
  dateParse.setMonth(Number(testDate.value[3] + testDate.value[4]) - 1);
  dateParse.setFullYear(
    Number(
      testDate.value[6] +
        testDate.value[7] +
        testDate.value[8] +
        testDate.value[9]
    )
  );
  dateParse.setHours(
    Number(testDate.value[12] + testDate.value[13]),
    Number(testDate.value[15] + testDate.value[16]),
    Number(testDate.value[18] + testDate.value[19])
  );
  const studentsIds: number[] = [];
  Object.entries(studentsClosedModel).forEach((value) =>
    value[1] ? studentsIds.push(Number(value[0])) : null
  );
  await groupStore.openTestForStudents(
    studentsIds,
    Number(selectedTest.value),
    Math.floor(dateParse.getTime() / 1000)
  );
  await testsStore.getStudentsOpenTest(Number(selectedTest.value));

  closeOpenTestModal();
};

const closeTest = async () => {
  const studentsIds: number[] = [];
  Object.entries(studentsOpenedModel).forEach((value) =>
    value[1] ? studentsIds.push(Number(value[0])) : null
  );
  await groupStore.closeTestForStudents(
    studentsIds,
    Number(selectedTest.value)
  );
  await testsStore.getStudentsOpenTest(Number(selectedTest.value));

  closeOpenTestModal();
};

watch(disciplineId, async () => {
  if (!Array.isArray(disciplineId.value) && disciplineId.value) {
    await disciplineStore.getDisciplineInfo(disciplineId.value);
    await attendanceStore.getDisciplineLessons(Number(disciplineId.value));
  }
});
</script>

<style lang="scss" scoped>
.modal {
  background-color: white;
  min-width: 650px;
  padding: 10px 20px;
}

.test-modal {
  background-color: white;
  min-width: 950px;
  padding: 10px 20px;
}

.test-modal-test {
  min-width: 500px;
}

@media screen and (max-width: 600px) {
  .test-modal {
    width: 90vw;
    min-width: 0;
  }

  .test-modal-test {
    width: 90vw;
    min-width: none;
  }
}

.add-group-modal {
  width: 300px;
  background-color: white;
  padding: 10px;
}

.dialog-title {
  font-weight: 500;
  font-size: 18px;
}

.test-title {
  font-weight: 600;
  font-size: 20px;
}

.section-name {
  font-weight: 500;
  font-size: 18px;
}
</style>
