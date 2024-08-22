<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Семинаристы
    </banner-component>
    <template v-if="seminarians">
      <banner-component>
        <q-list class="q-pa-none">
          <template
            v-for="(seminarian, index) in seminarians"
            :key="seminarian.seminarian_id"
          >
            <q-separator
              v-if="index !== 0"
              spaced
              inset
              class="q-px-none q-mx-none"
            />
            <lecturer-seminarian :seminarian="seminarian" />
          </template>
        </q-list>
      </banner-component>
    </template>
    <banner-component v-else class="text-primary profile-name">
      Нет семинаристов
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from 'src/components/BannerComponent.vue';
import LecturerSeminarian from './components/LecturerSeminarian.vue';
import { useSeminariansStore } from 'src/stores/seminarians';
import { computed } from 'vue';

const seminariansStore = useSeminariansStore();

const seminarians = computed(() => seminariansStore.seminarians);

seminariansStore.fetchSeminarians();
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}
</style>
