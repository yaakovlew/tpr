<template>
  <div class="q-pa-md flex g-m column">
    <banner-component class="page-title text-primary flex items-center">
      <countdown
        v-if="testStart && testDuration && closeDate"
        :time="
          (testDuration * 60 -
            (Math.floor(new Date().getTime() / 1000) - testStart)) *
          1000
        "
        @finish="passTest(false)"
      />
    </banner-component>
    <banner-component>
      <q-list separator>
        <q-item
          v-for="(question, index) in test"
          :key="question.questions.question_id"
        >
          <test-question
            v-if="answers[question.questions.question_id]"
            :question="question"
            v-model="answers[question.questions.question_id].answer"
            :question-number="index + 1"
          />
        </q-item>
      </q-list>
    </banner-component>
    <banner-component>
      <div class="flex items-center" style="flex-wrap: nowrap">
        <q-circular-progress
          rounded
          :value="progressValue"
          size="150px"
          color="primary"
          class="q-ma-md"
          label="Общий прогресс"
          track-color="grey-3"
        />
        <div
          class="flex text-primary profile-name text-h6"
          v-if="!isAllAnswered"
          style="height: fit-content"
        >
          Вы не ответили на вопросы под номерами:
          <div class="flex g-m cursor-pointer" style="height: fit-content">
            <div
              v-for="q in questionsWithoutAnswer"
              :key="q.questions.question_id"
              @click="scrollIntoView(q.questions.question_id)"
              style="height: 30px"
            >
              {{ q.number }}
            </div>
          </div>
        </div>
        <div
          class="flex text-primary profile-name text-h6"
          v-else
          style="height: fit-content"
        >
          Вы ответили на все вопросы!
        </div>
      </div>
    </banner-component>
    <banner-component>
      <q-btn
        flat
        color="primary"
        label="Отправить"
        @click="passTest()"
        class="full-width"
      />
    </banner-component>
    <q-dialog v-model="isAllAnsweredModal">
      <banner-component class="text-primary profile-name text-h6">
        <div class="flex column g-xl">
          <div>
            Вы не ответили на вопросы под номерами:
            {{ questionsWithoutAnswer?.map((q) => q.number).join(', ') }}
          </div>
          <q-btn
            label="Отправить все равно"
            @click="passTest(false)"
            flat
            class="full-width"
          />
        </div>
      </banner-component>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import { useCurrentTestStore } from '../../../stores/currentTest';
import { computed, onMounted, reactive } from 'vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import TestQuestion from '../components/TestQuestion.vue';
import { ITest } from 'src/models/test/test';
import { useRouter } from 'vue-router';
import Countdown from 'vue3-countdown';
import { IQuestion } from 'src/models/test/question';
import { ref } from 'vue';

const store = useCurrentTestStore();

const test = computed(() => store.test);
const currentTest = computed(() => store.currentTest);
const closeDate = computed(() => store.closeDate);
const testStart = computed(() => store.testStart);
const testDuration = computed(() => store.testDuration);

const scrollIntoView = (id: number) => {
  setTimeout(() => {
    const element = document.querySelector('#question-' + id);
    if (element) {
      const y = element?.getBoundingClientRect().top + window.scrollY - 200;
      window.scroll({
        top: y,
        behavior: 'smooth',
      });
    }
  }, 100);
};

export interface QuestionAnswer {
  answer: string[];
  question_id: number;
  theme_id: number;
}

const answers: Record<number, QuestionAnswer> = reactive({});

onMounted(async () => {
  const localTest = localStorage.getItem(`test-${currentTest.value}`);
  if (localTest) {
    store.test = JSON.parse(localTest) as IQuestion.StudentTest;
  } else {
    await store.getTest();
    localStorage.setItem(
      `test-${currentTest.value}`,
      JSON.stringify(test.value)
    );
  }
  test.value?.forEach((question) => {
    const fromLocal = JSON.parse(
      localStorage.getItem(`question-${question.questions.question_id}`) ?? '[]'
    );
    answers[question.questions.question_id] = {
      answer: fromLocal,
      question_id: question.questions.question_id,
      theme_id: question.theme_id,
    };
  });
});

const router = useRouter();

const questionsWithoutAnswer = computed(() => {
  return test.value
    ?.map((q, i) => ({ ...q, number: i + 1 }))
    ?.filter(
      (question) => !answers[question.questions.question_id].answer.length
    );
});

const isAllAnswered = computed(() => {
  return Object.values(answers).every((answer) => answer.answer.length > 0);
});

const progressValue = computed(() => {
  return isAllAnswered.value
    ? 100
    : Math.round(
        (Object.values(answers).filter((answer) => answer.answer.length > 0)
          .length /
          Object.values(answers).length) *
          100
      );
});

const isAllAnsweredModal = ref(false);

const passTest = async (withCheck = true) => {
  const answersArray = Object.values(answers);
  const data = {
    answers: { answers: answersArray },
    test_id: currentTest.value,
  };
  if (withCheck && !isAllAnswered.value) {
    isAllAnsweredModal.value = true;
  } else {
    await store.passTest(data as unknown as ITest.PassTest);
    answersArray.forEach((answer) => {
      localStorage.removeItem(`question-${answer.question_id}`);
    });
    localStorage.removeItem(`test-${currentTest.value}`);
    store.setCurrentTest(null);
    router.push({ name: 'student-profile' });
  }
};
</script>

<style lang="scss" scoped></style>
