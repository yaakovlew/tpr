<template>
  <div class="main flex items-center column q-ma-xs-md">
    <h4>Новости</h4>
    <div class="main__news-container flex items-center column">
      <single-news-container title="Лекции и семинары">
        <template #content>
          <div>
            Видеозаписи лекций и семинаров прошлых лет находятся по адресу
          </div>
        </template>
      </single-news-container>
    </div>
  </div>
</template>

<script setup lang="ts">
import SingleNewsContainer from './components/SingleNewsContainer.vue';
</script>

<style scoped lang="scss">
.main {
  &__news-container {
    width: 100%;
    gap: 15px;
  }
}
</style>
