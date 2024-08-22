<template>
  <div class="flex g-m items-center justify-between student">
    <div>{{ student.student_name }}</div>
    <q-checkbox v-model="isPresent" />
  </div>
</template>

<script lang="ts" setup>
import { IVisiting } from 'src/models/attendance/visiting';
import { useAttendanceStore } from 'src/stores/attendance';
import { computed, ref, watch } from 'vue';

const props = defineProps<{
  student: IVisiting.StudentVisiting;
  seminarId: string;
}>();

const store = useAttendanceStore();

const isPresent = ref(!props.student.is_absent);

const isAbsent = computed(() => !isPresent.value);

watch(isAbsent, async () => {
  await store.changeSeminarVisitingSeminarian(
    isAbsent.value,
    Number(props.student.student_id),
    Number(props.seminarId)
  );
});
</script>

<style lang="scss" scoped>
.student {
  border-bottom: 1px solid var(--q-primary);
}
</style>
