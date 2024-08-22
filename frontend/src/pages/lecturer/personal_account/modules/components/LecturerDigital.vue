<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Лекции
    </banner-component>
    <template v-if="lessons">
      <banner-component>
        <q-list class="q-pa-none">
          <template v-for="(lesson, index) in lessons" :key="lesson.name">
            <q-separator
              v-if="index !== 0"
              spaced
              inset
              class="q-px-none q-mx-none"
            />
            <lecturer-single-digital :lesson="lesson" />
          </template>
        </q-list>
      </banner-component>
    </template>
    <template v-else>
      <banner-component class="text-primary profile-name">
        Нет лекций
      </banner-component>
    </template>
    <banner-component class="text-primary profile-name" @click="openNewLesson">
      Добавить лекцию
      <q-icon name="add" color="primary" size="50px" class="cursor-pointer">
        <q-tooltip> Добавить лекцию </q-tooltip>
      </q-icon>
    </banner-component>
    <q-dialog v-model="newLesson" persistent @hide="closeNewLesson" full-width>
      <q-card style="width: 80vw">
        <q-card-section>
          <div class="text-h6">Название лекции</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input dense v-model="newLessonName" autofocus />
        </q-card-section>

        <q-card-section>
          <div class="text-h6">Описание</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-editor
            v-model="newLessonDescription"
            :dense="$q.screen.lt.md"
            :toolbar="[
              [
                {
                  label: 'Выравнивание',
                  icon: $q.iconSet.editor.align,
                  fixedLabel: true,
                  list: 'only-icons',
                  options: ['left', 'center', 'right', 'justify'],
                },
              ],
              [
                'bold',
                'italic',
                'strike',
                'underline',
                'subscript',
                'superscript',
              ],
              ['token', 'hr', 'link', 'custom_btn'],
              ['print', 'fullscreen'],
              [
                {
                  label: 'Формат',
                  icon: $q.iconSet.editor.formatting,
                  list: 'no-icons',
                  options: ['p', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'code'],
                },
                {
                  label: 'Размер',
                  icon: $q.iconSet.editor.fontSize,
                  fixedLabel: true,
                  fixedIcon: true,
                  list: 'no-icons',
                  options: [
                    'size-1',
                    'size-2',
                    'size-3',
                    'size-4',
                    'size-5',
                    'size-6',
                    'size-7',
                  ],
                },
                {
                  label: 'Шрифт',
                  icon: $q.iconSet.editor.font,
                  fixedIcon: true,
                  list: 'no-icons',
                  options: [
                    'default_font',
                    'arial',
                    'arial_black',
                    'comic_sans',
                    'courier_new',
                    'impact',
                    'lucida_grande',
                    'times_new_roman',
                    'verdana',
                  ],
                },
                'removeFormat',
              ],
              ['quote', 'unordered', 'ordered', 'outdent', 'indent'],

              ['undo', 'redo'],
              ['viewsource'],
            ]"
            :fonts="{
              arial: 'Arial',
              arial_black: 'Arial Black',
              comic_sans: 'Comic Sans MS',
              courier_new: 'Courier New',
              impact: 'Impact',
              lucida_grande: 'Lucida Grande',
              times_new_roman: 'Times New Roman',
              verdana: 'Verdana',
            }"
          />
          <!-- <q-input v-model="newLessonDescription" /> -->
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn flat label="Добавить лекцию" @click="addNewLesson" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import { useDigitalLessonStore } from 'src/stores/digitalLesson';
import { computed, onMounted, ref } from 'vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import LecturerSingleDigital from './LecturerSingleDigital.vue';

const store = useDigitalLessonStore();

const lessons = computed(() => store.digitalLessons);

const newLesson = ref(false);

const newLessonName = ref('');
const newLessonDescription = ref('');

const openNewLesson = () => {
  newLesson.value = true;
};

const closeNewLesson = () => {
  newLesson.value = false;
  newLessonName.value = '';
  newLessonDescription.value = '';
};

const addNewLesson = async () => {
  await store.addDigitalLesson({
    name: newLessonName.value,
    name_en: newLessonName.value,
    description: newLessonDescription.value,
    description_en: newLessonDescription.value,
  });
  closeNewLesson();
  await store.getDigitalLessons();
};

onMounted(async () => {
  await store.getDigitalLessons();
});
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}
</style>
