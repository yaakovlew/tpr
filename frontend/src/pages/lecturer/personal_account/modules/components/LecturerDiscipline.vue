<template>
  <div
    class="discipline-name text-primary flex items-center justify-between no-wrap"
  >
    <div>
      {{ discipline.name }}
      <q-popup-edit
        v-model="disciplineName"
        title="Измените название дисциплины"
        label-set="Сохранить"
        label-cancel="Отменить"
        v-slot="scope"
      >
        <q-input
          v-model="scope.value"
          dense
          autofocus
          counter
          @keyup.enter="scope.set"
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
              @click.stop.prevent="changeDisciplineName(scope.value), scope.set"
              :disable="
                scope.validate(scope.value) === false ||
                scope.initialValue === scope.value
              "
            />
          </template>
        </q-input>
      </q-popup-edit>
    </div>
    <div class="action-buttons">
      <q-btn
        flat
        dense
        icon="edit"
        class="cursor-pointer"
        @click="goToEditDiscipline"
      />
      <q-btn
        flat
        dense
        color="red"
        icon="delete"
        @click.stop.prevent="deleteDialog = true"
      />
    </div>
    <q-dialog v-model="deleteDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="dialog-title">
            Вы уверены что хотите удалить дисциплину
            {{ `${discipline.name}` }}
          </div>
        </q-card-section>
        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Закрыть" v-close-popup />
          <q-btn flat label="Удалить" v-close-popup @click="deleteDiscipline" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script lang="ts" setup>
import { IDiscipline } from 'src/models/discipline/discipline';
import { useDisciplinesStore } from 'src/stores/disciplines';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const props = defineProps<{
  discipline: IDiscipline;
}>();

const disciplineName = ref(props.discipline.name);

const store = useDisciplinesStore();

const router = useRouter();

const goToEditDiscipline = () => {
  router.push({
    name: 'edit-discipline',
    params: { disciplineId: props.discipline.discipline_id },
  });
};

const changeDisciplineName = async (name: string) => {
  await store.changeDisciplineName({
    name: name,
    name_en: name,
    discipline_id: Number(props.discipline.discipline_id),
  });
  await store.getDisciplines();
};

const deleteDialog = ref(false);

const deleteDiscipline = async () => {
  await store.deleteDiscipline(Number(props.discipline.discipline_id));
  await store.getDisciplines();
};
</script>

<style lang="scss" scoped>
.discipline-name {
  font-size: 20px;
  font-weight: 600;
}

@media screen and (max-width: 600px) {
  .action-buttons {
    width: 50px;
  }
}
</style>
