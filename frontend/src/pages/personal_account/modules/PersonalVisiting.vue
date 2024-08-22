<template>
  <div class="flex column g-m">
    <banner-component class="text-primary results-title text-center">
      Посещения
    </banner-component>
    <banner-component>
      <q-table
        :rows="rows"
        :columns="columns"
        row-key="name"
        hide-bottom
        flat
        table-header-class="results-table-header"
      >
        <template v-slot:body="props">
          <q-tr :props="props" @click="onRowClick(props.row)">
            <q-td key="name" :props="props">
              {{ props.row.name }}
            </q-td>
            <q-td key="date" :props="props">
              {{ props.row.date }}
            </q-td>
            <q-td key="result" :props="props">
              <q-icon
                :name="props.row.result ? 'done' : 'close'"
                size="32px"
                :color="props.row.result ? 'green' : 'negative'"
              />
            </q-td>
          </q-tr>
        </template>
      </q-table>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from '../../../components/BannerComponent.vue';

const columns = [
  {
    name: 'name',
    label: 'Название',
    field: 'name',
    align: 'left',
  },
  {
    name: 'date',
    label: 'Дата',
    field: 'date',
  },
  {
    name: 'result',
    label: 'Посещение',
    field: 'result',
  },
];

const rows = [
  {
    name: 'Мягкие исчисления',
    date: '2020-01-01',
    result: true,
  },
  {
    name: 'Мягкие исчисления',
    date: '2020-01-01',
    result: false,
  },
  {
    name: 'Мягкие исчисления',
    date: '2020-01-01',
    result: false,
  },
  {
    name: 'Мягкие исчисления',
    date: '2020-01-01',
    result: true,
  },
];
</script>

<style lang="scss" scoped>
.results-title {
  font-size: 24px;
  font-weight: bold;
}
</style>
<style lang="scss">
.results-table-header {
  th {
    font-size: 16px;
  }
}
</style>
