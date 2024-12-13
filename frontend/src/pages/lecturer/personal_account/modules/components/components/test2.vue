<template>
  <q-dialog v-model="modelValue" transition-show="scale" transition-hide="scale">
    <div class="q-pa-md">
      <h3>Модальное окно</h3>
      <q-btn label="Закрыть" color="primary" @click="$emit('update:modelValue', false)" />
    </div>
  </q-dialog>
</template>

<script>
export default {
  props: {
    modelValue: {
      type: Boolean,
      required: true,
    },
  },
};
</script>

<style scoped>
.q-dialog {
  text-align: center;
}
</style>
