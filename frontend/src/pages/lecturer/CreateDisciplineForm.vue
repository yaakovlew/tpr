<template>
  <div class="q-pa-lg column g-m">
    <banner-component class="page-title text-center text-primary">
      Создание дисциплины
    </banner-component>
    <div class="column g-m">
      <banner-component class="flex column g-m">
        <div class="row g-m">
          <q-input
            v-model="type"
            label="Введите название дисциплины"
            class="flex-grow"
          />
          <q-input
            v-model="sectionsCount"
            type="number"
            label="Введите количество разделов"
            class="flex-grow"
          />
        </div>
        <q-input
          v-model="description"
          label="Введите описание дисциплины"
          class="flex-grow"
          type="textarea"
        />
        <q-input
          v-model="visitingMark"
          label="Введите количество баллов за посещение"
          class="flex-grow"
          type="number"
        />
      </banner-component>
      <banner-component class="column g-m">
        <add-test-field
          v-for="(text, index) in testsCount"
          :key="index"
          :typeOptions="typeOptions"
          :sectionOptions="sectionOptions"
          v-model="addTestFieldValue[index]"
        />
        <q-icon
          name="add"
          color="primary"
          size="50px"
          class="cursor-pointer"
          @click="addTest"
        >
          <q-tooltip> Добавить тест </q-tooltip>
        </q-icon>
      </banner-component>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, WritableComputedRef } from 'vue';
import BannerComponent from '../../components/BannerComponent.vue';
import AddTestField from './components/AddTestField.vue';
import { ICreateDiscipne } from '../../types/disciplines';

const type = ref('');
const sectionsCount = ref(1);
const description = ref('');
const visitingMark = ref('');

const addTestFieldValue: WritableComputedRef<ICreateDiscipne[]> = computed({
  get() {
    const result: ICreateDiscipne[] = [
      {
        type: '',
        section: 1,
        mark: 0,
      },
    ];

    for (let i = 1; i < testsCount.value; i++) {
      result.push({
        type: '',
        section: 1,
        mark: 0,
      });
    }
    return result;
  },
  set(newValue) {
    console.log(newValue);
  },
});

const typeOptions = ['Тест', 'Лабараторная работа'];
const sectionOptions = computed(() => {
  const conut = ['1'];
  for (let i = 1; i < sectionsCount.value; i++) {
    conut.push(String(i + 1));
  }
  return conut;
});

const testsCount = ref(1);

const addTest = () => {
  testsCount.value++;
};
</script>

<style lang="scss" scoped>
.add-icon {
  //   border: 1px solid rgba(201, 201, 201, 0.715);
  //   width: fit-content;
  //   border-radius: 100%;
}
</style>
