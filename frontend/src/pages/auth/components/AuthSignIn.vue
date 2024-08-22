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
    <div>
      <q-btn label="Авторизироваться" type="submit" color="primary" />
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
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

const store = useAuthStore();

const auth = computed(() => store.state);
const userType = computed(() => store.userType);

const router = useRouter();

const email = ref('');
const password = ref('');
const isPwd = ref(true);

if (userType.value) {
  switch (userType.value) {
    case 'student':
      router.replace({ name: 'student-profile' });
      break;
    case 'lecturer':
      router.replace({ name: 'lecturer-profile' });
      break;
    case 'seminarian':
      router.replace({ name: 'seminarian-profile' });
      break;
    default:
      break;
  }
}

const onSubmit = async () => {
  await store.login({
    email: email.value,
    password: password.value,
  });

  if (userType.value) {
    switch (userType.value) {
      case 'student':
        router.replace({ name: 'student-profile' });
        break;
      case 'lecturer':
        console.log('here');
        router.replace({ name: 'lecturer-profile' });
        break;
      case 'seminarian':
        router.replace({ name: 'seminarian-profile' });
        break;
      default:
        break;
    }
  }
};

const onReset = () => {
  email.value = '';
  password.value = '';
  isPwd.value = true;
};
</script>
