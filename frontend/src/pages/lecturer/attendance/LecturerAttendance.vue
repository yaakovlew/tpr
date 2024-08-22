<template>
  <div class="q-ma-lg column g-m">
    <banner-component class="text-primary profile-name">
      Лекции для {{ discipline?.discipline_name }}
    </banner-component>
    <template v-if="lessons">
      <banner-component>
        <q-list class="q-pa-none">
          <template v-for="(lesson, index) in lessons" :key="lesson.id">
            <q-separator
              v-if="index !== 0"
              spaced
              inset
              class="q-px-none q-mx-none"
            />
            {{ lesson.name }}
          </template>
        </q-list>
      </banner-component>
    </template>
    <template v-else>
      <banner-component class="text-primary profile-name">
        Нет лекций
      </banner-component>
    </template>
    <banner-component
      class="text-primary profile-name"
      @click="openNewLessonDialog"
    >
      Добавить лекцию
      <q-icon name="add" color="primary" size="50px" class="cursor-pointer">
        <q-tooltip> Добавить лекцию </q-tooltip>
      </q-icon>
    </banner-component>
    <q-dialog
      v-model="newLessonDialog"
      persistent
      @hide="newLessonDialog = false"
    >
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Название лекции</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input
            dense
            v-model="newLessonName"
            autofocus
            @keyup.enter="addNewLesson"
          />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn flat label="Добавить урок" @click="addNewLesson" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import { useAttendanceStore } from 'src/stores/attendance';
import { useDisciplinesStore } from 'src/stores/disciplines';
import { computed, ref, watch } from 'vue';

const store = useAttendanceStore();
const disciplineStore = useDisciplinesStore();
const lessons = computed(() => store.lessons);
const currentDisciplineId = computed(() => store.currentDiscipline);
const discipline = computed(() => disciplineStore.discipline);

store.getLessons(currentDisciplineId.value);
disciplineStore.getDisciplineInfo(currentDisciplineId.value);

watch(currentDisciplineId, () => {
  store.getLessons(currentDisciplineId.value);
  disciplineStore.getDisciplineInfo(currentDisciplineId.value);
});

const newLessonDialog = ref(false);

const newLessonName = ref('');

const openNewLessonDialog = () => {
  newLessonDialog.value = true;
  newLessonName.value = '';
};

const addNewLesson = async () => {
  await store.createLesson(newLessonName.value);
  await store.getLessons(currentDisciplineId.value);
  newLessonDialog.value = false;
};
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}
</style>
