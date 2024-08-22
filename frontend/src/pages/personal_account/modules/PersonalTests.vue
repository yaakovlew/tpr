<template>
  <div class="flex column g-m">
    <banner-component class="banner-title"> Открытые тесты </banner-component>
    <banner-component v-if="openedTests">
      <q-table
        hide-pagination
        flat
        :rows="openedTests"
        :columns="columns"
        row-key="test_id"
      >
        <template v-slot:body-cell-open="props">
          <q-td :props="props">
            <q-btn
              flat
              label="Открыть тест"
              color="primary"
              @click="
                openTest(
                  props.row.test_id,
                  props.row.closed_date,
                  props.row.minutes_duration
                )
              "
            />
          </q-td>
        </template>
      </q-table>
    </banner-component>
    <banner-component v-else class="banner-title">
      Нет открытых тестов
    </banner-component>
    <banner-component v-if="doneTests">
      <q-table
        hide-pagination
        flat
        :rows="doneTests"
        :columns="columnsDone"
        row-key="test_id"
      >
        <template v-slot:body-cell-download="props">
          <q-td :props="props">
            <q-btn
              flat
              label="Скачать отчет"
              color="primary"
              icon="download"
              @click="dowloadTestReport(props.row.test_id)"
            />
          </q-td>
        </template>
        <template #body-cell-mark="props">
          <q-td :props="props">
            {{
              testsMarkStudent.find(
                (test) => props.row?.test_id === test?.test_id
              )?.mark
            }}
          </q-td>
        </template>
      </q-table>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import { useTestsStore } from '../../../stores/test';
import { onMounted, computed } from 'vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import { useCurrentTestStore } from '../../../stores/currentTest';
import { useRouter } from 'vue-router';
import { useMarksStore } from 'src/stores/mark';

const currentTestStore = useCurrentTestStore();
const store = useTestsStore();
const marksStore = useMarksStore();

const openedTests = computed(() => store.openedTestsStudent);
const doneTests = computed(() => store.studentsDoneTests);
const testsMarkStudent = computed(() => marksStore.testsMarkStudent);

const router = useRouter();

const openTest = async (id: number, closeDate: number, duration: number) => {
  await currentTestStore.setCurrentTest(id);
  currentTestStore.setCloseDate(closeDate);
  currentTestStore.setTestDuration(duration);
  currentTestStore.setTestStart(Math.floor(new Date().getTime() / 1000));
  router.push({ name: 'student-test' });
};

const dowloadTestReport = async (id: number, name?: string) => {
  await store.testReport(id, name);
};

const columns = [
  {
    name: 'name',
    field: 'name',
    label: 'Название',
    align: 'left',
    sortable: true,
  },
  {
    name: 'minutes_duration',
    field: 'minutes_duration',
    label: 'Продолжительность',
    align: 'left',
    sortable: true,
  },
  {
    name: 'task_description',
    field: 'task_description',
    label: 'Описание',
    align: 'left',
  },
  {
    name: 'default_mark',
    field: 'default_mark',
    align: 'left',
    label: 'Максильная оценка',
  },
  {
    name: 'open',
  },
];

const columnsDone = [
  {
    name: 'name',
    field: 'name',
    label: 'Название',
    align: 'left',
    sortable: true,
  },
  {
    name: 'minutes_duration',
    field: 'minutes_duration',
    label: 'Продолжительность',
    align: 'left',
    sortable: true,
  },
  {
    name: 'task_description',
    field: 'task_description',
    label: 'Описание',
    align: 'left',
  },
  {
    name: 'mark',
    field: 'mark',
    align: 'left',
    label: 'Оценка',
  },
  {
    name: 'download',
  },
];

onMounted(async () => {
  await store.getOpenedTests();
  await store.getDoneTestsStudent();
  await marksStore.getDoneTestMarksStudent(
    doneTests.value.map((test) => test.test_id)
  );
  // store.testReport(9);
});
</script>

<style lang="scss" scoped></style>
