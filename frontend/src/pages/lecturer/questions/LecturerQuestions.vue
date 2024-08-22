<template>
  <div class="q-px-md flex column g-m">
    <banner-with-back class="text-primary profile-name">
      Вопросы
    </banner-with-back>
    <banner-component
      class="text-primary profile-name cursor-pointer"
      @click="importQuestionsModal = true"
    >
      Загрузить вопросы
    </banner-component>
    <banner-component style="width: 100%">
      <q-table
        flat
        title="Вопросы"
        :rows="questions"
        :columns="columns"
        row-key="name"
        :pagination="pagination"
        hide-bottom
        style="width: 100%"
      >
        <template v-slot:body="props">
          <lecturer-question :question="props.row" />
        </template>
      </q-table>
    </banner-component>
    <banner-component
      class="text-primary profile-name flex row cursor-pointer"
      @click="openNewQuestion"
    >
      Добавить вопрос <q-icon name="add" color="primary" size="28px" />
    </banner-component>
    <q-dialog v-model="addNewQuestionModal">
      <div class="new-question q-pa-md flex column g-m">
        <q-input label="Вопрос" v-model="question" />
        <q-select v-model="variable" :options="variableVariants" />
        <q-btn
          flat
          color="primary"
          :disable="!question"
          label="Добавить вопрос"
          @click="addNewQuestion"
        />
      </div>
    </q-dialog>
    <q-dialog v-model="importQuestionsModal">
      <div class="new-question q-pa-md flex column g-m">
        <q-file v-model="file" label="Выберите файл" filled />
        <q-btn
          flat
          color="primary"
          :disable="!file"
          label="Добавить вопросы"
          @click="importQuestions"
        />
      </div>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import BannerWithBack from 'src/components/BannerWithBack.vue';
import LecturerQuestion from './LecturerQuestion.vue';
import { useTestsStore } from '../../../stores/test';
import { Ref, computed, onMounted, ref, watch } from 'vue';

const pagination = {
  rowsPerPage: 0,
};

const importQuestionsModal = ref(false);

const store = useTestsStore();

const file: Ref<File | null> = ref(null);

const questions = computed(() => store.allQuestions);

const importQuestions = async () => {
  if (file.value) {
    let formData = new FormData();
    formData.append('file', file.value);
    await store.importQuestions({
      file: formData,
    });
  }
  importQuestionsModal.value = false;
  await store.getAllQuestions();
};

const columns = ref([
  {
    name: 'question',
    label: 'Вопрос',
    field: 'question',
    align: 'left',
    style: 'width: 500px',
    headerStyle: 'width: 500px',
  },
  {
    name: 'variable',
    label: 'Тип вопроса',
    field: 'variable',
    align: 'left',
  },
  {
    name: 'theme',
    label: 'Тема',
    field: 'theme',
    align: 'left',
  },
  {
    name: 'answers',
    label: 'Ответы',
    field: 'answers',
    align: 'left',
  },
  {
    name: 'actions',
    label: 'Действия',
    field: 'actions',
    align: 'left',
  },
]);

const variableVariants = [
  {
    value: 0,
    label: 'Свободный выбор ответа',
  },
  {
    value: 1,
    label: 'Один вариант',
  },
  {
    value: 2,
    label: 'Множественный вариант',
  },
];

const question = ref('');
const variable = ref(variableVariants[0]);

const addNewQuestionModal = ref(false);

const openNewQuestion = () => {
  addNewQuestionModal.value = true;
};

const closeNewQuestion = () => {
  question.value = '';
  variable.value = variableVariants[0];
};

watch(addNewQuestionModal, () => {
  if (!addNewQuestionModal.value) {
    closeNewQuestion();
  }
});

const addNewQuestion = async () => {
  await store.createQuestion({
    question: question.value,
    question_en: question.value,
    is_variable: variable.value.value as 0 | 1 | 2,
  });
  await store.getAllQuestions();
  addNewQuestionModal.value = false;
};

onMounted(async () => {
  await store.getAllQuestions();
});
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}

.new-question {
  min-width: 500px;
  background-color: white;
}
</style>
