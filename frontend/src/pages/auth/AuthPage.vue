<template>
  <div class="q-pa-md container">
    <q-tabs v-model="tab" dense class="text-white backgrond-primary q-pa-sm">
      <q-tab name="sign-in" label="Войти" />
      <q-tab name="sign-up" label="Зарегистрироваться" />
      <q-tab name="recovery" label="Восстановить" />
    </q-tabs>
    <q-tab-panels v-model="tab" animated>
      <q-tab-panel name="sign-in">
        <auth-sign-in />
      </q-tab-panel>

      <q-tab-panel name="sign-up">
        <auth-sign-up />
      </q-tab-panel>

      <q-tab-panel name="recovery">
        <auth-recover />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import AuthSignIn from './components/AuthSignIn.vue';
import AuthSignUp from './components/AuthSignUp.vue';
import AuthRecover from './components/RestorePassword.vue';

const tab = ref('sign-in');
</script>

<style lang="scss" scoped>
@import '../../css/quasar.variables.scss';
.container {
  max-width: 600px;
  margin: 0 auto;
}
.backgrond-primary {
  background-color: $primary;
}
</style>
