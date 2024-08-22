<template>
  <q-item class="q-px-small flex column g-m">
    <banner-with-back> Новая дисциплина </banner-with-back>

    <banner-component>
      <q-input v-model="disciplineName" label="Название дисциплины" />
    </banner-component>

    <banner-component>
      <div class="flex justify-between">
        <q-input
          v-model="sectionCounter"
          type="number"
          label="Количество разделов"
          class="input-numbers"
          outlined
          :error="enterDiscipline"
        />
        <q-input
          v-model="examMark"
          type="number"
          label="Количество баллов за экзамен"
          class="input-numbers"
          outlined
        />
        <q-input
          v-model="attendanceSeminarMark"
          type="number"
          label="Количество баллов за посещение семинаров"
          class="input-numbers"
          outlined
        />
        <q-input
          v-model="attendanceLessonMark"
          type="number"
          label="Количество баллов за посещение лекций"
          class="input-numbers"
          outlined
        />
      </div>
    </banner-component>
    <template v-if="sections.length">
      <banner-component
        class="text-primary title"
        v-for="(section, index) in sections"
        :key="index"
      >
        <div class="flex column g-m">
          {{ section.name }}
          <q-input v-model="section.name" label="Название раздела" />
          <template v-if="section.tests">
            <q-list separator>
              <q-item
                v-for="(test, index) in section.tests"
                :key="index"
                class="q-px-none flex column g-m"
              >
                <div class="test-info flex justify-between">
                  <div>
                    <div>Название: {{ test.name }}</div>
                    <div>Описание: {{ test.task_description }}</div>
                    <div>Количество баллов: {{ test.default_mark }}</div>
                  </div>
                  <q-btn
                    flat
                    label="Удалить"
                    @click="deleteTest(section, test)"
                  />
                </div>
                <q-select
                  v-model="section.tests[index]"
                  :options="avalibaleToAddTests"
                  option-label="name"
                  option-value="test_id"
                />
              </q-item>
            </q-list>
          </template>
          <div v-else>Тестов нет</div>
          <q-btn
            v-if="avalibaleToAddTests.length"
            flat
            color="primary"
            label="Добавить тест"
            @click="addTest(section)"
          />
          <template v-if="section.labs">
            <q-list separator>
              <q-item
                v-for="(lab, index) in section.labs"
                :key="index"
                class="q-px-none flex column g-m"
              >
                <div class="test-info flex justify-between">
                  <div>
                    <div>Название: {{ lab.name }}</div>
                    <div>Описание: {{ lab.task_description }}</div>
                    <div>
                      Количество баллов:
                      {{ section.labs_marks[lab.external_laboratory_id] }}
                    </div>
                  </div>
                  <q-btn
                    flat
                    label="Удалить"
                    @click="deleteLab(section, lab)"
                  />
                </div>
                <div>
                  <q-input
                    v-model.number="
                      section.labs_marks[lab.external_laboratory_id]
                    "
                    type="number"
                    label="Количество баллов"
                  />
                  <q-select
                    v-model="section.labs[index]"
                    :options="avaliableToAddLabs"
                    option-label="name"
                    option-value="external_labaratory_id"
                  />
                </div>
              </q-item>
            </q-list>
          </template>
          <div v-else>Лабораторных работ нет</div>
          <q-btn
            v-if="avalibaleToAddTests.length"
            flat
            color="primary"
            label="Добавить лабораторную работу"
            @click="addLab(section)"
          />
        </div>
      </banner-component>
    </template>
    <banner-component class="text-primary title" v-else>
      Разделов нет
    </banner-component>
    <banner-component
      class="text-primary title cursor-pointer"
      @click="addDiscipline"
    >
      <div class="flex items-center">Итоговая сумма баллов: {{ totalSum }}</div>
      <div>
        Добавить дисциплину
        <q-icon name="add" color="primary" size="30px" class="cursor-pointer">
          <q-tooltip> Добавить дисциплину </q-tooltip>
        </q-icon>
      </div>
    </banner-component>
  </q-item>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import { useDisciplinesStore } from 'src/stores/disciplines';
import { useRouter } from 'vue-router';
import { Ref, computed, reactive, ref, watch } from 'vue';
import BannerWithBack from 'src/components/BannerWithBack.vue';
import { Notify } from 'quasar';
import { useSectionStore } from '../../../../../stores/section';
import { useTestsStore } from '../../../../../stores/test';
import { ITest } from 'src/models/test/test';
import { useLabsStore } from '../../../../../stores/labs';
import { ILabaratory } from 'src/models/labaratory/labaratory';

const store = useDisciplinesStore();
const sectionStore = useSectionStore();
const testStore = useTestsStore();
const labStore = useLabsStore();

const router = useRouter();

interface SectionToAdd {
  name: string;
  tests: ITest.Test[];
  labs: ILabaratory.ExternalLabaratory[];
  labs_marks: Record<number, number>;
}

const sections: Ref<SectionToAdd[]> = ref([]);

const disciplineName = ref('');
const sectionCounter: Ref<number> = ref(0);
const examMark: Ref<number> = ref(0);
const attendanceSeminarMark: Ref<number> = ref(0);
const attendanceLessonMark: Ref<number> = ref(0);

watch(sectionCounter, (newValue, oldValue) => {
  if (newValue < 0) {
    return [];
  }
  if (newValue > oldValue) {
    for (let i = oldValue; i < newValue; i++) {
      sections.value.push({
        name: `Раздел ${i + 1}`,
        tests: [],
        labs: [],
        labs_marks: {},
      });
    }
  } else if (newValue < oldValue) {
    sections.value.length = newValue;
  }
});

watch(sectionCounter, () => {
  if (sectionCounter.value !== 0) {
    enterDiscipline.value = false;
  }
});

const enterDiscipline = ref(false);

const tests: ITest.Test[] = reactive([]);

const allTests = computed(() => testStore.tests);
const allLabs = computed(() => labStore.labs);

const avalibaleToAddTests = computed(() =>
  allTests.value.filter((test) => {
    for (const section of sections.value) {
      if (section.tests.find((t) => t.test_id === test.test_id)) {
        return false;
      }
    }
    return true;
  })
);

const avaliableToAddLabs = computed(() =>
  allLabs.value?.filter((lab) => {
    for (const section of sections.value) {
      if (
        section.labs.find(
          (t) => t.external_laboratory_id === lab.external_laboratory_id
        )
      ) {
        return false;
      }
    }
    return true;
  })
);

testStore.getTests();
labStore.getLabs();

const addTest = (section: SectionToAdd) => {
  if (avalibaleToAddTests.value.length) {
    section.tests.push(avalibaleToAddTests.value[0]);
  }
};

const labMark: Ref<number> = ref(0);

const addLab = (section: SectionToAdd) => {
  if (avaliableToAddLabs.value?.length) {
    section.labs.push(avaliableToAddLabs.value[0]);
    section.labs_marks[avaliableToAddLabs.value[0].external_laboratory_id] = 0;
    labMark.value = 0;
  }
};

const deleteTest = (section: SectionToAdd, test: ITest.Test) => {
  section.tests = section.tests.filter((t) => t.test_id !== test.test_id);
};

const deleteLab = (
  section: SectionToAdd,
  lab: ILabaratory.ExternalLabaratory
) => {
  section.labs = section.labs.filter(
    (t) => t.external_laboratory_id !== lab.external_laboratory_id
  );
};

const totalSum = computed(() => {
  let sum = 0;
  sum += Number(examMark.value);
  sum += Number(attendanceSeminarMark.value);
  sum += Number(attendanceLessonMark.value);

  sections.value.forEach((section) => {
    section.tests.forEach((test) => {
      sum += Number(test.default_mark);
    });
  });

  sections.value.forEach((section) => {
    Object.values(section.labs_marks).forEach((mark) => {
      sum += Number(mark);
    });
  });

  return sum;
});

const disciplines = computed(() => store.disciplines);
const disciplineSections = computed(() => sectionStore.sections);

const addDiscipline = async () => {
  let message: string | null = null;

  if (!disciplineName.value) {
    message = 'Введите название дисциплины';
  } else if (sectionCounter.value <= 0) {
    message = 'Количество разделов должно быть больше 0';
  } else if (
    Number(examMark.value) < 0 ||
    Number(attendanceSeminarMark.value) < 0 ||
    Number(attendanceLessonMark.value) < 0
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
  } else {
    await store.createDiscipline({
      discipline_name: disciplineName.value,
      discipline_name_en: disciplineName.value,
      seminar_visiting_mark: Number(attendanceSeminarMark.value),
      lesson_visiting_mark: Number(attendanceLessonMark.value),
      exam_mark: Number(examMark.value),
    });
    await store.getDisciplines();
    const createdDiscipline = disciplines.value?.find(
      (d) => d.name === disciplineName.value
    );
    const sectionsPromises: Promise<void>[] = [];
    if (createdDiscipline) {
      sections.value.forEach((section) => {
        sectionsPromises.push(
          sectionStore.createSection({
            name: section.name,
            discipline_id: Number(createdDiscipline.discipline_id),
            name_en: section.name,
          })
        );
      });
      await Promise.all(sectionsPromises);
      await sectionStore.getSections(Number(createdDiscipline.discipline_id));
      const testsPromises: Promise<void>[] = [];
      if (disciplineSections.value) {
        sections.value.forEach((section) => {
          const foundSection = disciplineSections.value?.find(
            (s) => s.name === section.name
          );
          if (foundSection) {
            section.tests.forEach((test) => {
              testsPromises.push(
                sectionStore.addTestToSection({
                  section_id: foundSection.section_id,
                  test_id: test.test_id,
                })
              );
            });
            section.labs.forEach((lab) => {
              testsPromises.push(
                sectionStore.addLabToSection({
                  section_id: foundSection.section_id,
                  external_lab_id: lab.external_laboratory_id,
                  default_mark: Number(
                    section.labs_marks[lab.external_laboratory_id]
                  ),
                })
              );
            });
          }
        });
      }
      await Promise.all(testsPromises);
    }
    router.back();
  }
};
</script>

<style lang="scss" scoped>
.title {
  font-weight: 600;
  font-size: 18px;
}

.input-numbers {
  width: 45%;
}

.test-info {
  font-size: 16px;
}
</style>
