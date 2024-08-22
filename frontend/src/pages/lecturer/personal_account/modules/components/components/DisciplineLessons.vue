<template>
  <banner-component>
    <div>Лекции</div>
    <q-list separator>
      <q-item
        v-for="lesson in fullDisciplineLessons"
        :key="lesson.study_guide_id"
        class="q-px-none flex items-center g-m full-width justify-between"
      >
        <div>
          {{ lesson.name }}
        </div>
        <div class="flex g-m">
          <q-icon
            name="delete"
            class="cursor-pointer"
            @click="deleteLesson(lesson.study_guide_id)"
          />
        </div>
      </q-item>
    </q-list>
    <div
      class="cursor-pointer flex g-m items-center"
      @click="openAddLessonModal"
    >
      <div>Добавить лекцию</div>
      <q-icon name="add" size="lg" />
    </div>
  </banner-component>
  <q-dialog v-model="addLessonModal" @hide="closeAddLessonModal">
    <div class="modal g-m">
      <q-select
        v-model="lessonName"
        :options="digitaLessons"
        option-label="name"
        option-value="study_guide_id"
      />
      <q-btn label="Добавить" :disable="!lessonName" @click="addLesson" />
    </div>
  </q-dialog>
  <q-dialog v-model="changeLessonModal" @hide="closeChangeLessonModal">
    <div class="modal g-m">
      <q-input label="Введите название лекции" v-model="lessonNameEdit" />
      <q-btn
        label="Изменить"
        :disable="!lessonNameEdit"
        @click="changeLessonName"
      />
    </div>
  </q-dialog>
</template>

<script setup lang="ts">
import BannerComponent from 'src/components/BannerComponent.vue';
import { IDiscipline } from '../../../../../../models/discipline/discipline';
import { useAttendanceStore } from '../../../../../../stores/attendance';
import { computed, ref, watch, Ref, onMounted } from 'vue';
import { ILesson } from '../../../../../../models/attendance/attendance';
import { useDigitalLessonStore } from 'src/stores/digitalLesson';
import { IDigitalLesson } from 'src/models/digitalLesson/digital-lesson';

const props = defineProps<{
  discipline: IDiscipline.DisciplineFullInfo;
}>();

const store = useAttendanceStore();
const digitalStore = useDigitalLessonStore();

onMounted(async () => {
  await digitalStore.getDigitalLessons();
  if (props.discipline) {
    await digitalStore.getDisciplineLessons(props.discipline.id);
  }
});

const digitaLessons = computed(() => digitalStore.digitalLessons);
const fullDisciplineLessons = computed(
  () => digitalStore.fullDisciplineLessons
);

const lessons = computed(() => store.lessons);

watch(
  () => props.discipline,
  async () => {
    await digitalStore.getDigitalLessons();
    await digitalStore.getDisciplineLessons(props.discipline.id);
    store.setCurrentDiscipline(String(props.discipline.id));
  }
);

const deleteLesson = async (lessonId: number) => {
  await digitalStore.removeDigitalLessonFromDiscipline({
    discipline_id: props.discipline.id,
    digital_material_id: Number(lessonId),
  });
  await digitalStore.getDisciplineLessons(props.discipline.id);
};

const addLessonModal = ref(false);
const changeLessonModal = ref(false);

const openAddLessonModal = () => {
  addLessonModal.value = true;
};

const closeAddLessonModal = () => {
  // lessonName.value = '';
  addLessonModal.value = false;
};

const closeChangeLessonModal = () => {
  lessonNameEdit.value = '';
  selectedLesson.value = '';
  changeLessonModal.value = false;
};

const lessonName: Ref<IDigitalLesson.DigitalLesson | undefined> = ref();
const lessonNameEdit = ref('');

const selectedLesson = ref('');

const changeLessonName = async () => {
  await store.changeLessonName(selectedLesson.value, lessonNameEdit.value);
  await store.getLessons(String(props.discipline.id));
  closeChangeLessonModal();
};

const openChangeLesson = (lesson: ILesson) => {
  selectedLesson.value = lesson.lesson_id;
  lessonNameEdit.value = lesson.name;
  changeLessonModal.value = true;
};

const addLesson = async () => {
  if (lessonName.value) {
    await store.createLesson(lessonName.value.name);
    await digitalStore.addDigitalLessonToDiscipline({
      digital_material_id: lessonName.value.study_guide_id,
      discipline_id: props.discipline.id,
    });
  }
  await digitalStore.getDisciplineLessons(props.discipline.id);
  await store.getLessons(String(props.discipline.id));
  closeAddLessonModal();
};
</script>

<style scoped lang="scss">
.modal {
  background-color: white;
  min-width: 300px;
  padding: 20px;
  display: flex;
  flex-direction: column;
}
</style>
