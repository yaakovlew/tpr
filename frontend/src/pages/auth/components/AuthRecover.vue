<template>
  <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
    <q-input
      v-model="password"
      filled
      :type="isPwd ? 'password' : 'text'"
      label="Пароль*"
      lazy-rules
      :rules="[
        (val: string) => (val && val.length > 0) || 'Поле должно быть заполненым',
      ]"
    >
      <template v-slot:append>
        <q-icon
          :name="isPwd ? 'visibility_off' : 'visibility'"
          class="cursor-pointer"
          @click="isPwd = !isPwd"
        />
      </template>
    </q-input>
    <q-input
      v-model="confirmationPassword"
      filled
      :type="isPwdConfirmation ? 'password' : 'text'"
      label="Подтверждение пароля*"
      lazy-rules
      :rules="[(val: string) => (val && val === password) || 'Поля дожны совпадать']"
    >
      <template v-slot:append>
        <q-icon
          :name="isPwdConfirmation ? 'visibility_off' : 'visibility'"
          class="cursor-pointer"
          @click="isPwdConfirmation = !isPwdConfirmation"
        />
      </template>
    </q-input>
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
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { Notify } from 'quasar';

const store = useAuthStore();

const password = ref('');
const confirmationPassword = ref(null);
const isPwd = ref(true);
const isPwdConfirmation = ref(true);

const route = useRoute();

const onSubmit = async () => {
  if (password.value && route.query.token) {
    await store.restorePassword(
      password.value,
      (route.query.token as string) ?? ''
    );
    Notify.create({
      type: 'positive',
      message:
        'Пароль отправлен вам на почту, если сообщения нет, проверьте вкладку спам',
    });
  }
};

const onReset = () => {
  password.value = '';
  confirmationPassword.value = null;
  isPwd.value = true;
  isPwdConfirmation.value = true;
};
</script>
