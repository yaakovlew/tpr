<template>
  <div class="q-pa-md flex column g-m">
    <banner-component>
      <div class="page-title text-primary">
        Дисциплина: {{ discipline?.discipline_name }}
      </div>
      <div class="page-title text-primary">
        Группа: {{ currentGroup?.name }}
      </div>
      <div class="primary-text">
        Максимальное количество баллов за экзамен: {{ discipline?.exam_mark }}
      </div>
      <div class="primary-text">
        Максимальное количество баллов за посещение лекций:
        {{ discipline?.lesson_visiting_mark }}
      </div>
    </banner-component>
    <banner-component
      class="page-title text-primary cursor-pointer"
      @click="getReport"
    >
      Получить отчет
    </banner-component>
    <banner-component class="page-title text-primary">
      Оценки
    </banner-component>
    <banner-component>
      <q-table
        :columns="columns"
        :rows="rows"
        hide-bottom
        flat
        :pagination="pagination"
      >
        <template #body-cell-exam="props">
          <q-td key="name" :props="props">
            {{ props.row.exam }}
            <q-popup-edit v-model="props.row.exam" v-slot="scope">
              <q-input
                v-model="scope.value"
                dense
                autofocus
                counter
                type="number"
                @keyup.enter="
                  changeExamMark(props.row.id, scope.value), scope.set()
                "
              />
            </q-popup-edit>
          </q-td>
        </template>
        <template
          v-for="test in allTests"
          #[`body-cell-${test.name}`]="props"
          :key="test.test_id"
        >
          <q-td :key="test.test_id" :props="props">
            {{ props.row[test.name] ?? 'Не пройден' }}
            <q-icon
              v-if="props.row[test.name] !== null"
              name="download"
              color="primary"
              size="18px"
              class="cursor-pointer"
              @click.stop="
                getTestReport(
                  test.test_id,
                  props.row.id,
                  `${props.row.name}_${props.row.surname}`
                )
              "
            >
              <q-tooltip> Скачать отчет </q-tooltip>
            </q-icon>
            <q-popup-edit v-model="props.row[test.name]" v-slot="scope">
              <q-input
                v-model="scope.value"
                dense
                autofocus
                counter
                type="number"
                @keyup.enter="
                  changeTestMark(props.row.id, scope.value, test.test_id),
                    scope.set()
                "
              />
            </q-popup-edit> </q-td
        ></template>
      </q-table>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { Ref, computed, onMounted, ref, watch } from 'vue';
import { useMarksStore } from '../../../../../../stores/mark';
import { useGroupsStore } from '../../../../../../stores/groups';
import { useTestsStore } from 'src/stores/test';
import { useSectionStore } from '../../../../../../stores/section';
import { ITest } from 'src/models/test/test';
import BannerComponent from 'src/components/BannerComponent.vue';
import { useDisciplinesStore } from 'src/stores/disciplines';
import { useReportStore } from 'src/stores/report';
import { download } from 'src/utils/download';
import { IGroup } from 'src/models/group/group';

const route = useRoute();

const disciplineId = computed(() => route.params.disciplineId);
const groupId = computed(() => route.params.groupId);

const report = useReportStore();

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

const getReport = async () => {
  const res = await report.getReport(
    Number(groupId.value),
    Number(disciplineId.value),
    true
  );
  saveFile(
    res.data,
    `Отчет_${currentGroup.value?.name}_${discipline.value?.discipline_name}`
  );
  // download('Отчет.xlsx', res.data);
};

const store = useMarksStore();
const groupStore = useGroupsStore();
const testsStore = useTestsStore();
const sectionStore = useSectionStore();
const disciplineStore = useDisciplinesStore();

const allSectionsTests = computed(() => testsStore.allSectionsTests);
const tests = computed(() => testsStore.sectionTests);
const sections = computed(() => sectionStore.sections);
const students = computed(() => groupStore.groupStudents);
const examMarks = computed(() => store.examMarks);
const allTestsMarks = computed(() => store.allTestsMarks);
const discipline = computed(() => disciplineStore.discipline);
const groups = computed(() => groupStore.groups);

const changeExamMark = async (userId: number, mark: number) => {
  await store.postExamMark({
    discipline_id: Number(disciplineId.value),
    user_id: Number(userId),
    mark: Number(mark),
  });
  await getInfo();
};

const changeTestMark = async (userId: number, mark: number, testId: number) => {
  await store.postTestMark({
    user_id: userId,
    mark: Number(mark),
    test_id: testId,
  });
  await getInfo();
};

const getTestReport = async (testId: number, userId: number, name: string) => {
  const fileName = `Отчет_${currentGroup.value?.name}_${discipline.value?.discipline_name}_${name}.txt`;
  await testsStore.getTestReport(
    {
      test_id: testId,
      user_id: userId,
    },
    fileName
  );
};

const columns = computed(() => {
  const columns: any = [
    {
      name: 'name',
      field: 'name',
      label: 'Имя',
      align: 'left',
    },
    {
      name: 'surname',
      field: 'surname',
      label: 'Фамилия',
      align: 'left',
    },
  ];

  allTests.value?.forEach((test) => {
    columns.push({
      name: test.name,
      field: test.name,
      label: `Тест: ${test.name} (${test.default_mark})`,
      align: 'left',
    });
  });

  columns.push({
    name: 'exam',
    field: 'exam',
    label: 'Экзамен',
    align: 'left',
  });

  return columns;
});

const rows = computed(() => {
  const rows: any = [];

  students.value?.forEach((student) => {
    const obj: Record<string, number | string | null> = {
      name: student.name,
      surname: student.surname,
      exam:
        examMarks.value?.find(
          (mark) => mark.user_id === Number(student.student_id)
        )?.mark ?? 0,
      id: student.student_id,
    };
    allTests.value?.forEach((test) => {
      obj[test.name] =
        allTestsMarks.value[test.test_id]?.find(
          (test) => test.user_id === Number(student.student_id)
        )?.mark ?? null;
    });
    rows.push(obj);
  });

  return rows;
});

const allTests = computed(() => {
  const array = Object.values(allSectionsTests.value)?.map((tests) => tests);
  const res: ITest.Test[] = [];
  array?.forEach((arr) => {
    if (arr) {
      res.push(...arr);
    }
  });
  return res;
});

const pagination = {
  rowsPerPage: 0,
};

const getInfo = async () => {
  await groupStore.getGroupStudents(String(groupId.value));
  store.allTestsMarks = {};
  testsStore.allSectionsTests = {};
  await store.getExamMarks({
    group_id: Number(groupId.value),
    discipline_id: Number(disciplineId.value),
  });
  await sectionStore.getSections(Number(disciplineId.value));
  if (sections.value) {
    await testsStore.getAllSectionsTests(
      sections.value?.map((s) => s.section_id)
    );
  }
  await store.getTestsMarks(
    allTests.value.map((t) => t.test_id),
    Number(groupId.value)
  );
};

const currentGroup: Ref<IGroup.Group | null> = ref(null);

watch(groupId, async () => {
  if (!Array.isArray(groupId.value)) {
    await groupStore.getGroups();
    const gr = groups.value?.find(
      (g) => String(g.group_id) === String(groupId.value)
    );
    if (gr) {
      currentGroup.value = gr;
    }
  }
});

onMounted(async () => {
  await disciplineStore.getDisciplineInfo(String(disciplineId.value));
  await groupStore.getGroups();
  await getInfo();
  if (groupId.value) {
    currentGroup.value =
      groups.value?.find((g) => g.group_id == groupId.value) ?? null;
  }
});
</script>

<style lang="scss" scoped></style>
