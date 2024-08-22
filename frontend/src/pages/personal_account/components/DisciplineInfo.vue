<template>
  <div class="q-pa-md flex column g-m">
    <banner-component>
      <template #title>
        <div class="title">{{ disciplineName }}</div>
      </template>
    </banner-component>
    <banner-component>
      <q-list separator>
        <q-item
          class="text-primary flex justify-between items-center q-px-none"
        >
          <div class="name">Название</div>
          <div class="description">Описание</div>
          <div class="name">Длительность</div>
          <div class="name">Макс. оценка</div>
        </q-item>
        <q-item
          v-for="test in disciplineTests"
          :key="test.test_id"
          class="text-primary flex justify-between items-center q-px-none"
        >
          <div class="name">{{ test.name }}</div>
          <div class="description">{{ test.task_description }}</div>

          <div class="name">{{ test.minutes_duration }}</div>
          <div class="name">{{ test.default_mark }}</div>
        </q-item>
      </q-list>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import { useTestsStore } from 'src/stores/test';
import { computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useMarksStore } from '../../../stores/mark';
import { useDisciplinesStore } from 'src/stores/disciplines';
import { useSectionStore } from 'src/stores/section';
import { useDigitalLessonStore } from 'src/stores/digitalLesson';

const testStore = useTestsStore();
const marksStore = useMarksStore();
const disciplineStore = useDisciplinesStore();
const sectionStore = useSectionStore();
const digitalLessonStore = useDigitalLessonStore();

const sections = computed(() => sectionStore.sections);
const studentMaterials = computed(() => digitalLessonStore.studentMaterials);

const route = useRoute();

const disciplineId = route.params.disciplineId;
const disciplineIdNumber = computed(() =>
  !Array.isArray(disciplineId) ? Number(disciplineId) : 0
);

const disciplines = computed(() => disciplineStore.disciplines);

const disciplineName = computed(
  () =>
    disciplines.value?.find(
      (discipline) =>
        Number(discipline.discipline_id) === disciplineIdNumber.value
    )?.name
);

const disciplineTests = computed(() => testStore.disciplineTests);

onMounted(async () => {
  await testStore.getDoneTestsStudent();
  await disciplineStore.getDisciplines();
  await marksStore.getTestMarksStudent(disciplineIdNumber.value);
  await sectionStore.getStudentDisciplineSections(disciplineIdNumber.value);
  await digitalLessonStore.getStudentMaterials(disciplineIdNumber.value);
  if (sections.value) {
    sections.value.forEach((section) => {
      testStore.getDisciplineTests(section.section_id);
    });
  }
});
</script>

<style lang="scss" scoped>
.title {
  font-size: 24px;
}
.name {
  font-size: 20px;
  width: 100px;
  font-weight: 500;
}

.description {
  font-size: 18px;
  width: 100px;
  font-weight: 500;
}
</style>
