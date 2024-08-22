<template>
  <q-item class="q-px-none">
    <q-item-section class="row">
      <q-item-label class="row items-center justify-between full-width">
        <div class="row g-m">
          <q-btn
            flat
            dense
            color="primary"
            icon="folder"
            @click.stop.prevent="openChangeGroup"
          />
          <q-btn
            flat
            dense
            color="primary"
            icon="key"
            @click.stop.prevent="openChangePassword"
          />
          <q-btn
            flat
            dense
            color="red"
            icon="delete"
            @click.stop.prevent="deleteDialog = true"
          />
        </div>
      </q-item-label>
    </q-item-section>
    <q-dialog v-model="changeGroupDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="dialog-title">
            Выберите новую группу для {{ `${student.surname} ${student.name}` }}
          </div>
          <q-select
            v-model="newGroup"
            :options="groups"
            :option-value="
                (opt : IGroup) => (Object(opt) === opt && 'name' in opt ? opt.name : null)
              "
            :option-label="
                (opt : IGroup) =>
                  Object(opt) === opt && 'name' in opt ? opt.name : ''
              "
            map-options
          />
        </q-card-section>
        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn
            flat
            label="Изменить группу"
            v-close-popup
            @click="changeGroup"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
    <q-dialog v-model="changePasswordDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="dialog-title">
            Изменение пароля для студента
            {{ `${student.surname} ${student.name}` }}
          </div>
        </q-card-section>
        <q-card-section>
          <q-input v-model="newPassword" />
        </q-card-section>
        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn
            flat
            label="Изменить пароль"
            v-close-popup
            @click="changePassword"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
    <q-dialog v-model="deleteDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="dialog-title">
            Вы уверены что хотите удалить студента
            {{ `${student.surname} ${student.name}` }}
          </div>
        </q-card-section>
        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn flat label="Удалить" v-close-popup @click="deleteUser" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-item>
</template>

<script lang="ts" setup>
import { IGroup } from 'src/models/group/group';
import { IStudent } from 'src/models/student/student';
import { useGroupsStore } from 'src/stores/groups';
import { useStudentsStore } from 'src/stores/students';
import { Ref, computed, ref } from 'vue';

const props = defineProps<{
  student: IStudent;
}>();

const store = useStudentsStore();
const groupsStore = useGroupsStore();

const groups = computed(() => groupsStore.commonGroups);

const changePasswordDialog = ref(false);
const changeGroupDialog = ref(false);
const deleteDialog = ref(false);

const newGroup: Ref<IGroup | null> = ref(null);
const newPassword: Ref<string> = ref('');

const openChangeGroup = async () => {
  await groupsStore.getCommonGroups();
  changeGroupDialog.value = true;
  newGroup.value = null;
};

const changeGroup = async () => {
  if (newGroup.value) {
    await store.changeStudentData(
      props.student.student_id,
      newGroup.value.group_id
    );
  }
  await store.getStudents();
  changeGroupDialog.value = false;
};

const openChangePassword = () => {
  newPassword.value = '';
  changePasswordDialog.value = true;
};

const changePassword = async () => {
  if (newPassword.value) {
    await store.changeStudentPassword(
      props.student.student_id,
      newPassword.value
    );
  }
  changePasswordDialog.value = false;
};

const deleteUser = async () => {
  await store.deleteStudent(props.student.student_id);
  store.getStudents();
};
</script>

<style lang="scss" scoped>
.group-name {
  font-size: 18px;
}
.dialog-title {
  font-weight: 500;
  font-size: 18px;
}
</style>
