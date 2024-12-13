<template>
  <div>
    <q-btn label="Открыть модалку" @click="openTestModal = true" />
    <q-dialog v-model="openTestModal" transition-show="scale" transition-hide="scale">
      <div class="q-pa-md">
        <h3>Модальное окно</h3>
        <q-btn label="Закрыть" color="primary" @click="openTestModal = false" />
      </div>
    </q-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      openTestModal: false,
    };
  },
};
</script>

<style scoped>
.q-dialog {
  text-align: center;
}
</style>
