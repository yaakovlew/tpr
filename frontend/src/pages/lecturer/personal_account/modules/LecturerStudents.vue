<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Студенты
    </banner-component>
    <template v-if="students">
      <banner-component>
        <q-table
          :columns="columns"
          :rows="students"
          :pagination="pagination"
          hide-bottom
          :filter="filter"
          dense
          flat
        >
          <template v-slot:top-right>
            <q-input
              borderless
              dense
              debounce="300"
              v-model="filter"
              placeholder="Поиск"
            >
              <template v-slot:append>
                <q-icon name="search" />
              </template>
            </q-input>
          </template>
          <template #body-cell="props">
            <q-td :props="props" class="text-primary group-name">
              <template v-if="props.col.name !== 'actions'">
                {{ props.value }}
              </template>
              <template v-else>
                <lecturer-student :student="props.row" />
              </template>
            </q-td>
          </template>
        </q-table>
      </banner-component>
    </template>
    <template v-else>
      <banner-component class="text-primary profile-name">
        Нет студентов
      </banner-component>
    </template>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from '../../../../components/BannerComponent.vue';
import LecturerStudent from './components/LecturerStudent.vue';
import { computed, ref } from 'vue';
import { useStudentsStore } from 'src/stores/students';

const columns = [
  {
    name: 'name',
    label: 'Имя',
    sortable: true,
    field: 'name',
    align: 'left',
  },
  {
    name: 'surname',
    label: 'Фамилия',
    sortable: true,
    field: 'surname',
    align: 'left',
    sort: (a: string, b: string) => {
      console.log(a, b);
      if (a.toLowerCase() < b.toLowerCase()) {
        return -1;
      }
      if (a.toLowerCase() > b.toLowerCase()) {
        return 1;
      }
      return 0;
    },
  },
  ,
  {
    name: 'group_name',
    label: 'Группа',
    sortable: true,
    field: 'group_name',
    align: 'left',
  },
  {
    name: 'actions',
    label: 'Действия',
    field: 'actions',
    align: 'left',
  },
];

const filter = ref('');

const pagination = {
  rowsPerPage: 0,
};

const store = useStudentsStore();

store.getStudents();

const students = computed(() => store.students);
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}

.group-name {
  font-size: 18px;
}
</style>
