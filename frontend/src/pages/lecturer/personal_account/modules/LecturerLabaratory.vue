<template>
  <div class="flex column g-m">
    <banner-component>
      <template #title>Лабораторные работы</template>
    </banner-component>
    <banner-component>
      <q-list separator class="q-pa-none">
        <lecturer-lab
          v-for="lab in labs"
          :key="lab.external_laboratory_id"
          :lab="lab"
        />
      </q-list>
    </banner-component>
    <banner-component
      class="text-primary profile-name"
      @click="addLabDialog = true"
    >
      <div class="flex items-center justify-between row cursor-pointer">
        Создать лабораторную работу
        <q-icon name="add" size="30px" />
      </div>
    </banner-component>
    <q-dialog v-model="addLabDialog">
      <div class="create-test flex column justify-between g-m q-pa-md">
        <q-input v-model="addLabForm.name" label="Название" />
        <q-input
          type="textarea"
          v-model="addLabForm.task_description"
          label="Описание"
        />
        <q-input
          type="number"
          v-model="addLabForm.default_mark"
          label="Максимальная оценка"
        />
        <q-input
          type="number"
          v-model="addLabForm.minutes_duration"
          label="Длительность в минутах"
        />
        <q-btn
          flat
          color="primary"
          label="Создать лабораторную работу"
          @click="addLab"
          :disable="
            !addLabForm.name &&
            !addLabForm.default_mark &&
            !addLabForm.minutes_duration
          "
        />
      </div>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import { useLabsStore } from '../../../../stores/labs';
import { onMounted, computed, ref, reactive } from 'vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import LecturerLab from './components/LecturerLab.vue';

const labsStore = useLabsStore();
const labs = computed(() => labsStore.labs);

const addLabDialog = ref(false);

const addLabForm = reactive({
  name: '',
  task_description: '',
  minutes_duration: 0,
  default_mark: 0,
});

const addLab = async () => {
  addLabDialog.value = false;
  await labsStore.addLab({
    ...addLabForm,
    minutes_duration: Number(addLabForm.minutes_duration),
    default_mark: Number(addLabForm.default_mark),
    task_description_en: addLabForm.task_description,
    name_en: addLabForm.name,
    linc: '123',
    token: '123',
    day_fine: 0,
  });
  addLabForm.name = '';
  addLabForm.task_description = '';
  addLabForm.minutes_duration = 0;
  addLabForm.default_mark = 0;
  await labsStore.getLabs();
};

onMounted(async () => {
  await labsStore.getLabs();
});
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 20px;
  font-weight: 600;
  flex-grow: 100;
}

.create-test {
  background-color: white;
  width: 500px;
}
</style>
