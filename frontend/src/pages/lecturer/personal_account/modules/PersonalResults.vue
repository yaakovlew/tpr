<template>
  <div class="flex column g-m">
    <banner-component class="text-primary results-title text-center">
      Результаты тестирования
    </banner-component>
    <banner-component>
      <q-table
        :rows="rows"
        :columns="columns"
        row-key="name"
        hide-bottom
        flat
        table-header-class="results-table-header"
      />
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from '../../../../components/BannerComponent.vue';

const columns = [
  {
    name: 'name',
    label: 'Название',
    field: 'name',
    align: 'left',
  },
  {
    name: 'result',
    label: 'Количество баллов',
    field: 'result',
  },
];

const rows = [
  {
    name: 'Мягкие исчисления',
    result: 20,
  },
  {
    name: 'Мягкие исчисления',
    result: 20,
  },
  {
    name: 'Мягкие исчисления',
    result: 20,
  },
  {
    name: 'Мягкие исчисления',
    result: 20,
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
