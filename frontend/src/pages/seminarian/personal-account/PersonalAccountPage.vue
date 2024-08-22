<template>
  <div class="flex q-pa-lg">
    <q-tabs
      v-model="tab"
      vertical
      active-bg-color="grey-4"
      :active-class="'mdi-border-radius'"
      class="text-grey-9"
    >
      <q-tab name="profile" label="Профиль" />
      <!-- <q-tab name="results" label="Результаты тестирования" /> -->
      <q-tab name="disciplins" label="Дисциплины" />
    </q-tabs>
    <q-tab-panels v-model="tab" class="bg-none q-pa-none flex-grow">
      <q-tab-panel name="profile" class="q-py-none">
        <personal-info />
      </q-tab-panel>
      <q-tab-panel name="disciplins" class="q-py-none">
        <personal-disciplines />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, watch } from 'vue';
import PersonalInfo from './modules/PersonalInfo.vue';
import PersonalDisciplines from './modules/PersonalDisciplines.vue';
import { useAuthStore } from 'src/stores/auth';
import { useRouter } from 'vue-router';

const tab = ref('profile');

const authStore = useAuthStore();

const userType = computed(() => authStore.userType);

const router = useRouter();

if (userType.value) {
  switch (userType.value) {
    case 'student':
      break;
    case 'lecturer':
      router.replace({ name: 'lecturer-profile' });
      break;
    case 'seminarian':
      router.replace({ name: 'seminarian-profile' });
      break;
    default:
      router.replace({ path: '/auth' });
      break;
  }
}

watch(userType, () => {
  if (userType.value) {
    switch (userType.value) {
      case 'student':
        break;
      case 'lecturer':
        router.replace({ name: 'lecturer-profile' });
        break;
      case 'seminarian':
        router.replace({ name: 'seminarian-profile' });
        break;
      default:
        router.replace({ path: '/auth' });
        break;
    }
  }
});
</script>

<style lang="scss" scoped></style>
