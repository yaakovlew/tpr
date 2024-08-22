<template>
  <div class="row justify-between">
    <q-select
      v-model="type"
      :options="typeOptions"
      label="Выберите тип"
      @update:model-value="updateModelValue"
      class="flex-grow field-element"
    />
    <q-select
      v-model="section"
      :options="sectionOptions"
      label="Выберите номер раздела"
      @update:model-value="updateModelValue"
      class="flex-grow field-element"
    />
    <q-input
      v-model="mark"
      type="number"
      label="Введите максимальную оценку"
      @update:model-value="updateModelValue"
      class="flex-grow field-element"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { ICreateDiscipne } from '../../../types/disciplines';

const props = defineProps<{
  modelValue: ICreateDiscipne;
  typeOptions: string[];
  sectionOptions: string[];
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: ICreateDiscipne): void;
}>();

const updateModelValue = () => {
  emit('update:modelValue', {
    type: type.value,
    section: section.value,
    mark: mark.value,
  });
};

const type = ref(props.modelValue.type);
const section = ref(props.modelValue.section);
const mark = ref(props.modelValue.mark);
</script>

<style lang="scss" scoped>
.field-element {
  max-width: 32%;
}
</style>
