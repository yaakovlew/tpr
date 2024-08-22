import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { IQuestion } from 'src/models/test/question';
import { TestService } from 'src/services/test';
import { ITest } from 'src/models/test/test';
import { IDigitalLesson } from 'src/models/digitalLesson/digital-lesson';
import { DigitalLessonService } from 'src/services/digitalLesson';

export const useDigitalLessonStore = defineStore('digital-lesson', () => {
  const digitalLessons: Ref<IDigitalLesson.DigitalLesson[] | undefined> = ref();
  const disciplineLessons: Ref<
    IDigitalLesson.AddDigitalLessonToDiscipline[] | undefined
  > = ref();
  const fullDisciplineLessons: Ref<IDigitalLesson.DigitalLesson[] | undefined> =
    ref([]);

  const studentMaterials: Ref<IDigitalLesson.StudentMaterial[] | undefined> =
    ref();

  const getDigitalLessons = async () => {
    const res = await DigitalLessonService.getDigitalLessons();
    if (res.data) {
      digitalLessons.value = res.data.ru;
    }
  };

  const addDigitalLesson = async (data: IDigitalLesson.AddDigitalLesson) => {
    await DigitalLessonService.addDigitalLesson(data);
  };

  const addDigitalLessonToDiscipline = async (
    data: IDigitalLesson.AddDigitalLessonToDiscipline
  ) => {
    await DigitalLessonService.addDigitalLessonToDiscipline(data);
  };

  const changeDigitalLesson = async (
    data: IDigitalLesson.ChangeDigitalLesson
  ) => {
    await DigitalLessonService.changeDigitalLesson(data);
  };

  const getDisciplineLessons = async (disciplineId: number) => {
    const res = await DigitalLessonService.getDisciplineLessons(disciplineId);
    if (res.data) {
      disciplineLessons.value = res.data.digital_guides;
    }
    await getDigitalLessons();
    fullDisciplineLessons.value = [];
    disciplineLessons.value?.forEach((lesson) => {
      digitalLessons.value?.forEach((l) => {
        if (lesson.digital_material_id === l.study_guide_id) {
          fullDisciplineLessons.value?.push(l);
        }
      });
    });
  };

  const removeDigitalLessonFromDiscipline = async (
    data: IDigitalLesson.AddDigitalLessonToDiscipline
  ) => {
    await DigitalLessonService.removeDigitalLessonFromDiscipline(data);
  };

  const removeDigitalLesson = async (id: number) => {
    await DigitalLessonService.removeDigitalLesson(id);
  };

  const getStudentMaterials = async (id: number) => {
    const res = await DigitalLessonService.getStudentMaterials(id);
    if (res.data) {
      studentMaterials.value = res.data.en;
    }
  };

  const uploadDigitalLesson = async (
    data: IDigitalLesson.AddDigitalMaterial
  ) => {
    await DigitalLessonService.uploadDigitalLesson(data);
  };

  const downloadDigitalLesson = async (id: number) => {
    const res = await DigitalLessonService.downloadDigitalLesson(id);
    if (res.data) {
      return res.data;
    }
  };

  const getFilesId = async (id: number) => {
    const res = await DigitalLessonService.getFilesId(id);
    if (res.data) {
      return res.data;
    }
  };

  return {
    getDigitalLessons,
    addDigitalLesson,
    addDigitalLessonToDiscipline,
    changeDigitalLesson,
    digitalLessons,
    getDisciplineLessons,
    fullDisciplineLessons,
    removeDigitalLessonFromDiscipline,
    removeDigitalLesson,
    getStudentMaterials,
    studentMaterials,
    uploadDigitalLesson,
    downloadDigitalLesson,
    getFilesId,
  };
});
