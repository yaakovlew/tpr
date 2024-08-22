<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Дисциплины
    </banner-component>
    <banner-component v-if="disciplines">
      <q-list class="q-pa-none" separator>
        <template v-for="discipline in disciplines" :key="discipline.name">
          <q-item class="q-px-none flex justify-between items-center">
            <q-item-label class="text-primary name">
              {{ discipline.name }}
            </q-item-label>
            <q-icon
              name="info"
              color="primary"
              size="24px"
              @click="goToDiscipline(Number(discipline.discipline_id))"
              class="cursor-pointer"
            />
            <!-- <q-item-label caption lines="2">
                {{ discipline.description }}
              </q-item-label> -->
          </q-item>
        </template>
      </q-list>
    </banner-component>
    <banner-component v-else class="text-primary profile-name">
      Нет назначенных дисциплин
    </banner-component>
  </div>
</template>

<script lang="ts" setup>
import { useDisciplinesStore } from 'src/stores/disciplines';
import BannerComponent from '../../../components/BannerComponent.vue';
import { computed } from 'vue';
import { useRouter } from 'vue-router';

const store = useDisciplinesStore();

store.getDisciplines();

const disciplines = computed(() => store.disciplines);

const router = useRouter();

const goToDiscipline = (id: number) => {
  router.push({ name: 'student-discipline', params: { disciplineId: id } });
};
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}

.name {
  font-size: 20px;
  font-weight: 500;
}
</style>
