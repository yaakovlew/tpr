<template>
  <q-item class="text-primary flex g-m items-center justify-between">
    <div>
      <div v-if="lab.name">{{ lab.name }}</div>
      <div v-if="lab.task_description">{{ lab.task_description }}</div>
    </div>
    <q-btn
        flat
        dense
        icon="edit"
        class="cursor-pointer"
        @click="openEditLabWindow"
    />
    <q-dialog v-model="editLabDialog">
      <div class="create-test flex column justify-between g-m q-pa-md">
        <q-input v-model="editLabForm.name" label="Название" />
        <q-input
          type="textarea"
          v-model="editLabForm.description"
          label="Описание"
        />
        <q-btn
          flat
          color="primary"
          label="Применить"
          @click="editLab"
          :disable="
            !editLabForm.name
          "
        />
      </div>
    </q-dialog>
    <!-- <q-btn flat color="red" icon="delete" @click="deleteLab" /> -->
  </q-item>
</template>

<script lang="ts" setup>
import { ILabaratory } from 'src/models/labaratory/labaratory';
import { useLabsStore } from '../../../../../stores/labs';
import { reactive, ref } from 'vue';

const props = defineProps<{
  lab: ILabaratory.ExternalLabaratory;
}>();
const editLabDialog = ref(false);
const labsStore = useLabsStore();

const editLabForm = reactive({
  name: props.lab.name,
  description: props.lab.task_description,
  name_en: props.lab.name,
  description_en: props.lab.task_description,
  laboratory_id: props.lab.external_laboratory_id.toString(),
  // token: '',
  // linc: ''
});

const editLab = async () => {
  editLabDialog.value = false;
  await labsStore.editLab({
    ...editLabForm,
    description_en: editLabForm.description,
    name_en: editLabForm.name,
  });
  editLabForm.name = '';
  editLabForm.description = '';
  await labsStore.getLabs();
};

const openEditLabWindow = async () => {
  editLabDialog.value = true
  // await labsStore.editLab(editLabForm)
  // await store.deleteLab(props.lab.external_laboratory_id);
  // await store.getLabs();
};
</script>

<style lang="scss" scoped>
  .create-test {
    background-color: white;
    width: 500px;
  }
</style>
