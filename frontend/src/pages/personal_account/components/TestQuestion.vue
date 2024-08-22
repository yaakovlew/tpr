<template>
  <div class="full-width" :id="`question-${question.questions.question_id}`">
    <component
      :is="questionComponent"
      v-model="currentValue"
      :question="question"
      :questionNumber="questionNumber"
    />
  </div>
</template>

<script lang="ts" setup>
import { IQuestion } from '../../../models/test/question';
import TestQuestionFree from './TestQuestionFree.vue';
import TestQuestionSingle from './TestQuestionSingle.vue';
import TestQuestionMultiple from './TestQuestionMultiple.vue';
import { computed, ref, watch } from 'vue';

const props = defineProps<{
  question: IQuestion.SingleQuestion;
  modelValue: string[];
  questionNumber: number;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string[]): void;
}>();

const currentValue = ref(props.modelValue);

watch(currentValue, () => {
  localStorage.setItem(
    `question-${props.question.questions.question_id}`,
    JSON.stringify(currentValue.value)
  );
  emit('update:modelValue', currentValue.value);
});

const questionComponent = computed(() => {
  switch (props.question.questions.is_variable) {
    case 0:
      return TestQuestionFree;
    case 1:
      return TestQuestionSingle;
    case 2:
      return TestQuestionMultiple;
    default:
      return TestQuestionFree;
  }
});
</script>

<style lang="scss" scoped></style>
