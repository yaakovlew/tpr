<template>
  <div class="flex column g-m">
    <banner-component class="text-primary profile-name">
      Группы
    </banner-component>
    <template v-if="groups">
      <banner-component>
        <q-list class="q-pa-none">
          <template v-for="(group, index) in groups" :key="group.name">
            <q-separator
              v-if="index !== 0"
              spaced
              inset
              class="q-px-none q-mx-none"
            />
            <lecturer-group :group="group" />
          </template>
        </q-list>
      </banner-component>
    </template>
    <template v-else>
      <banner-component class="text-primary profile-name">
        Нет групп
      </banner-component>
    </template>
    <banner-component class="text-primary profile-name" @click="openNewGroup">
      Добавить группу
      <q-icon name="add" color="primary" size="50px" class="cursor-pointer">
        <q-tooltip> Добавить группу </q-tooltip>
      </q-icon>
    </banner-component>
    <q-dialog v-model="newGroup" persistent @hide="closeNewGroup">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">Имя группы</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input
            dense
            v-model="newGroupName"
            autofocus
            @keyup.enter="addGroup"
          />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn flat label="Добавить группу" @click="addGroup" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import BannerComponent from '../../../../components/BannerComponent.vue';
import LecturerGroup from './components/LecturerGroup.vue';
import { computed, ref } from 'vue';
import { useGroupsStore } from 'src/stores/groups';
import { Notify } from 'quasar';

const store = useGroupsStore();

store.getGroups();

const groups = computed(() => store.groups);

const newGroup = ref(false);
const newGroupName = ref('');

const addGroup = () => {
  if (newGroupName.value) {
    store.addNewGroup(newGroupName.value);
    closeNewGroup();
  } else {
    Notify.create({
      type: 'negative',
      message: 'Название группы не должно быть пустым',
    });
  }
};

const openNewGroup = () => {
  newGroup.value = true;
};

const closeNewGroup = () => {
  newGroup.value = false;
  newGroupName.value = '';
  store.getGroups();
};
</script>

<style lang="scss" scoped>
.profile-name {
  font-size: 24px;
  font-weight: 600;
  flex-grow: 100;
}
</style>
