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
  lessonId: string;
}>();

const store = useAttendanceStore();

const isPresent = ref(!props.student.is_absent);

const isAbsent = computed(() => !isPresent.value);

watch(isAbsent, async () => {
  await store.changeLessonVisiting(
    isAbsent.value,
    Number(props.lessonId),
    Number(props.student.student_id)
  );
});

watch(
  () => props.student,
  async () => {
    isPresent.value = !props.student.is_absent;
  }
);
</script>

<style lang="scss" scoped>
.student {
  border-bottom: 1px solid var(--q-primary);
}
</style>
