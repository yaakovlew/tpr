<template>
  <q-item class="q-px-none">
    <q-item-section class="row">
      <q-item-label class="row items-center justify-between full-width">
        <div class="group-name">
          {{ seminarian.name }}
          {{ seminarian.surname }}
        </div>
      </q-item-label>
    </q-item-section>
    <q-item-section class="row">
      <q-item-label class="row items-center justify-between full-width">
        <div class="group-name">
          {{ seminarian.email }}
        </div>
      </q-item-label>
    </q-item-section>
    <!-- <q-item-section class="row">
      <q-item-label class="row items-center justify-end">
        <q-icon
          name="info"
          color="primary"
          size="20px"
          class="cursor-pointer"
        />
      </q-item-label>
    </q-item-section> -->
  </q-item>
</template>

<script lang="ts" setup>
import { ISeminarian } from 'src/models/seminarian/seminarian';

defineProps<{
  seminarian: ISeminarian;
}>();
</script>

<style lang="scss" scoped>
.group-name {
  font-size: 18px;
}
.dialog-title {
  font-weight: 500;
  font-size: 18px;
}
</style>
