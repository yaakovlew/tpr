<template>
    <div class="lab-modal">
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
                  studentsWithClosedLab && studentsWithClosedLab.length > 0
                "
              />
              <q-item
                v-for="student in studentsWithClosedLab"
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
                  v-model="labDate"
                  :mask="mask"
                  color="purple"
                  :locale="myLocale"
                />
                <q-time
                  v-model="labDate"
                  :mask="mask"
                  color="purple"
                  format24h
                />
              </div>
              <q-btn label="Открыть" @click="createLab" />
            </div>
          </div>
        </q-tab-panel>

        <q-tab-panel name="close">
          <div class="flex column g-m justify-between">
            <q-list separator>
              <q-btn
                v-if="
                  studentsWithOpenedLab && studentsWithOpenedLab.length > 0
                "
                flat
                color="primary"
                label="Отметить всех студентов"
                @click="checkAllStudents(studentsOpenedModel)"
              />
              <q-item
                v-for="student in studentsWithOpenedLab"
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

            <q-btn label="Закрыть" @click="closeLab" />
          </div>
        </q-tab-panel>
      </q-tab-panels>
    </div>
</template>


<script lang="ts" setup>
import { useRoute } from 'vue-router';
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
import { ILesson } from 'src/models/attendance/attendance';
import { useGroupsStore } from 'src/stores/groups';
import { useSectionStore } from 'src/stores/section';
import { useLabsStore } from 'src/stores/labs';

const props = defineProps<{
  modelValue: boolean;
}>();

const route = useRoute();

const tab = ref('open');

const disciplineId = computed(() => route.params.disciplineId);
const groupId = computed(() => route.params.groupId);

const disciplineStore = useDisciplinesStore();
const groupStore = useGroupsStore();

const groups = computed(() => groupStore.groups);
const currentGroupName = ref('');

// const disciplineInfo = computed(() => disciplineStore.discipline);

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


const labsStore = useLabsStore();

const selectedLab = ref(0);

const modelValue = ref(props.modelValue);

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

const closeOpenLabModal = () => {
  modelValue.value = false;
  labDate.value = getDate();
};

const getDate = () => {
  const newDate = new Date();
  newDate.setSeconds(0, 0);
  return newDate.toLocaleString('ru');
};

const labDate = ref(getDate());

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

const studentsOpenedLab = computed(() => labsStore.studentsOpenedLab);
const groupStudents = computed(() => groupStore.groupStudents);

const studentsClosedModel: Record<number, boolean> = reactive({});
const studentsOpenedModel: Record<number, boolean> = reactive({});

const studentsWithClosedLab = computed(() => {
  const res = groupStudents.value?.filter(
    (student) =>
      !studentsOpenedLab.value?.find(
        (st) => Number(student.student_id) === st.student_id
      )
  );
  res?.forEach((student) => {
    studentsClosedModel[Number(student.student_id)] = false;
  });
  return res;
});

const studentsWithOpenedLab = computed(() => {
  const res = studentsOpenedLab.value?.filter((student) =>
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

const createLab = async () => {
  const dateParse = new Date();
  dateParse.setDate(Number(labDate.value[0] + labDate.value[1]));
  dateParse.setMonth(Number(labDate.value[3] + labDate.value[4]) - 1);
  dateParse.setFullYear(
    Number(
      labDate.value[6] +
        labDate.value[7] +
        labDate.value[8] +
        labDate.value[9]
    )
  );
  dateParse.setHours(
    Number(labDate.value[12] + labDate.value[13]),
    Number(labDate.value[15] + labDate.value[16]),
    Number(labDate.value[18] + labDate.value[19])
  );
  const studentsIds: number[] = [];
  Object.entries(studentsClosedModel).forEach((value) =>
    value[1] ? studentsIds.push(Number(value[0])) : null
  );
  await groupStore.openLabForStudents(
    studentsIds,
    Number(selectedLab.value),
    Math.floor(dateParse.getTime() / 1000)
  );
  await labsStore.getStudentsOpenLabs(Number(selectedLab.value));

  closeOpenLabModal();
};

const closeLab = async () => {
  const studentsIds: number[] = [];
  Object.entries(studentsOpenedModel).forEach((value) =>
    value[1] ? studentsIds.push(Number(value[0])) : null
  );
  await groupStore.closeLabForStudents(
    studentsIds,
    Number(selectedLab.value)
  );
  await labsStore.getStudentsOpenLabs(Number(selectedLab.value));

  closeOpenLabModal();
};

watch(disciplineId, async () => {
  if (!Array.isArray(disciplineId.value) && disciplineId.value) {
    await disciplineStore.getDisciplineInfo(disciplineId.value);
    await attendanceStore.getDisciplineLessons(Number(disciplineId.value));
  }
});
</script>

<style lang="scss" scoped>
.add-group-modal {
  max-width: 800px;
  margin: auto;
  background-color: white;
  padding: 10px;
}
.flex {
  display: flex;
}
.row {
  flex-direction: row;
}
.column {
  flex-direction: column;
}
.section-title {
  margin-bottom: 16px;
}
.lab-section {
  flex: 1;
  border-right: 1px solid #ccc;
  padding-right: 16px;
}
.lab-section {
  flex: 1;
  padding-left: 16px;
}
.dialog-title {
  margin-bottom: 8px;
}


.modal {
  background-color: white;
  min-width: 650px;
  padding: 10px 20px;
}

.lab-modal {
  background-color: white;
  min-width: 950px;
  padding: 10px 20px;
}

.lab-modal-lab {
  min-width: 500px;
}

@media screen and (max-width: 600px) {
  .lab-modal {
    width: 90vw;
    min-width: 0;
  }

  .lab-modal-lab {
    width: 90vw;
    min-width: none;
  }
}

// .add-group-modal {
//   width: 300px;
//   background-color: white;
//   padding: 10px;
// }

.dialog-title {
  font-weight: 500;
  font-size: 18px;
}

.lab-title {
  font-weight: 600;
  font-size: 20px;
}

.section-name {
  font-weight: 500;
  font-size: 18px;
}
</style>
