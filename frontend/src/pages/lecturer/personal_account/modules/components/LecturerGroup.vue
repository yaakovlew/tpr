<template>
  <q-item class="q-px-none">
    <q-item-section class="row">
      <q-item-label class="row items-center justify-between full-width">
        <div class="group-name">
          {{ group.name }}
          <q-popup-edit
            v-model="name"
            :validate="(val: string) => val !== ''"
            v-slot="scope"
          >
            <q-input
              autofocus
              dense
              v-model="scope.value"
              hint="Название группы"
              :rules="[
                (val: string) => scope.validate(val) || 'Поле не должно быть пустым',
              ]"
            >
              <template v-slot:after>
                <q-btn
                  flat
                  dense
                  color="negative"
                  icon="cancel"
                  @click.stop.prevent="scope.cancel"
                />

                <q-btn
                  flat
                  dense
                  color="positive"
                  icon="check_circle"
                  @click.stop.prevent="changeGroupName(scope.value), scope.set"
                  :disable="
                    scope.validate(scope.value) === false ||
                    scope.initialValue === scope.value
                  "
                />
              </template>
            </q-input>
          </q-popup-edit>
        </div>
        <div class="row g-m">
          <q-btn
            flat
            dense
            color="primary"
            icon="person"
            @click.stop.prevent="getGroupStudents"
          />
          <q-btn
            flat
            dense
            color="primary"
            icon="folder"
            @click.stop.prevent="getGroupDisciplines"
          />
          <q-btn
            flat
            dense
            color="negative"
            icon="cancel"
            @click.stop.prevent="deleteGroupDialog = true"
          />
        </div>
      </q-item-label>
    </q-item-section>
    <q-dialog v-model="deleteGroupDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="dialog-title">
            Вы уверены что хотите удалить группу {{ group.name }}
          </div>
        </q-card-section>
        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn
            flat
            label="Удалить группу"
            v-close-popup
            @click="deleteGroup"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
    <q-dialog v-model="getStudentsDialog">
      <q-card style="min-width: 350px">
        <template v-if="students">
          <q-card-section>
            <div class="dialog-title">Студенты группы {{ group.name }}</div>
          </q-card-section>
          <q-card-section class="q-pt-none">
            <q-list separator>
              <q-item
                v-for="student in students"
                :key="student.user_id"
                v-ripple
                class="q-px-none"
              >
                <q-item-section>
                  {{ `${student.name} ${student.surname}` }}
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
        </template>
        <q-card-section v-else>
          <div class="dialog-title">
            В группе {{ group.name }} нет студентов
          </div>
        </q-card-section>
      </q-card>
    </q-dialog>
    <q-dialog v-model="getDisciplinesDialog">
      <q-card style="min-width: 350px">
        <template v-if="disciplines">
          <q-card-section>
            <div class="dialog-title">Дисциплины группы {{ group.name }}</div>
          </q-card-section>
          <q-card-section class="q-pt-none">
            <q-list separator class="q-pa-none">
              <q-item
                v-for="discipline in disciplines"
                :key="discipline.discipline_id"
                v-ripple
                class="q-px-none"
              >
                <q-item-section>
                  <div class="flex items-center justify-between">
                    <div>
                      {{ `${discipline.name}` }}
                    </div>
                    <q-icon
                      class="cursor-pointer"
                      name="arrow_forward"
                      color="primary"
                      size="24px"
                      @click="
                        goToDiscipline(discipline.discipline_id, group.group_id)
                      "
                    />
                  </div>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
        </template>
        <q-card-section v-else>
          <div class="dialog-title">
            У группы {{ group.name }} нет назначенных дисциплин
          </div>
        </q-card-section>
      </q-card>
    </q-dialog>
  </q-item>
</template>

<script lang="ts" setup>
import { IGroup } from 'src/models/group/group';
import { useGroupsStore } from 'src/stores/groups';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';

const props = defineProps<{
  group: IGroup;
}>();

const store = useGroupsStore();

const students = computed(() => store.groupStudents);
const disciplines = computed(() => store.groupDisciplines);

const name = ref(props.group.name);
const deleteGroupDialog = ref(false);
const getStudentsDialog = ref(false);
const getDisciplinesDialog = ref(false);

const changeGroupName = async (name: string) => {
  await store.changeGroupName(name, Number(props.group.group_id));
  store.getGroups();
};

const deleteGroup = async () => {
  await store.deleteGroup(props.group.group_id);
  store.getGroups();
};

const getGroupStudents = async () => {
  getStudentsDialog.value = true;
  await store.getGroupStudents(props.group.group_id);
};

const getGroupDisciplines = async () => {
  getDisciplinesDialog.value = true;
  await store.getGroupDiscipline(props.group.group_id);
};

const router = useRouter();

const goToDiscipline = (disciplineId: string, groupId: string) => {
  router.push({
    name: 'group-discipline',
    params: { disciplineId: disciplineId, groupId: groupId },
  });
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
