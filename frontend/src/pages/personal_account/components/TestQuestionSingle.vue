<template>
  <div>
    <div class="primary-text">
      {{ questionNumber + '. ' + question.questions.question }}
    </div>
    <q-option-group v-model="currentValue" :options="options" />
  </div>
</template>

<script lang="ts" setup>
import { IQuestion } from 'src/models/test/question';
import { ref, watch, computed } from 'vue';

const props = defineProps<{
  question: IQuestion.SingleQuestion;
  modelValue: string[];
  questionNumber: number;
}>();

const options = computed(() => {
  return props.question.answers.map((answer) => {
    return {
      label: answer.name,
      value: answer.name,
    };
  });
});

const emit = defineEmits<{
  (e: 'update:modelValue', value: string[]): void;
}>();

const currentValue = ref(props.modelValue.length ? props.modelValue[0] : '');

watch(currentValue, () => {
  emit('update:modelValue', [currentValue.value]);
});
</script>

<style lang="scss" scoped></style>
