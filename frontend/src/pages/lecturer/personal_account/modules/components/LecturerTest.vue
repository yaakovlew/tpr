<template>
  <q-tr class="text-primary test">
    <q-td key="question" class="test">
      {{ test.name }}
      <q-popup-edit v-model="testName" v-slot="scope">
        <q-input
          autofocus
          dense
          v-model="scope.value"
          :model-value="scope.value"
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
              @click.stop.prevent="
                changeTestDescription(
                  test.task_description,
                  test.minutes_duration,
                  scope.value
                );
                scope.set();
              "
              :disable="scope.value === ''"
            />
          </template>
        </q-input>
      </q-popup-edit>
    </q-td>
    <q-td key="task_description" class="test">
      {{ test.task_description }}
      <q-popup-edit v-model="testDescription" v-slot="scope">
        <q-input
          autofocus
          dense
          v-model="scope.value"
          :model-value="scope.value"
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
              @click.stop.prevent="
                changeTestDescription(
                  scope.value,
                  test.minutes_duration,
                  test.name
                );
                scope.set();
              "
              :disable="scope.value === ''"
            />
          </template>
        </q-input>
      </q-popup-edit>
    </q-td>
    <q-td key="minutes_duration" class="test">
      {{ test.minutes_duration }}
      <q-popup-edit v-model="testMinutesDuration" v-slot="scope">
        <q-input
          autofocus
          dense
          v-model="scope.value"
          :model-value="scope.value"
          type="number"
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
              @click.stop.prevent="
                changeTestDescription(
                  test.task_description,
                  scope.value,
                  test.name
                );
                scope.set();
              "
              :disable="scope.value === ''"
            />
          </template>
        </q-input>
      </q-popup-edit>
    </q-td>
    <q-td key="default_mark" class="test">
      {{ test.default_mark }}
    </q-td>
    <q-td key="actions">
      <q-btn icon="list" color="primary" flat @click="getTestThemes" />
      <q-btn icon="delete" color="red" flat @click="deleteTest" />
    </q-td>
  </q-tr>
  <q-dialog v-model="themesDialog">
    <div class="modal flex q-pa-md column g-m no-wrap">
      <div class="profile-name text-primary">Темы</div>
      <q-table
        :columns="columns"
        :rows="themesWithCount"
        :pagination="pagination"
        hide-bottom
        dense
        flat
      >
        <template #body-cell="props">
          <q-td :props="props" style="width: 30px">
            <div v-if="props.col.name === 'actions'" style="width: 30px">
              <q-icon
                color="negative"
                name="delete"
                size="18px"
                class="cursor-pointer"
                @click="deleteThemeFromTest(props.row.theme_id)"
                style="width: 30px"
              />
            </div>
            <div v-else-if="props.col.name === 'count'">
              {{ `${props.row.count} / ${props.row.total_count}` }}
              <q-popup-edit v-model="props.row.count" v-slot="scope">
                <q-input
                  autofocus
                  dense
                  v-model="scope.value"
                  type="number"
                  :model-value="scope.value"
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
                      @click.stop.prevent="
                        changeThemeCount(props.row.theme_id, scope.value);
                        scope.set();
                      "
                      :disable="scope.value === ''"
                    />
                  </template>
                </q-input>
              </q-popup-edit>
            </div>
            <div v-else class="name-in-table">
              {{ props.value }}
            </div>
          </q-td>
        </template>
      </q-table>
      <q-btn
        v-if="!addThemeForm && avalibaleThemes.length"
        label="Добавить тему"
        @click="addThemeForm = true"
        flat
        color="primary"
      />
      <div v-else-if="avalibaleThemes.length" class="flex column g-m">
        <q-select
          v-model="themeToAdd"
          :options="avalibaleThemes"
          option-label="name"
          option-value="theme_id"
          label="Тема"
        />
        <q-input type="number" v-model="count" label="Колчиство вопросов" />
        <q-btn
          label="Добавить тему"
          color="primary"
          flat
          :disable="!themeToAdd || !count"
          @click="addTheme"
        />
      </div>
    </div>
  </q-dialog>
</template>

<script lang="ts" setup>
import { ITest } from 'src/models/test/test';
import { useTestsStore } from '../../../../../stores/test';
import { computed, ref, Ref } from 'vue';
import { ITheme } from 'src/models/test/theme';
const props = defineProps<{
  test: ITest.Test;
}>();

const store = useTestsStore();

const themes = computed(() => store.themes);

const avalibaleThemes = computed(() =>
  store.allThemes.filter(
    (theme) => !themes.value?.find((t) => t.theme_id === theme.theme_id)
  )
);

const getQuestionsCount = async (id: number) => {
  const count = await store.getQuestionsCount(id);
  return count;
};

const pagination = {
  rowsPerPage: 0,
};

const themesWithCount: Ref<ITheme.ThemeTestWithTotalCount[]> = ref([]);

const getThemesWithQuestionsCount = async () => {
  const allPromises: Promise<Record<number, number> | undefined>[] = [];
  themes.value?.forEach((theme) => {
    allPromises.push(store.getQuestionsCount(theme.theme_id));
  });
  const all = await Promise.all(allPromises);
  themesWithCount.value = themes.value?.map((theme) => ({
    ...theme,
    total_count: all.find((t) => t?.[theme.theme_id])?.[theme.theme_id] ?? 0,
  }));
};

const getTestThemes = async () => {
  themesDialog.value = true;
  await store.getThemes(props.test.test_id);
  await getThemesWithQuestionsCount();
};

const themesDialog = ref(false);

const addThemeForm = ref(false);

const themeToAdd: Ref<ITheme.Theme | null> = ref(null);
const count = ref(0);

const testName = ref(props.test.name);
const testDescription = ref(props.test.task_description);
const testMinutesDuration = ref(props.test.minutes_duration);
const testDefaultMark = ref(props.test.default_mark);

// TODO: переделать для одного метода changeTest

const changeTestDescription = async (
  description: string,
  minutes: number,
  name: string
) => {
  await store.changeTestDescription({
    test_id: props.test.test_id,
    description: description,
    description_en: description,
    minutes_duration: minutes,
    name: name,
    name_en: name,
  });
  await store.getTests();
};

const addTheme = async () => {
  if (themeToAdd.value) {
    await store.addThemeToTest({
      test_id: props.test.test_id,
      theme_id: themeToAdd.value.theme_id,
      count: Number(count.value),
    });
  }
  themeToAdd.value = null;
  addThemeForm.value = false;
  await store.getThemes(props.test.test_id);
  await getThemesWithQuestionsCount();
};

const deleteThemeFromTest = async (themeId: number) => {
  await store.deleteThemeFromTest({
    test_id: props.test.test_id,
    theme_id: themeId,
  });
  await store.getThemes(props.test.test_id);
  await getThemesWithQuestionsCount();
};

const changeThemeCount = async (themeId: number, count: number) => {
  await store.changeThemeCount({
    test_id: props.test.test_id,
    theme_id: themeId,
    count: Number(count),
  });
  await store.getThemes(props.test.test_id);
};

const deleteTest = async () => {
  await store.deleteTest(props.test.test_id);
  await store.getTests();
};

const columns = [
  {
    name: 'name',
    label: 'Название',
    field: 'name',
    align: 'left',
    sortable: true,
  },
  {
    name: 'count',
    label: 'Количество вопросов',
    field: 'count',
    align: 'left',
  },
  {
    name: 'actions',
    label: '',
    field: 'actions',
    align: 'left',
  },
];
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 20px;
  font-weight: 500;
}
.modal {
  min-width: 600px;
  background-color: white;
}

.test {
  font-size: 18px;
  font-weight: 500;
  width: 300px;
}

.name-in-table {
  width: 300px;
  white-space: break-spaces;
}
</style>
