<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Дисциплины
    </banner-component>
    <banner-component v-if="disciplines">
      <q-list class="q-pa-none">
        <template
          v-for="(discipline, index) in disciplines"
          :key="discipline.name"
        >
          <q-separator
            v-if="index !== 0"
            spaced
            inset
            class="q-px-none q-mx-none"
          />

          <lecturer-discipline :discipline="discipline" />
        </template>
      </q-list>
    </banner-component>
    <banner-component v-else class="text-primary profile-name">
      Нет назначенных дисциплин
    </banner-component>
    <banner-component
      class="text-primary profile-name cursor-pointer"
      @click="openNewDiscipline"
    >
      Добавить дисциплину
      <q-icon name="add" color="primary" size="30px" class="cursor-pointer">
        <q-tooltip> Добавить дисциплину </q-tooltip>
      </q-icon>
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import { useDisciplinesStore } from 'src/stores/disciplines';
import BannerComponent from '../../../../components/BannerComponent.vue';
import { computed } from 'vue';
import LecturerDiscipline from './components/LecturerDiscipline.vue';
import { useRouter } from 'vue-router';

const store = useDisciplinesStore();

store.getDisciplines();

const disciplines = computed(() => store.disciplines);

const router = useRouter();

const openNewDiscipline = () => {
  router.push('new-discipline');
};
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}
</style>
