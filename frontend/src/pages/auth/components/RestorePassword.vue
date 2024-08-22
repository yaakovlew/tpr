<template>
  <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
    <q-input
      filled
      v-model="email"
      type="email"
      label="Email*"
      lazy-rules
      :rules="[
          (val: string) => (val && val.length > 0) || 'Поле должно быть заполненым',
        ]"
    />
    <div>
      <q-btn label="Восстановить" type="submit" color="primary" />
      <q-btn
        label="Обновить"
        type="reset"
        color="primary"
        flat
        class="q-ml-sm"
      />
    </div>
  </q-form>
</template>

<script setup lang="ts">
import { useAuthStore } from 'src/stores/auth';
import { ref } from 'vue';

const email = ref(null);

const store = useAuthStore();

const onSubmit = async () => {
  if (email.value) {
    await store.forgetPassword(email.value);
  }
  console.log('submit');
};

const onReset = () => {
  email.value = null;
};
</script>
