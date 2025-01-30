<template>
  <div class="flex column g-m">
    <banner-component class="banner-title"> Открытые Лабораторные работы </banner-component>
    <banner-component v-if="openedLabs">
      <q-table
        hide-pagination
        flat
        :rows="openedLabs"
        :columns="columns"
        row-key="laboratory_id"
      >
        <template v-slot:body-cell-minutes_duration="props">
          <q-td :props="props">
            90 мин
          </q-td>
        </template>
        <template v-slot:body-cell-open="props">
          <q-td :props="props">
            <q-btn
              flat
              label="Открыть"
              color="primary"
              @click="
                openLab(props.row.link)
                // openLab(
                //   props.row.laboratory_id,
                //   props.row.closed_date,
                //   props.row.minutes_duration
                // )
              "
            />
          </q-td>
        </template>
      </q-table>
    </banner-component>
    <banner-component v-else class="banner-title">
      Нет открытых лабораторных работ
    </banner-component>
    <banner-component v-if="doneLabs">
      <q-table
        hide-pagination
        flat
        :rows="doneLabs"
        :columns="columnsDone"
        row-key="laboratory_id"
      >
        <template v-slot:body-cell-download="props">
          <q-td :props="props">
            <!-- <q-btn
              flat
              label="Скачать отчет"
              color="primary"
              icon="download"
              @click="dowloadTestReport(props.row.test_id)"
            /> -->
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
import { useLabsStore } from '../../../stores/labs';
import { onMounted, computed } from 'vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import { useCurrentTestStore } from '../../../stores/currentTest';
import { useRouter } from 'vue-router';
import { useMarksStore } from 'src/stores/mark';

const currentTestStore = useCurrentTestStore();
const store = useLabsStore();
const marksStore = useMarksStore();

const openedLabs = computed(() => store.openedLabsStudent);
const doneLabs = computed(() => store.studentsDoneLabs);
const testsMarkStudent = computed(() => marksStore.testsMarkStudent);

const router = useRouter();

const openLab = async (link: string) => {
  link = `${link}?jwt=${localStorage.getItem('token')}`
  if (link.startsWith('http') || link.startsWith('//')) {
    // Open external link in a new tab
    window.open(link, '_blank');
  } else {
    // Navigate within the app
    router.push(link);
  }



  // await currentTestStore.setCurrentTest(id);
  // currentTestStore.setCloseDate(closeDate);
  // currentTestStore.setTestDuration(duration);
  // currentTestStore.setTestStart(Math.floor(new Date().getTime() / 1000));
  // router.push({ name: 'student-test' });
};

// const dowloadLabReport = async (id: number, name?: string) => {
//   await store.labReport(id, name);
// };

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
  // {
  //   name: 'minutes_duration',
  //   field: 'minutes_duration',
  //   label: 'Продолжительность',
  //   align: 'left',
  //   sortable: true,
  // },
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
  // {
  //   name: 'download',
  // },
];

onMounted(async () => {
  await store.getOpenedLabs();
  await store.getDoneLabsStudent();
  console.log('doneLabs', doneLabs.value)
  // await marksStore.getDoneTestMarksStudent(
  //   doneLabs.value.map((lab) => lab.laboratory_id)
  // );
  // store.testReport(9);
});
</script>

<style lang="scss" scoped></style>
