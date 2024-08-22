<template>
  <q-form @submit="onSubmit" @reset="onReset" class="q-gutter-md">
    <q-input
      filled
      v-model="name"
      label="Имя*"
      lazy-rules
      :rules="[
        (val: string) => (val && val.length > 0) || 'Поле должно быть заполненым',
      ]"
    />
    <q-input
      filled
      v-model="surname"
      label="Фамилия*"
      lazy-rules
      :rules="[
        (val: string) => (val && val.length > 0) || 'Поле должно быть заполненым',
      ]"
    />
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
    <q-select
      v-if="!isSeminarian"
      v-model="groupName"
      label="Группа"
      :options="commonGroups"
      :option-value="
                (opt : IGroup) => (Object(opt) === opt && 'name' in opt ? opt.name : null)
              "
      :option-label="
                (opt : IGroup) =>
                  Object(opt) === opt && 'name' in opt ? opt.name : ''
              "
      map-options
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
    <q-checkbox v-model="isSeminarian" label="Семинарист" />
    <div>
      <q-btn label="Зарегистрироваться" type="submit" color="primary" />
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
import { IGroup } from 'src/models/group/group';
import { useAuthStore } from 'src/stores/auth';
import { useGroupsStore } from 'src/stores/groups';
import { computed, ref, Ref } from 'vue';
import { useRouter } from 'vue-router';

const store = useAuthStore();
const groupStore = useGroupsStore();

const router = useRouter();

const userType = computed(() => store.userType);

const isSeminarian = ref(false);

const name = ref('');
const surname = ref('');
const groupName: Ref<IGroup | null> = ref(null);
const type = ref(null);
const email = ref('');
const password = ref('');
const confirmationPassword = ref(null);
const isPwd = ref(true);
const isPwdConfirmation = ref(true);

groupStore.getCommonGroups();

const commonGroups = computed(() => groupStore.commonGroups);

const onSubmit = async () => {
  if (groupName.value?.name) {
    await store.register({
      name: name.value,
      surname: surname.value,
      email: email.value,
      password: password.value,
      post: 'student',
      group_name: groupName.value?.name,
    });
  } else if (isSeminarian.value) {
    await store.register({
      name: name.value,
      surname: surname.value,
      email: email.value,
      password: password.value,
      post: 'seminarian',
    });
  }
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
  name.value = '';
  surname.value = '';
  email.value = '';
  type.value = null;
  password.value = '';
  confirmationPassword.value = null;
  isPwd.value = true;
  isPwdConfirmation.value = true;
};
</script>
