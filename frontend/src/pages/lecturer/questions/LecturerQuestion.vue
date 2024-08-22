<template>
  <q-tr>
    <q-td key="question">
      <div class="text">
        {{ question.question }}
      </div>
      <q-popup-edit v-model="questionName" v-slot="scope">
        <q-input
          autofocus
          dense
          v-model="scope.value"
          :model-value="scope.value"
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
              @click.stop.prevent="
                changeQuestion(scope.value);
                scope.set();
              "
              :disable="scope.value === ''"
            />
          </template>
        </q-input>
      </q-popup-edit>
    </q-td>
    <q-td key="variable">
      {{
        question.is_variable === 1
          ? 'Один вариант'
          : question.is_variable === 0
          ? 'Свободный вариант'
          : 'Несколько вариантов'
      }}
    </q-td>
    <q-td key="theme">
      {{ themes?.map((theme) => theme.name).join(', ') }}
    </q-td>
    <q-td key="answers" class="text">
      <q-list class="q-pa-none" separator>
        <q-item
          v-for="answer in answersToThisQuestion"
          :key="answer.answer_id"
          class="flex items-center justify-between q-px-none q-py-none answer"
          :class="answer.is_right ? 'text-positive' : 'text-negative'"
        >
          {{ answer.answer }}
        </q-item>
      </q-list>
    </q-td>
    <q-td key="actions">
      <div class="flex g-m row" style="width: 55px">
        <q-icon
          name="list"
          class="cursor-pointer"
          @click="getAnswers"
          color="primary"
          size="24px"
        />
        <q-icon
          name="delete"
          color="negative"
          size="24px"
          class="cursor-pointer"
          @click="deleteModal = true"
        />
      </div>
    </q-td>
    <q-dialog v-model="answersModal">
      <div class="modal q-pa-md text-primary flex column">
        <div class="title">Темы в которых есть вопрос</div>
        <q-list class="q-pa-none" separator>
          <q-item
            class="flex items-center justify-between q-px-none answer"
            v-for="theme in themes"
            :key="theme.theme_id"
          >
            {{ theme.name }}
          </q-item>
        </q-list>
        <div class="title" style="margin-top: 20px">Ответы</div>
        <q-list class="q-pa-none" separator>
          <q-item
            v-for="answer in answers"
            :key="answer.answer_id"
            class="flex items-center justify-between q-px-none answer"
            :class="answer.is_right ? 'text-positive' : 'text-negative'"
          >
            {{ answer.answer }}
            <q-popup-edit
              v-model="answer.answer"
              v-slot="scope"
              @show="openPopup(answer)"
            >
              <div class="text-primary is-right">
                Правильный
                <q-checkbox v-model="changeIsRight" />
              </div>
              <q-input
                autofocus
                dense
                v-model="scope.value"
                :model-value="scope.value"
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
                    @click.stop.prevent="
                      changeAnswer(answer.answer_id, scope.value);
                      scope.set();
                    "
                    :disable="scope.value === ''"
                  />
                </template>
              </q-input>
            </q-popup-edit>
            <q-icon
              color="negative"
              class="cursor-pointer"
              size="20px"
              name="delete"
              @click="deleteAnswer(answer.answer_id)"
            />
          </q-item>
        </q-list>
        <q-btn
          v-if="!addNewForm"
          label="Добавить ответ"
          flat
          class="align-right self-end"
          @click="addNewForm = true"
        />
        <div v-else class="flex column">
          <q-input autofocus label="Ответ" v-model="newAnswer" />
          <div class="text-primary is-right">
            Правильный
            <q-checkbox v-model="isNewRight" />
          </div>
          <div class="flex row g-m items-center justify-end">
            <q-btn
              label="Добавить"
              flat
              class="align-right self-end"
              @click="addNewAnswer"
              :disable="!newAnswer"
            />
            <q-btn
              label="Отмена"
              flat
              class="align-right self-end"
              @click="closeNewForm"
            />
          </div>
        </div>
      </div>
    </q-dialog>
    <q-dialog v-model="deleteModal">
      <div class="delete-modal flex g-m text-primary q-pa-md">
        <div>
          Вы уверены что хотите удалить вопрос: "{{ question.question }}"?
        </div>
        <div class="flex g-m justify-between self-end">
          <q-btn flat label="Отмена" @click="deleteModal = false" />
          <q-btn flat label="Удалить" @click="deleteQuestion" />
        </div>
      </div>
    </q-dialog>
  </q-tr>
</template>

<script lang="ts" setup>
import { ITheme } from 'src/models/test/theme';
import { IQuestion } from '../../../models/test/question';
import { useTestsStore } from '../../../stores/test';
import { computed, ref, watch, Ref, onMounted } from 'vue';

const props = defineProps<{
  question: IQuestion;
}>();

const variableVariants = [
  {
    value: 0,
    label: 'Один вариант',
  },
  {
    value: 1,
    label: 'Множественный вариант',
  },
  {
    value: 2,
    label: 'Свободный выбор ответа',
  },
];

const questionName = ref(props.question.question);

const changeQuestion = async (question: string) => {
  await store.changeQuestion({
    name: question,
    question_id: props.question.question_id,
  });
  await store.getAllQuestions();
};

const store = useTestsStore();

const themes: Ref<ITheme.Theme[] | null> = ref(null);

const getThemes = async () => {
  const res = await store.getThemesByQuestion(props.question.question_id);
  if (res) {
    themes.value = res.themes;
  }
};

const answersToThisQuestion: Ref<IQuestion.Answer[] | null> = ref(null);

onMounted(async () => {
  await getThemes();
  answersToThisQuestion.value =
    (await store.getAnswers(props.question.question_id)) ?? null;
});

const answers = computed(() => store.answers);

const answersModal = ref(false);

const newAnswer = ref('');
const isNewRight = ref(false);

const addNewForm = ref(false);

const changeIsRight = ref(false);

const openPopup = (answer: IQuestion.Answer) => {
  changeIsRight.value = answer.is_right;
};

const closeNewForm = () => {
  addNewForm.value = false;
  newAnswer.value = '';
  isNewRight.value = false;
};

const deleteModal = ref(false);

const deleteQuestion = async () => {
  await store.deleteQuestion(props.question.question_id);
  await store.getAllQuestions();
  deleteModal.value = false;
};

const addNewAnswer = async () => {
  await store.addAnswerToQuestion({
    is_right: isNewRight.value,
    name: newAnswer.value,
    name_en: newAnswer.value,
    question_id: props.question.question_id,
  });
  await store.getAnswers(props.question.question_id);
  closeNewForm();
};

const deleteAnswer = async (id: number) => {
  await store.deleteAnswer(id);
  await store.getAnswers(props.question.question_id);
};

const changeAnswer = async (id: number, newName: string) => {
  await store.changeAnswerName({
    name: newName,
    answer_id: id,
    is_right: changeIsRight.value ? '1' : '0',
  });

  await store.getAnswers(props.question.question_id);
};

watch(answersModal, () => {
  if (!answersModal.value) {
    closeNewForm();
  }
});

const getAnswers = async () => {
  await getThemes();
  await store.getAnswers(props.question.question_id);
  answersModal.value = true;
};
</script>

<style lang="scss" scoped>
.question {
  font-size: 18px;
  font-weight: 500;
  width: 250px;
}

.modal {
  background-color: white;
  min-width: 500px;
}

.title {
  font-size: 18px;
  font-weight: 600;
}

.answer {
  font-weight: 400;
}

.is-right {
  font-size: 16px;
}

.delete-modal {
  font-size: 20px;
  font-weight: 500;
  width: 400px;
  background-color: white;
}

.text {
  white-space: break-spaces;
  width: 300px;
}
</style>
