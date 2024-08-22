<template>
  <q-item class="text-primary flex g-m items-center justify-between">
    <div>
      <div v-if="lab.name">{{ lab.name }}</div>
      <div v-if="lab.task_description">{{ lab.task_description }}</div>
    </div>
    <q-btn flat color="red" icon="delete" @click="deleteLab" />
  </q-item>
</template>

<script lang="ts" setup>
import { ILabaratory } from 'src/models/labaratory/labaratory';
import { useLabsStore } from '../../../../../stores/labs';

const props = defineProps<{
  lab: ILabaratory.ExternalLabaratory;
}>();

const store = useLabsStore();

const deleteLab = async () => {
  await store.deleteLab(props.lab.external_laboratory_id);
  await store.getLabs();
};
</script>

<style lang="scss" scoped></style>
