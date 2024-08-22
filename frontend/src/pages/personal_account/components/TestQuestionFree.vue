<template>
  <div class="flex column g-m">
    <div class="primary-text">
      {{ questionNumber + '. ' + question.questions.question }}
    </div>
    <q-input square filled v-model="currentValue" label="Введите ответ" />
  </div>
</template>

<script lang="ts" setup>
import { IQuestion } from 'src/models/test/question';
import { ref, watch } from 'vue';

const props = defineProps<{
  question: IQuestion.SingleQuestion;
  modelValue: string[];
  questionNumber: number;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string[]): void;
}>();

const currentValue = ref(
  props.modelValue.length ? props.modelValue[0].trim() : ''
);

watch(currentValue, () => {
  emit('update:modelValue', [currentValue.value.trim()]);
});
</script>

<style lang="scss" scoped></style>
