<template>
  <div>
    <q-btn label="Открыть модалку" @click="openTestModal = true" />
    <ModalDialog v-model="openTestModal" />
  </div>
</template>

<script>
import ModalDialog from "./ModalDialog.vue";

export default {
  components: { ModalDialog },
  data() {
    return {
      openTestModal: false,
    };
  },
};
</script>
