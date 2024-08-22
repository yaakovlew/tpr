<template>
  <div
    class="discipline-name text-primary flex items-center justify-between no-wrap"
  >
    <div>
      <div>
        <span class="text-weight-bold"> Название: </span>
        {{ lesson.name }}
      </div>
      <div>
        <span class="text-weight-bold"> Описание: </span>
      </div>
      <div v-html="lesson.description" />
    </div>
    <div>
      <q-btn flat icon="download" @click="downloadDigitalLesson" />
      <q-btn flat icon="edit" @click="openEditLesson" />
      <q-btn flat icon="delete" @click="removeDisgitalLesson" />
    </div>
  </div>
  <q-dialog v-model="editLesson" persistent @hide="closeEditLesson" full-width>
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
        <q-file v-model="file" label="Выберите файл" filled />
      </q-card-section>

      <q-card-actions align="right" class="text-primary">
        <q-btn flat label="Закрыть" v-close-popup />
        <q-btn flat label="Изменить лекцию" @click="changeLesson" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script lang="ts" setup>
import { IDigitalLesson } from 'src/models/digitalLesson/digital-lesson';
import { useDigitalLessonStore } from 'src/stores/digitalLesson';
import { download } from 'src/utils/download';
import { Ref, ref } from 'vue';

const props = defineProps<{
  lesson: IDigitalLesson.DigitalLesson;
}>();

const store = useDigitalLessonStore();

const editLesson = ref(false);

const newLessonName = ref(props.lesson.name);
const newLessonDescription = ref(props.lesson.description);

const file: Ref<File | null> = ref(null);

const openEditLesson = () => {
  editLesson.value = true;
};

function saveFile(blob, filename) {
  if (window.navigator.msSaveOrOpenBlob) {
    window.navigator.msSaveOrOpenBlob(blob, filename);
  } else {
    const a = document.createElement('a');
    document.body.appendChild(a);
    const url = window.URL.createObjectURL(blob);
    a.href = url;
    a.download = filename;
    a.click();
    setTimeout(() => {
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    }, 0);
  }
}

const downloadDigitalLesson = async () => {
  const files = await store.getFilesId(props.lesson.study_guide_id);
  if (files && files.files[0]) {
    const a = await store.downloadDigitalLesson(files.files[0].file_id);
    saveFile(a, `${props.lesson.name}`);
  }
};

const closeEditLesson = () => {
  editLesson.value = false;
  newLessonName.value = props.lesson.name;
  newLessonDescription.value = props.lesson.description;
};

const changeLesson = async () => {
  await store.changeDigitalLesson({
    name: newLessonName.value,
    description: encodeURIComponent(newLessonDescription.value),
    digital_guide_id: props.lesson.study_guide_id,
    name_en: newLessonName.value,
    description_en: newLessonDescription.value,
  });

  if (file.value) {
    let formData = new FormData();
    formData.append('file', file.value);
    await store.uploadDigitalLesson({
      id: props.lesson.study_guide_id,
      file: formData,
    });
  }

  closeEditLesson();
  await store.getDigitalLessons();
};

const removeDisgitalLesson = async () => {
  await store.removeDigitalLesson(props.lesson.study_guide_id);
  await store.getDigitalLessons();
};
</script>

<style lang="scss" scoped></style>
