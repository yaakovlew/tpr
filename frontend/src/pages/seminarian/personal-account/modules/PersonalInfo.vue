<template>
  <div name="profile" class="flex column bg-none g-m">
    <div class="flex g-m items-stretch">
      <banner-component class="text-primary profile-name">
        {{ me?.name + ' ' + me?.surname }}
      </banner-component>
      <banner-component class="exit-icon cursor-pointer" @click="logout">
        <q-icon
          name="exit_to_app"
          class="full-height"
          size="2em"
          color="primary"
        />
      </banner-component>
    </div>

    <banner-component class="flex column g-m">
      <text-with-caption required>
        <template #caption> Фамилия </template>
        <span v-if="!isInEditMode">
          {{ me?.surname }}
        </span>
        <q-input v-else v-model="editData.surname" />
      </text-with-caption>
      <text-with-caption required>
        <template #caption> Имя </template>
        <span v-if="!isInEditMode">
          {{ me?.name }}
        </span>
        <q-input v-else v-model="editData.name" />
      </text-with-caption>
      <text-with-caption v-if="!isInEditMode" required>
        <template #caption> Электронная почта </template>
        <span>
          {{ me?.email }}
        </span>
      </text-with-caption>
    </banner-component>
    <banner-component
      v-if="isInEditMode"
      @click="saveEdit"
      class="text-center cursor-pointer text-primary edit-button flex items-center relative-position"
    >
      <q-icon name="edit" />
      <div class="flex-grow">Сохранить</div>
    </banner-component>
    <banner-component
      v-if="isInEditMode"
      @click="closeEdit"
      class="text-center cursor-pointer text-primary edit-button flex items-center relative-position"
    >
      <q-icon name="close" />
      <div class="flex-grow">Отменить</div>
    </banner-component>
    <banner-component
      v-if="!isInEditMode"
      @click="isInEditMode = !isInEditMode"
      class="text-center cursor-pointer text-primary edit-button flex items-center relative-position"
    >
      <q-icon name="edit" />
      <div class="flex-grow">Редактировать профиль</div>
    </banner-component>
    <banner-component
      v-if="!isInEditMode"
      @click="isPasswordEdit = !isPasswordEdit"
      class="text-center cursor-pointer text-primary edit-button flex items-center relative-position"
    >
      <q-icon name="edit" />
      <div class="flex-grow">Изменить пароль</div>
    </banner-component>
    <q-dialog v-model="isPasswordEdit" persistent @hide="closeChangePassword">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Изменение пароля</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-form @submit="changePassword" class="q-gutter-md">
            <q-input
              v-model="changePasswordData.old_password"
              filled
              :type="isPwd ? 'password' : 'text'"
              label="Старый пароль*"
              lazy-rules
              :rules="[
                (val: string) =>
                  (val && val.length > 0) || 'Поле должно быть заполненым',
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
              v-model="changePasswordData.new_password"
              filled
              :type="isPwdConfirmation ? 'password' : 'text'"
              label="Новый пароль*"
              lazy-rules
              :rules="[
                (val: string) =>
                  (val && val.length > 0) || 'Поле должно быть заполненым',
              ]"
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
              <q-btn label="Поменять пароль" type="submit" color="primary" />
              <q-btn
                label="Отменить"
                color="primary"
                flat
                class="q-ml-sm"
                @click="closeChangePassword"
              />
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import { usePersonalInfoStore } from 'src/stores/personalInfo';
import BannerComponent from 'src/components/BannerComponent.vue';
import TextWithCaption from '../components/TextWithCaption.vue';
import { computed, reactive, ref, watch } from 'vue';
import { useAuthStore } from 'src/stores/auth';

const store = usePersonalInfoStore();
const authStore = useAuthStore();

store.getPersonalInfo();

const me = computed(() => store.me);

const editData = reactive({
  name: me.value?.name,
  surname: me.value?.surname,
});

const logout = () => {
  authStore.logout();
};

const isInEditMode = ref(false);

watch(isInEditMode, () => {
  if (isInEditMode.value) {
    editData.name = me.value?.name;
    editData.surname = me.value?.surname;
  }
});

watch(
  () => me,
  () => {
    editData.name = me.value?.name;
    editData.surname = me.value?.surname;
  }
);

const saveEdit = async () => {
  await store.changeName(editData.name ?? '', editData.surname ?? '');
  await store.getPersonalInfo();
  closeEdit();
};

const closeEdit = () => {
  isInEditMode.value = false;
};

const changePasswordData = reactive({
  new_password: '',
  old_password: '',
});

const isPasswordEdit = ref(false);

const isPwd = ref(true);
const isPwdConfirmation = ref(true);

const changePassword = async () => {
  await store.changePassword(
    changePasswordData.old_password,
    changePasswordData.new_password
  );
  closeChangePassword();
};

const closeChangePassword = () => {
  isPasswordEdit.value = false;
  changePasswordData.new_password = '';
  changePasswordData.old_password = '';
};
</script>

<style lang="scss" scoped>
.bg-none {
  background: none;
}

.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}

.exit-icon {
  font-weight: 600;
}

.edit-button {
  font-size: 22px;
  font-weight: 600;
}
</style>
