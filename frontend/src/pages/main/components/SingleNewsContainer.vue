<template>
  <div class="news-container">
    <div class="news-container__title q-pa-md">
      <slot name="title">
        <span v-if="title !== undefined">
          {{ title }}
        </span>
      </slot>
    </div>
    <div class="news-container__content">
      <slot name="content" />
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title?: string;
}>();
</script>

<style scoped lang="scss">
@import '../../../css/quasar.variables.scss';
.news-container {
  width: 100%;
  border: 2px solid $secondary;

  &__title {
    width: 100%;
    background-color: $secondary;
    font-size: 18px;
    color: white;
    text-align: center;
  }

  &__content {
    margin: 10px 20px;
    font-size: 16px;
    text-align: justify;
    font-weight: 500;
  }
}
</style>
