<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Тесты
    </banner-component>
    <banner-component v-if="tests" class="text-primary">
      <q-table
        :columns="columns"
        :rows="tests"
        :pagination="pagination"
        hide-bottom
        dense
        flat
      >
        <template v-slot:header-cell="props">
          <q-th :props="props" class="text-primary test">
            {{ props.col.label }}
          </q-th>
        </template>
        <template v-slot:body="props">
          <lecturer-test :test="props.row" />
        </template>
      </q-table>
    </banner-component>
    <banner-component v-else class="text-primary profile-name">
      Тестов нет
    </banner-component>
    <banner-component
      class="text-primary profile-name"
      @click="openCreateTestModal"
    >
      <div class="flex items-center justify-between row cursor-pointer">
        Создать тест
        <q-icon name="add" size="30px" />
      </div>
    </banner-component>
    <banner-component class="text-primary profile-name" @click="goToQuestions">
      <div class="flex items-center justify-between row cursor-pointer">
        Вопросы
        <q-icon name="forward" />
      </div>
    </banner-component>
    <banner-component class="text-primary profile-name" @click="goToThemes">
      <div class="flex items-center justify-between row cursor-pointer">
        Темы
        <q-icon name="forward" />
      </div>
    </banner-component>
  </div>
  <q-dialog v-model="createTestModal">
    <div class="create-test flex column justify-between g-m q-pa-md">
      <q-input v-model="name" label="Название" />
      <q-input type="textarea" v-model="description" label="Описание" />
      <q-input
        type="number"
        v-model="defaultMark"
        label="Максимальная оценка"
      />
      <q-input type="number" v-model="minutes" label="Длительность в минутах" />
      <q-btn
        flat
        color="primary"
        label="Создать тест"
        @click="createTest"
        :disable="!name && !defaultMark && !minutes"
      />
    </div>
  </q-dialog>
</template>

<script lang="ts" setup>
import BannerComponent from '../../../../components/BannerComponent.vue';
import LecturerTest from './components/LecturerTest.vue';
import { useTestsStore } from '../../../../stores/test';
import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const store = useTestsStore();

const tests = computed(() => store.tests);

const router = useRouter();

const columns = [
  {
    name: 'name',
    label: 'Название',
    sortable: true,
    field: 'name',
    align: 'left',
  },
  {
    name: 'task_description',
    label: 'Описание',
    sortable: true,
    field: 'task_description',
    align: 'left',
  },
  ,
  {
    name: 'minutes_duration',
    label: 'Продолжительность',
    field: 'minutes_duration',
    align: 'left',
  },
  {
    name: 'default_mark',
    label: 'Оценка',
    field: 'default_mark',
    align: 'left',
    sortable: false,
  },
  {
    name: 'actions',
    label: '',
    field: 'actions',
    align: 'left',
    sortable: false,
  },
];

onMounted(async () => {
  await store.getAllThemes();
});

const pagination = {
  rowsPerPage: 0,
};

const goToQuestions = () => {
  router.push({ name: 'questions' });
};

const goToThemes = () => {
  router.push({ name: 'themes' });
};

const defaultMark = ref(0);
const minutes = ref(0);
const name = ref('');
const description = ref('');

const createTestModal = ref(false);

const openCreateTestModal = () => {
  createTestModal.value = true;
  defaultMark.value = 0;
  minutes.value = 0;
  name.value = '';
  description.value = '';
};

const createTest = async () => {
  await store.createTest({
    name: name.value,
    task_description: description.value,
    default_mark: Number(defaultMark.value),
    minutes_duration: Number(minutes.value),
    task_description_en: description.value,
    name_en: name.value,
  });
  createTestModal.value = false;
  await store.getTests();
};

onMounted(async () => {
  await store.getTests();
});
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}

.create-test {
  background-color: white;
  width: 500px;
}

.test {
  font-size: 20px;
  font-weight: 500;
}
</style>
