<template>
  <q-layout view="lHh Lpr lFf" class="main-layout bg-grey-2">
    <q-header elevated class="main-layout__header">
      <q-toolbar>
        <q-toolbar-title to="/" @click="toMain" class="cursor-pointer">
          Справочно обучающая система
        </q-toolbar-title>

        <q-btn to="/auth" flat round dense icon="person" />
      </q-toolbar> </q-header
    >**
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const leftDrawerOpen = ref(false);

const toMain = () => {
  router.push('/');
};

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>

<style scoped lang="scss">
$max-width: 1440px;
.main-layout {
  max-width: $max-width;
  margin: 0 auto;

  &__header {
    max-width: $max-width;
    margin: 0 auto;
  }
}
</style>
