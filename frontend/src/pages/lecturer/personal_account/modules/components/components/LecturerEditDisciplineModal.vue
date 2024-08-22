<template>
  <div class="flex column g-m">
    <div class="text-primary text-h5 text-bold">
      Название: {{ discipline?.discipline_name }}
    </div>
    <div>
      <div
        class="number-edit flex items-center justify-between q-py-md cursor-pointer text-primary"
      >
        <div class="flex justify-between items-center g-m">
          Количество баллов за экзамен:
          <q-input v-model="disciplineExamMark" type="number" dense />
        </div>
      </div>
      <div
        class="number-edit flex items-center justify-between q-py-md cursor-pointer text-primary"
      >
        <div class="flex justify-between items-center g-m">
          Количество баллов за семинары:
          <q-input v-model="disciplineSeminarMark" type="number" dense />
        </div>
      </div>
      <div
        class="number-edit flex items-center justify-between q-py-md cursor-pointer text-primary"
      >
        <div class="flex justify-between items-center g-m">
          Количество баллов за лекции:
          <q-input v-model="disciplineLessonMark" type="number" dense />
        </div>
      </div>
    </div>
    <div v-for="section in sections" :key="section.section_id">
      <div class="text-primary text-h6 text-bolder flex items-center">
        <div class="flex items-center g-m">
          Раздел: <q-input v-model="section.name" dense />
        </div>
        <q-btn
          label="Удалить"
          flat
          color="negative"
          @click="deleteSection(section.section_id)"
        />
      </div>
      <div>
        Тесты:
        <q-list separator>
          <q-item
            v-for="test in sectionsTests?.find(
              (t) => t.id === section.section_id
            )?.tests"
            :key="test.test_id"
            class="flex items-center justify-between"
          >
            <div class="text-primary flex column g-m q-px-none">
              <div class="dialog-title">Название: {{ test.name }}</div>
              <div class="dialog-title">
                Описание: {{ test.task_description }}
              </div>
              <div class="dialog-title">Оценка: {{ test.default_mark }}</div>
              <div class="dialog-title">
                Длительность: {{ test.minutes_duration }}
              </div>
            </div>
            <q-btn
              label="Удалить"
              flat
              color="negative"
              @click="deleteTest(test)"
            />
          </q-item>
        </q-list>
        <div class="flex items-center g-m" v-if="avaliableToAddTests?.length">
          <q-select
            v-model="testToAdd"
            :options="avaliableToAddTests"
            option-label="name"
            option-value="test_id"
            style="min-width: 300px"
          />
          <q-btn
            v-if="avaliableToAddTests?.length"
            label="Добавить тест"
            flat
            color="primary"
            @click="addTestToSection(section.section_id)"
          />
        </div>
      </div>
      <div>
        Лабораторные работы:
        <q-list separator>
          <q-item
            v-for="lab in sectionsLabs?.find(
              (l) => l?.id === section.section_id
            )?.labs"
            :key="lab.external_laboratory_id"
            class="flex items-center justify-between"
          >
            <div class="text-primary flex column g-m q-px-none">
              <div class="dialog-title">Название: {{ lab.name }}</div>
              <div class="dialog-title">
                Описание: {{ lab.task_description }}
              </div>
              <div class="dialog-title flex items-center g-m">
                Оценка:
                <q-input
                  v-model.number="lab.default_mark"
                  type="number"
                  dense
                />
              </div>
            </div>
            <q-btn
              label="Удалить"
              flat
              color="negative"
              @click="deleteLab(lab)"
            />
          </q-item>
        </q-list>
        <div class="flex items-center g-m" v-if="avaliableToAddLabs?.length">
          <q-select
            v-model="labToAdd"
            :options="avaliableToAddLabs"
            option-label="name"
            option-value="external_laboratory_id"
            style="min-width: 300px"
          />
          <q-btn
            label="Добавить лабораторную работу"
            flat
            color="primary"
            @click="addLabToSection(section.section_id)"
          />
        </div>
      </div>
    </div>
    <q-btn label="Добавить раздел" flat color="primary" @click="addSection" />
    <div class="text-primary text-h6 text-bolder">
      Сумма баллов: {{ totalSum }}
    </div>
    <q-btn label="Изменить" flat color="primary" @click="changeDiscipline" />
  </div>
</template>

<script lang="ts" setup>
import { useDisciplinesStore } from 'src/stores/disciplines';
import { useSectionStore } from 'src/stores/section';
import { IDiscipline } from 'src/models/discipline/discipline';
import { ref, onMounted, Ref, computed } from 'vue';
import { useTestsStore } from 'src/stores/test';
import { ITest } from 'src/models/test/test';
import { useLabsStore } from 'src/stores/labs';
import { ILabaratory } from 'src/models/labaratory/labaratory';
import { ISection } from '../../../../../../models/section/section';
import { Notify } from 'quasar';

const props = defineProps<{
  discipline: IDiscipline.DisciplineFullInfo;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const disciplineStore = useDisciplinesStore();
const sectionsStore = useSectionStore();
const testsStore = useTestsStore();
const labsStore = useLabsStore();

const disciplineExamMark = ref(props.discipline.exam_mark);
const disciplineSeminarMark = ref(props.discipline.seminar_visiting_mark);
const disciplineLessonMark = ref(props.discipline.lesson_visiting_mark);

const allTests = computed(() => testsStore.tests);
const allLabs = computed(() => labsStore.labs);

const avaliableToAddTests = computed(() => {
  return allTests.value?.filter((test) => {
    return sectionsTests.value?.every((section) => {
      return !section.tests?.find((t) => t.test_id === test.test_id);
    });
  });
});

const avaliableToAddLabs = computed(() => {
  return allLabs.value?.filter((lab) => {
    return sectionsLabs.value?.every((section) => {
      return !section?.labs?.find(
        (l) => l.external_laboratory_id === lab.external_laboratory_id
      );
    });
  });
});

const labToAdd = ref<ILabaratory.ExternalLabaratory>();
const testToAdd = ref<ITest.Test>();

const addLabToSection = (id: number) => {
  if (labToAdd.value) {
    let labs = sectionsLabs.value?.find((section) => section?.id === id)?.labs;
    if (labs) {
      labs?.push({
        ...labToAdd.value,
        default_mark: 1,
        laboratory_id: 0,
      });
    } else if (
      sectionsLabs.value?.find((section) => section?.id === id) &&
      sectionsLabs.value?.find((section) => section?.id === id)?.labs === null
    ) {
      sectionsLabs.value.find((section) => section?.id === id)!.labs = [
        {
          ...labToAdd.value,
          default_mark: 1,
          laboratory_id: 0,
        },
      ];
    }
    labToAdd.value = undefined;
  }
};

const addTestToSection = (id: number) => {
  if (testToAdd.value) {
    let tests = sectionsTests.value?.find(
      (section) => section?.id === id
    )?.tests;
    if (tests) {
      tests?.push(testToAdd.value);
    } else if (
      sectionsTests.value?.find((section) => section?.id === id) &&
      sectionsTests.value?.find((section) => section?.id === id)?.tests === null
    ) {
      sectionsTests.value.find((section) => section?.id === id)!.tests = [
        testToAdd.value,
      ];
    }
    testToAdd.value = undefined;
  }
};

const sections: Ref<ISection.Section[] | undefined> = ref([]);

const sectionsTests: Ref<
  { tests: ITest.Test[] | undefined; id: number }[] | undefined
> = ref(undefined);

const deleteTest = (test: ITest.Test) => {
  if (sectionsTests.value) {
    sectionsTests.value.forEach((t) => {
      t.tests = t.tests?.filter((t) => t.test_id !== test.test_id);
    });
  }
};

const deleteLab = (lab: ILabaratory.ExternalLaboratorySection) => {
  if (sectionsLabs.value) {
    sectionsLabs.value.forEach((l) => {
      if (l) {
        l.labs = l.labs?.filter(
          (l) => l.external_laboratory_id !== lab.external_laboratory_id
        );
      }
    });
  }
};

const sectionsLabs: Ref<
  | (
      | { labs: ILabaratory.ExternalLaboratorySection[]; id: number }
      | undefined
    )[]
  | undefined
> = ref(undefined);

const getSectionsTests = async () => {
  if (sections.value) {
    const promises: Promise<{ data: ITest.GetTest | undefined; id: number }>[] =
      [];
    sections.value.forEach((section) => {
      promises.push(testsStore.getSectionTests(section.section_id));
    });
    const all = await Promise.all(promises);
    return all.map((t) => {
      return {
        tests: t.data?.ru,
        id: t.id,
      };
    });
  }
};

const getSectionLabs = async () => {
  if (sections.value) {
    const promises: Promise<
      | {
          labs: ILabaratory.ExternalLaboratorySection[];
          id: number;
        }
      | undefined
    >[] = [];
    sections.value.forEach((section) => {
      promises.push(labsStore.getLabsFromSectionReturn(section.section_id));
    });
    return await Promise.all(promises);
  }
};

const newSectionsNumbers: Ref<number[]> = ref([]);

const getNewSectionNumber = (newNumber = -1): number => {
  const newSectionNumber = newNumber;
  if (newSectionsNumbers.value.find((n) => n === newSectionNumber)) {
    return getNewSectionNumber(newSectionNumber - 1);
  }
  return newSectionNumber;
};

const addSection = () => {
  const newSectionNumber = getNewSectionNumber();
  sections.value?.push({
    section_id: newSectionNumber,
    name: 'Введите название',
  });
  sectionsLabs.value?.push({
    id: newSectionNumber,
    labs: [],
  });
  sectionsTests.value?.push({
    id: newSectionNumber,
    tests: [],
  });
};

const deleteSection = (id: number) => {
  sections.value = sections.value?.filter((s) => s.section_id !== id);
  sectionsLabs.value = sectionsLabs.value?.filter((l) => l?.id !== id);
  sectionsTests.value = sectionsTests.value?.filter((t) => t?.id !== id);
};

const totalSum = computed(() => {
  let sum = 0;
  sum += Number(disciplineExamMark.value);
  sum += Number(disciplineSeminarMark.value);
  sum += Number(disciplineLessonMark.value);

  sectionsTests.value?.forEach((section) => {
    section.tests?.forEach((test) => {
      sum += Number(test.default_mark);
    });
  });

  sectionsLabs.value?.forEach((section) => {
    section?.labs?.forEach((lab) => {
      sum += Number(lab.default_mark);
    });
  });

  return sum;
});

const checkSections = async () => {
  const sectionsToCheck = await sectionsStore.getSectionsReturn(
    props.discipline.id
  );
  const sectionsChangedPromises: Promise<void>[] = [];
  sections.value?.forEach((section) => {
    const foundSection = sectionsToCheck?.find(
      (s) => s.section_id === section.section_id
    );
    if (foundSection) {
      sectionsChangedPromises.push(sectionsStore.changeSection(section));
    } else {
      sectionsChangedPromises.push(
        sectionsStore.createSection({
          name: section.name,
          name_en: section.name,
          discipline_id: props.discipline.id,
        })
      );
    }
  });
  sectionsToCheck?.forEach((section) => {
    if (!sections.value?.find((s) => s.section_id === section.section_id)) {
      sectionsChangedPromises.push(
        sectionsStore.deleteSection(section.section_id)
      );
      22;
    }
  });
  await Promise.all(sectionsChangedPromises);
};

const checkTests = async () => {
  const sections = await sectionsStore.getSectionsReturn(props.discipline.id);
  const testsToCheck = await getSectionsTests();
  const deleteAllTestsPromises: Promise<void>[] = [];
  testsToCheck?.forEach((section) => {
    section.tests?.forEach((test) => {
      deleteAllTestsPromises.push(
        sectionsStore.deleteTestFromSection({
          section_id: section.id,
          test_id: test?.test_id,
        })
      );
    });
  });
  await Promise.all(deleteAllTestsPromises);
  const addAllTestsPromises: Promise<void>[] = [];
  sectionsTests.value?.forEach((section) => {
    section.tests?.forEach((test) => {
      let sectionId: number | undefined = section.id;
      if (sectionId < 0) {
        const name = sections?.find((s) => s.section_id === sectionId)?.name;
        sectionId = sections?.find((s) => s.name === name)?.section_id;
      }
      if (!sectionId) {
        Notify.create({
          message: 'Произошла ошибка при добавлении теста, попробуйте еще раз',
          color: 'warning',
        });
        return;
      }
      addAllTestsPromises.push(
        sectionsStore.addTestToSection({
          section_id: sectionId,
          test_id: test?.test_id,
        })
      );
    });
  });
  await Promise.all(addAllTestsPromises);
};

const checkLabs = async () => {
  const sections = await sectionsStore.getSectionsReturn(props.discipline.id);
  const labsToCheck = await getSectionLabs();
  const deleteAllLabsPromises: Promise<void>[] = [];
  labsToCheck?.forEach((section) => {
    section?.labs?.forEach((lab) => {
      deleteAllLabsPromises.push(
        sectionsStore.deleteLabFromSection({
          section_id: section.id,
          laboratory_id: lab?.laboratory_id,
        })
      );
    });
  });
  await Promise.all(deleteAllLabsPromises);
  const addAllLabsPromises: Promise<void>[] = [];
  sectionsLabs.value?.forEach((section) => {
    section?.labs?.forEach((lab) => {
      let sectionId: number | undefined = section.id;
      if (sectionId < 0) {
        const name = sections?.find((s) => s.section_id === sectionId)?.name;
        sectionId = sections?.find((s) => s.name === name)?.section_id;
      }
      if (!sectionId) {
        Notify.create({
          message:
            'Произошла ошибка при добавлении лаборатории, попробуйте еще раз',
          color: 'warning',
        });
        return;
      }
      addAllLabsPromises.push(
        sectionsStore.addLabToSection({
          section_id: sectionId,
          external_lab_id: lab?.external_laboratory_id,
          default_mark: lab?.default_mark,
        })
      );
    });
  });
  await Promise.all(addAllLabsPromises);
};

const changeDiscipline = async () => {
  let message: string | null = null;
  if (sections.value?.length === 0) {
    message = 'Количество разделов должно быть больше 0';
  } else if (
    Number(disciplineExamMark.value) < 0 ||
    Number(disciplineSeminarMark.value) < 0 ||
    Number(disciplineLessonMark.value) < 0
  ) {
    message =
      'Колличество баллов за посещения и за экзамен не должно быть отрицательным числом';
  }
  if (totalSum.value !== 100) {
    message = 'Необходимо чтобы сумма баллов была равна 100';
  }

  if (message !== null) {
    Notify.create({
      message: message,
      color: 'warning',
    });
    return;
  }

  await disciplineStore.changeDisciplineExamMark({
    discipline_id: Number(props.discipline.id),
    exam_mark: Number(disciplineExamMark.value),
  });
  await disciplineStore.changeDisciplineSeminarMark({
    discipline_id: Number(props.discipline.id),
    seminar_mark: Number(disciplineSeminarMark.value),
  });
  await disciplineStore.changeDisciplineLessonMark({
    discipline_id: Number(props.discipline.id),
    lesson_mark: Number(disciplineLessonMark.value),
  });

  await checkSections();
  await checkTests();
  await checkLabs();
  emit('close');
};

onMounted(async () => {
  sections.value = await sectionsStore.getSectionsReturn(props.discipline.id);
  sectionsTests.value = await getSectionsTests();
  40;
  sectionsLabs.value = await getSectionLabs();
  await testsStore.getTests();
  await labsStore.getLabs();
});
</script>

<style lang="scss" scoped>
.number-edit {
  width: 100%;
  border-bottom: 1px solid var(--q-primary);
  font-size: 18px;
}
</style>
