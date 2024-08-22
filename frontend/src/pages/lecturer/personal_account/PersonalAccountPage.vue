<template>
  <div class="flex q-pa-lg no-wrap page-wrap">
    <q-tabs
      v-model="tab"
      active-bg-color="grey-4"
      :active-class="'mdi-border-radius'"
      class="text-grey-9"
      :vertical="width > 500"
    >
      <q-tab name="profile" label="Профиль" />
      <q-tab name="disciplins" label="Дисциплины" />
      <q-tab name="groups" label="Группы" />
      <q-tab name="students" label="Студенты" />
      <q-tab name="seminarians" label="Семинаристы" />
      <q-tab name="tests" label="Тесты" />
      <q-tab name="lessons" label="Лекции" />
      <q-tab name="labs" label="Лабораторные работы" />
    </q-tabs>
    <q-tab-panels v-model="tab" class="bg-none q-pa-none flex-grow">
      <q-tab-panel name="profile" class="q-py-none">
        <personal-info />
      </q-tab-panel>
      <q-tab-panel name="disciplins" class="q-py-none">
        <personal-disciplines />
      </q-tab-panel>
      <q-tab-panel name="groups" class="q-py-none">
        <lecturer-groups />
      </q-tab-panel>
      <q-tab-panel name="students" class="q-py-none">
        <lecturer-students />
      </q-tab-panel>
      <q-tab-panel name="seminarians" class="q-py-none">
        <lecturer-seminarians />
      </q-tab-panel>
      <q-tab-panel name="tests" class="q-py-none">
        <lecturer-tests />
      </q-tab-panel>
      <q-tab-panel name="lessons" class="q-py-none">
        <lecturer-digital />
      </q-tab-panel>
      <q-tab-panel name="labs" class="q-py-none">
        <lecturer-labaratory />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, watch } from 'vue';
import PersonalInfo from './modules/PersonalInfo.vue';
import PersonalDisciplines from './modules/PersonalDisciplines.vue';
import LecturerGroups from './modules/LecturerGroups.vue';
import LecturerStudents from './modules/LecturerStudents.vue';
import LecturerSeminarians from './modules/LecturerSeminarians.vue';
import LecturerTests from './modules/LecturerTests.vue';
import LecturerDigital from './modules/components/LecturerDigital.vue';
import LecturerLabaratory from './modules/LecturerLabaratory.vue';

const tab = ref(localStorage.getItem('personal-tab') || 'profile');

watch(tab, () => {
  localStorage.setItem('personal-tab', tab.value);
});

const width = computed(() => window.innerWidth);
</script>

<style lang="scss" scoped>
@media screen and (max-width: 500px) {
  .page-wrap {
    flex-direction: column;
  }
}
</style>
