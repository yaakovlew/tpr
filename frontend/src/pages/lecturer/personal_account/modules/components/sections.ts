import { sectionCounter } from './LecturerDiscipline.vue';

const sections = computed(() => {
  const array = [];
  if (sectionCounter.value) {
    for (let num = 1; num <= sectionCounter.value; num += 1)
      array.push(num + 1);
  }
});
