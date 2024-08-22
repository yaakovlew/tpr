<template>
  <q-item class="flex row items-center justify-between">
    <div class="question text-primary">
      {{ theme.name }}
      <q-popup-edit v-model="themeName" v-slot="scope">
        <q-input
          autofocus
          dense
          v-model="scope.value"
          :model-value="scope.value"
        >
          <template v-slot:after>
            <q-btn
              flat
              dense
              color="negative"
              icon="cancel"
              @click.stop.prevent="scope.cancel"
            />

            <q-btn
              flat
              dense
              color="positive"
              icon="check_circle"
              @click.stop.prevent="
                changeThemeName(scope.value, theme.weight);
                scope.set();
              "
              :disable="scope.value === ''"
            />
          </template>
        </q-input>
      </q-popup-edit>
    </div>
    <div class="flex g-m items-center">
      <q-icon
        name="list"
        class="cursor-pointer"
        @click="getThemeQuestions"
        color="primary"
        size="24px"
      />
      <q-icon
        name="delete"
        color="negative"
        size="24px"
        class="cursor-pointer"
        @click="deleteModal = true"
      />
    </div>
  </q-item>
  <q-dialog v-model="questionsModal">
    <div class="modal flex g-m text-primary q-pa-md column no-wrap">
      <div class="title">Вопросы</div>
      <q-list separator class="q-px-none">
        <q-item
          v-for="question in questions"
          :key="question.question_id"
          class="flex justify-between q-px-none"
        >
          <div>
            {{ question.question }}
          </div>
          <q-icon
            class="cursor-pointer"
            color="negative"
            size="18px"
            name="delete"
            @click="deleteQuestion(question.question_id)"
          />
        </q-item>
      </q-list>
      <q-btn
        v-if="!addQuestionsForm && allQuestions"
        label="Добавить вопрос"
        flat
        color="primary"
        @click="addQuestionsForm = true"
      />
      <div v-else-if="allQuestions" class="flex column g-m">
        <q-select
          v-model="selectedQuestion"
          :options="allQuestions"
          option-label="question"
          option-value="question_id"
        />
        <q-btn
          label="Добавить вопрос"
          flat
          color="primary"
          :disable="!selectedQuestion"
          @click="addQuestion"
        />
      </div>
    </div>
  </q-dialog>
  <q-dialog v-model="deleteModal">
    <div class="delete-modal flex g-m text-primary q-pa-md">
      <div>Вы уверены что хотите удалить тему: "{{ theme.name }}"?</div>
      <div class="flex g-m justify-between self-end">
        <q-btn flat label="Отмена" @click="deleteModal = false" />
        <q-btn flat label="Удалить" @click="deleteTheme" />
      </div>
    </div>
  </q-dialog>
</template>

<script lang="ts" setup>
import { ITheme } from 'src/models/test/theme';
import { useTestsStore } from '../../../stores/test';
import { Ref, computed, ref, watch } from 'vue';
import { IQuestion } from 'src/models/test/question';

const props = defineProps<{
  theme: ITheme.Theme;
}>();

const themeName = ref(props.theme.name);
const themeWeight = ref(props.theme.weight);

const store = useTestsStore();

const deleteModal = ref(false);

const changeThemeName = async (name: string, weight: number) => {
  await store.changeTestThemeName({
    name,
    theme_id: props.theme.theme_id,
    weight: Number(weight),
  });
  await store.getAllThemes();
};

const deleteTheme = async () => {
  await store.deleteTheme(props.theme.theme_id);
  await store.getAllThemes();
  deleteModal.value = false;
};

const questionsModal = ref(false);

const questions = computed(() => store.questions);

const getThemeQuestions = async () => {
  await store.getQuestions(props.theme.theme_id);
  questionsModal.value = true;
};

const allQuestions = computed(() =>
  store.allQuestions.filter(
    (question) =>
      !questions.value?.find((q) => q.question_id === question.question_id)
  )
);

store.getAllQuestions();

const selectedQuestion: Ref<IQuestion.Question | null> = ref(null);

const addQuestion = async () => {
  if (selectedQuestion.value) {
    await store.addQuestion({
      question_id: selectedQuestion.value?.question_id,
      theme_id: props.theme.theme_id,
    });
  }
  selectedQuestion.value = null;
  await store.getQuestions(props.theme.theme_id);
  await store.getAllQuestions();
  addQuestionsForm.value = false;
};

const addQuestionsForm = ref(false);

watch(addQuestionsForm, async () => {
  if (addQuestionsForm.value) {
    await store.getAllQuestions();
  }
});

const deleteQuestion = async (id: number) => {
  await store.deleteQuestionFromTheme({
    question_id: id,
    theme_id: props.theme.theme_id,
  });
  await store.getQuestionsWithoutTheme();
  await store.getQuestions(props.theme.theme_id);
};
</script>

<style lang="scss" scoped>
.question {
  font-size: 18px;
  font-weight: 500;
  width: 250px;
}

.modal {
  background-color: white;
  min-width: 500px;
}

.title {
  font-size: 18px;
  font-weight: 500;
}

.answer {
  font-weight: 500;
}

.is-right {
  font-size: 16px;
}

.delete-modal {
  font-size: 20px;
  font-weight: 500;
  width: 400px;
  background-color: white;
}
</style>
