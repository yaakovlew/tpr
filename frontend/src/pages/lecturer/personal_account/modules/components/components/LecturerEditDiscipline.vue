<template>
  <div class="q-px-md flex column g-m text-primary discipline-edit">
    <banner-with-back>
      <div class="cursor-pointer flex items-center justify-between">
        {{ discipline?.discipline_name }}
        <q-btn
          flat
          icon="edit"
          label="Редактировать баллы"
          @click="dicsiplineModal = true"
        />
      </div>
    </banner-with-back>
    <banner-component class="cursor-pointer flex items-center justify-between">
      <div class="flex row items-center justify-between full-width">
        Название: {{ discipline?.discipline_name }}
        <q-icon name="edit" />
        <q-popup-edit
          v-model="disciplineName"
          title="Измените название дисциплины"
          label-set="Сохранить"
          label-cancel="Отменить"
          v-slot="scope"
        >
          <q-input
            v-model="scope.value"
            dense
            autofocus
            counter
            @keyup.enter="changeDisciplineName(scope.value)"
            @keyup.enter.prevent="scope.set"
          >
            <template v-slot:after>
              <q-btn
                flat
                dense
                color="negative"
                icon="cancel"
                @click.stop.prevent="scope.cancel"
              />

              <q-btn
                flat
                dense
                color="positive"
                icon="check_circle"
                @click.stop.prevent="changeDisciplineName(scope.value)"
                @click="scope.set"
                :disable="
                  scope.validate(scope.value) === false ||
                  scope.initialValue === scope.value
                "
              />
            </template>
          </q-input>
        </q-popup-edit>
      </div>
    </banner-component>

    <banner-component class="flex justify-between g-m">
      <div
        class="number-edit flex items-center justify-between q-py-md cursor-pointer"
      >
        <div>Количество баллов за экзамен: {{ disciplineExamMark }}</div>
      </div>
      <div
        class="number-edit flex items-center justify-between q-py-md cursor-pointer"
      >
        <div>Количество баллов за семинары: {{ disciplineSeminarMark }}</div>
      </div>
      <div
        class="number-edit flex items-center justify-between q-py-md cursor-pointer"
      >
        <div>Количество баллов за лекции: {{ disciplineLessonMark }}</div>
      </div>
      <q-button label="Изменить" />
    </banner-component>
    <banner-component>
      <div>Разделы</div>
      <q-list separator>
        <q-item
          v-for="section in sections"
          :key="section.section_id"
          class="flex items-center justify-between"
        >
          <div>
            {{ section.name }}
          </div>
          <q-icon
            name="list"
            @click="getSectionTests(section.section_id)"
            size="24px"
            class="cursor-pointer"
          />
        </q-item>
      </q-list>
    </banner-component>
    <banner-component>
      <div>Группы</div>
      <q-list separator>
        <q-item
          v-for="group in groups"
          :key="group.group_id"
          class="q-px-none flex items-center g-m full-width justify-between"
        >
          <div>
            {{ group.name }}
          </div>
          <div class="flex g-m">
            <q-icon
              name="person"
              class="cursor-pointer"
              @click="getGroupSeminarians(group.group_id)"
            />
            <q-icon
              name="delete"
              class="cursor-pointer"
              @click="deleteGroupFromDiscipline(group.group_id)"
            />
          </div>
        </q-item>
      </q-list>
      <div
        class="cursor-pointer flex g-m items-center"
        @click="openAddGroupModal"
      >
        <div>Добавить группу</div>
        <q-icon name="add" size="lg" />
      </div>
    </banner-component>
    <discipline-lessons :discipline="discipline" />
  </div>
  <q-dialog v-model="addGroupModal">
    <div class="add-group-modal flex column g-m">
      <q-select
        v-model="groupToAdd"
        :options="avaliableGroups"
        option-value="group_id"
        option-label="name"
        label="Группа"
      />
      <q-btn
        label="Добавить"
        :disable="!groupToAdd"
        @click="addGroupToDiscipline"
      />
    </div>
  </q-dialog>
  <q-dialog
    v-model="groupSeminariansModal"
    @hide="addSeminariansOpenBoolean = false"
  >
    <q-card style="min-width: 350px">
      <template v-if="seminarians">
        <!-- <q-card-section>
            <div class="dialog-title">Ceминаристы группы {{ group.name }}</div>
          </q-card-section> -->
        <q-card-section class="q-pt-none">
          <q-list separator>
            <q-item
              v-for="seminarian in seminarians"
              :key="seminarian.seminarian_id"
              v-ripple
              class="q-px-none"
            >
              <q-item-section>
                {{ `${seminarian.name} ${seminarian.surname}` }}
              </q-item-section>
            </q-item>
          </q-list>
        </q-card-section>
      </template>
      <q-card-section v-else>
        <div class="dialog-title">У группы нет семинаристов</div>
      </q-card-section>
      <q-card-section>
        <div
          class="dialog-title"
          v-if="!addSeminariansOpenBoolean"
          @click="addSeminariansOpen"
        >
          Добавить семинриста
        </div>
        <div v-else class="flex justify-end g-m">
          <q-select
            v-model="selectedSeminarian"
            :options="avaliableSeminarians"
            option-value="seminarian_id"
            option-label="name"
            class="full-width"
          />
          <q-btn
            label="Добавить"
            @click="addSeminarian"
            :disable="!selectedSeminarian"
          />
        </div>
      </q-card-section>
    </q-card>
  </q-dialog>
  <q-dialog v-model="testsModal">
    <div
      class="add-group-modal flex column g-m q-pa-lg"
      style="min-width: 500px"
    >
      <div class="text-primary test-title text-bolder">Тесты</div>
      <q-list separator>
        <q-item
          v-for="test in tests"
          :key="test.test_id"
          class="text-primary flex column g-m q-px-none"
        >
          <div class="dialog-title">Название: {{ test.name }}</div>
          <div class="dialog-title">Описание: {{ test.task_description }}</div>
          <div class="dialog-title">Оценка: {{ test.default_mark }}</div>
          <div class="dialog-title">
            Длительность: {{ test.minutes_duration }}
          </div>
        </q-item>
      </q-list>
      <div class="text-primary test-title text-bolder">Лабораторные работы</div>
      <q-list separator>
        <q-item
          v-for="lab in labs"
          :key="lab.laboratory_id"
          class="text-primary flex column g-m q-px-none"
        >
          <div class="dialog-title">Название: {{ lab.name }}</div>
          <div class="dialog-title">Описание: {{ lab.task_description }}</div>
          <div class="dialog-title">Оценка: {{ lab.default_mark }}</div>
        </q-item>
      </q-list>
    </div>
  </q-dialog>
  <q-dialog v-model="dicsiplineModal">
    <div
      class="add-group-modal flex column g-m q-pa-lg"
      style="min-width: 700px"
      v-if="discipline"
    >
      <lecturer-edit-discipline-modal
        :discipline="discipline"
        @close="closeEditDisciplineModal"
      />
    </div>
  </q-dialog>
</template>

<script lang="ts" setup>
import BannerWithBack from 'src/components/BannerWithBack.vue';
import BannerComponent from 'src/components/BannerComponent.vue';
import { useRoute } from 'vue-router';
import { useDisciplinesStore } from 'src/stores/disciplines';
import { Ref, computed, ref, watch } from 'vue';
import { IGroup } from 'src/models/group/group';
import { useSeminariansStore } from '../../../../../../stores/seminarians';
import { ISeminarian } from '../../../../../../models/seminarian/seminarian';
import DisciplineLessons from './DisciplineLessons.vue';
import { useSectionStore } from '../../../../../../stores/section';
import { useTestsStore } from '../../../../../../stores/test';
import { useLabsStore } from 'src/stores/labs';
import LecturerEditDisciplineModal from './LecturerEditDisciplineModal.vue';

const route = useRoute();

const disciplineId = route.params.disciplineId as string;

const store = useDisciplinesStore();
store.getDisciplineInfo(disciplineId);

const sectionStore = useSectionStore();
sectionStore.getSections(Number(disciplineId));

const testsStore = useTestsStore();
const labsStore = useLabsStore();

const tests = computed(() => testsStore.sectionTests);
const labs = computed(() => labsStore.sectionLabs);

const testsModal = ref(false);

const getSectionTests = async (id: number) => {
  testsModal.value = true;
  await testsStore.getSectionTests(id);
  getSectionLabs(id);
};

const closeEditDisciplineModal = async () => {
  await store.getDisciplineInfo(disciplineId);
  await store.getDisciplineInfo(disciplineId);
  await sectionStore.getSections(Number(disciplineId));

  dicsiplineModal.value = false;
};

const getSectionLabs = async (id: number) => {
  testsModal.value = true;
  await labsStore.getLabsFromSection(id);
};

const sections = computed(() => sectionStore.sections);

const discipline = computed(() => store.discipline);
const seminarians = computed(() => store.seminarians);

const groups = computed(() => store.groups);
const avaliableGroups = computed(() => store.avalibaleGroups);
const addGroupModal = ref(false);
const groupToAdd: Ref<IGroup | null> = ref(null);

const disciplineName = ref(discipline.value?.discipline_name);
const disciplineExamMark = ref(discipline.value?.exam_mark);
const disciplineSeminarMark = ref(discipline.value?.seminar_visiting_mark);
const disciplineLessonMark = ref(discipline.value?.lesson_visiting_mark);

watch(discipline, () => {
  disciplineName.value = discipline.value?.discipline_name;
  disciplineExamMark.value = discipline.value?.exam_mark;
  disciplineSeminarMark.value = discipline.value?.seminar_visiting_mark;
  disciplineLessonMark.value = discipline.value?.lesson_visiting_mark;
  getAllGroups();
});

const dicsiplineModal = ref(false);

const changeDisciplineName = async (name: string) => {
  await store.changeDisciplineName({
    name: name,
    discipline_id: Number(disciplineId),
    name_en: name,
  });
  await store.getDisciplineInfo(disciplineId);
};

const changeDisciplineExamMark = async (mark: number) => {
  await store.changeDisciplineExamMark({
    discipline_id: Number(disciplineId),
    lesson_mark: Number(mark),
  });
  await store.getDisciplineInfo(disciplineId);
};

const changeDisciplineSeminarMark = async (mark: number) => {
  await store.changeDisciplineSeminarMark({
    discipline_id: Number(disciplineId),
    seminar_mark: Number(mark),
  });
  await store.getDisciplineInfo(disciplineId);
};

const changeDisciplineLessonMark = async (mark: number) => {
  await store.changeDisciplineLessonMark({
    discipline_id: Number(disciplineId),
    lesson_mark: Number(mark),
  });
  await store.getDisciplineInfo(disciplineId);
};

const getAllGroups = async () => {
  if (discipline.value?.id) {
    await store.getAllGroups(discipline.value?.id);
  }
};

const openAddGroupModal = async () => {
  if (discipline.value?.id) {
    await store.getAvalibaleGroups(discipline.value?.id);
  }
  addGroupModal.value = true;
};

const addGroupToDiscipline = async () => {
  if (groupToAdd.value && discipline.value?.id) {
    await store.addGroupToDiscipline(
      String(discipline.value.id),
      groupToAdd.value.group_id
    );
    await getAllGroups();
    groupToAdd.value = null;
    addGroupModal.value = false;
  }
};

const selectedGroup: Ref<string | null> = ref(null);

const deleteGroupFromDiscipline = async (group_id: string) => {
  if (discipline.value?.id) {
    await store.deleteGroupFromDiscipline(
      String(discipline.value.id),
      group_id
    );
    await getAllGroups();
  }
};

const groupSeminariansModal = ref(false);

const getGroupSeminarians = async (groupId: string) => {
  if (discipline.value?.id) {
    selectedGroup.value = groupId;
    await store.getGroupSeminarians(String(discipline.value?.id), groupId);
    groupSeminariansModal.value = true;
  }
};

const addSeminarian = async () => {
  if (selectedGroup.value) {
    await store.addSeminarianToGroup({
      group_id: Number(selectedGroup.value),
      seminarian_id: Number(selectedSeminarian.value?.seminarian_id),
      discipline_id: Number(disciplineId),
    });
    addSeminariansOpenBoolean.value = false;
    await getGroupSeminarians(selectedGroup.value);
    selectedSeminarian.value = null;
  }
};

const seminariansStore = useSeminariansStore();

const addSeminariansOpenBoolean = ref(false);

const allSeminarians = computed(() => seminariansStore.seminarians);

const avaliableSeminarians: Ref<ISeminarian.Seminarian[]> = ref([]);

const selectedSeminarian: Ref<ISeminarian | null> = ref(null);

const addSeminariansOpen = async () => {
  await seminariansStore.fetchSeminarians();
  if (allSeminarians.value) {
    console.log(allSeminarians.value, seminarians.value);
    avaliableSeminarians.value = allSeminarians.value.filter((sem) => {
      return !seminarians.value?.find(
        (s) => s.seminarian_id === sem.seminarian_id
      );
    });
  }
  addSeminariansOpenBoolean.value = true;
};
</script>

<style lang="scss" scoped>
.input-numbers {
  width: 45%;
}

.discipline-edit {
  font-weight: 600;
  font-size: 20px;
}

.number-edit {
  width: 100%;
  border-bottom: 1px solid var(--q-primary);
}

.add-group-modal {
  width: 300px;
  background-color: white;
  padding: 10px;
}

.dialog-title {
  font-weight: 500;
  font-size: 18px;
}

.test-title {
  font-weight: 600;
  font-size: 20px;
}
</style>
