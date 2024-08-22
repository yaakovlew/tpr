<template>
  <div class="q-px-md flex column g-m">
    <banner-with-back class="text-primary profile-name">
      Темы
    </banner-with-back>
    <banner-component>
      <q-list class="q-pa-none" separator>
        <q-item class="flex row items-center justify-between">
          <div class="question text-primary">Название темы</div>
          <div class="flex g-m items-center">
            <q-icon name="list" color="transparent" size="24px" />
            <q-icon name="delete" color="transparent" size="24px" />
          </div>
        </q-item>
        <lecturer-theme
          v-for="theme in allThemes"
          :key="theme.theme_id"
          :theme="theme"
        />
      </q-list>
    </banner-component>
    <banner-component
      class="text-primary profile-name flex row cursor-pointer"
      @click="openNewTheme"
    >
      Добавить тему <q-icon name="add" color="primary" size="28px" />
    </banner-component>
    <q-dialog v-model="addNewThemeModal">
      <div class="new-question q-pa-md flex column g-m no-wrap">
        <div class="flex row g-m items-center justify-between">
          <q-input label="Название темы" v-model="newTheme" />
          <div style="min-width: 45%">
            <div class="flex column g-m" v-if="allQuestions?.length">
              <!-- <div class="text-h6">Добавить вопрос</div> -->
              <q-select
                v-model="selectedQuestion"
                :options="allQuestions"
                option-label="question"
                option-value="question_id"
                label="Вопрос"
              />
              <q-btn
                v-if="selectedQuestion"
                label="Добавить вопрос"
                color="primary"
                flat
                @click="addQuestion"
              />
            </div>
            <div
              class="flex column g-m text-primary"
              v-if="selectedQuestions?.length"
            >
              <div class="text-h6">Добавленные вопросы</div>
              <div
                v-for="question in selectedQuestions"
                :key="question.question_id"
              >
                {{ question.question }}
              </div>
            </div>
          </div>
        </div>
        <q-btn
          flat
          color="primary"
          :disable="!newTheme"
          label="Добавить тему"
          @click="addNewTheme"
        />
      </div>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import BannerWithBack from 'src/components/BannerWithBack.vue';
import { useTestsStore } from '../../../stores/test';
import LecturerTheme from './LecturerTheme.vue';
import { computed, onMounted, Ref, ref, watch } from 'vue';
import { IQuestion } from 'src/models/test/question';

const store = useTestsStore();

const allThemes = computed(() => store.allThemes);
// const questions = computed(() => store.allQuestions);

const addQuestion = () => {
  selectedQuestions.value?.push(selectedQuestion.value as IQuestion.Question);
  selectedQuestion.value = undefined;
};

const selectedQuestion: Ref<IQuestion.Question | undefined> = ref(undefined);
const selectedQuestions: Ref<IQuestion.Question[] | null> = ref([]);

const allQuestions = computed(() =>
  store.allQuestions.filter(
    (question) =>
      !selectedQuestions.value?.find(
        (q) => q.question_id === question.question_id
      )
  )
);

const newTheme = ref('');
const weight = ref(1);

const addNewThemeModal = ref(false);

const openNewTheme = () => {
  addNewThemeModal.value = true;
};

const closeNewTheme = () => {
  newTheme.value = '';
  weight.value = 1;
  selectedQuestions.value = [];
  selectedQuestion.value = undefined;
};

watch(addNewThemeModal, () => {
  if (!addNewThemeModal.value) {
    closeNewTheme();
  }
});

const addNewTheme = async () => {
  await store.createTheme({
    name: newTheme.value,
    weight: Number(weight.value),
  });
  await store.getAllThemes();
  const addedTheme = allThemes.value.find(
    (theme) => theme.name === newTheme.value
  );
  const questionPromises: Promise<void>[] = [];
  selectedQuestions.value?.forEach((question) => {
    if (!addedTheme) return;
    questionPromises.push(
      store.addQuestion({
        question_id: question.question_id,
        theme_id: addedTheme?.theme_id,
      })
    );
  });
  await Promise.all(questionPromises);
  addNewThemeModal.value = false;
};

onMounted(async () => {
  await store.getAllThemes();
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

.question {
  font-size: 18px;
  font-weight: 500;
  width: 250px;
}
</style>
