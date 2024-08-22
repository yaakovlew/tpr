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

          <q-item class="q-px-none">
            <q-item-section>
              <div class="flex row items-center g-m justify-between">
                <q-item-label> {{ discipline.name }} </q-item-label>
                <q-btn
                  icon="person"
                  flat
                  @click="getGroups(discipline.discipline_id)"
                  color="primary"
                />
              </div>
              <!-- <q-item-label caption lines="2">
                {{ discipline.description }}
              </q-item-label> -->
            </q-item-section>
          </q-item>
        </template>
      </q-list>
    </banner-component>
    <banner-component v-else class="text-primary profile-name">
      Нет назначенных дисциплин
    </banner-component>
    <q-dialog v-model="groupDialog">
      <div class="dialog">
        <div class="dialog-title text-primary">Группы</div>
        <q-list class="q-pa-none" separator>
          <q-item
            v-for="group in disciplineGroups"
            :key="group.group_id"
            class="q-pa-none group-name text-primary flex g-m justify-between items-center"
          >
            <div>
              {{ group.name }}
            </div>
            <q-btn
              icon="chevron_right"
              flat
              @click="goToDiscipline(currentDiscipline, group.group_id)"
            />
          </q-item>
        </q-list>
      </div>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import { useDisciplinesStore } from 'src/stores/disciplines';
import BannerComponent from 'src/components/BannerComponent.vue';
import { computed, ref } from 'vue';
import { useGroupsStore } from '../../../../stores/groups';
import { useRouter } from 'vue-router';

const store = useDisciplinesStore();

store.getDisciplines();
const groupsStore = useGroupsStore();
const currentDiscipline = ref('');

const getGroups = async (id: string) => {
  await groupsStore.getSeminarianGroups(id);
  currentDiscipline.value = id;
  groupDialog.value = true;
};

const disciplines = computed(() => store.disciplines);

const disciplineGroups = computed(() => groupsStore.seminarianGroups);

const groupDialog = ref(false);

const router = useRouter();

const goToDiscipline = (disciplineId: string, groupId: string) => {
  router.push({
    name: 'seminarian-group-discipline',
    params: { disciplineId: disciplineId, groupId: groupId },
  });
};
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}

.dialog {
  background: white;
  width: 400px;
  padding: 20px;

  .dialog-title {
    font-size: 20px;
    font-weight: 600;
  }

  .group-name {
    font-size: 18px;
    font-weight: 500;
    min-height: fit-content;
  }
}
</style>
