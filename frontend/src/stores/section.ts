import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { ISection } from '../models/section/section';
import { SectionService } from '../services/section';

export const useSectionStore = defineStore('sections', () => {
  const sections: Ref<ISection[] | null> = ref(null);

  const getSections = async (id: number) => {
    const res = await SectionService.fetch(id);
    if (res.data) {
      sections.value = res.data.ru;
    }
  };

  const getSectionsReturn = async (id: number) => {
    const res = await SectionService.fetch(id);
    if (res.data) {
      return res.data.ru;
    }
  };

  const createSection = async (data: ISection.CreateSection) => {
    await SectionService.createSection(data);
  };

  const changeSection = async (data: ISection.Section) => {
    await SectionService.changeSection(data);
  };

  const addLabToSection = async (data: ISection.AddLabToSection) => {
    await SectionService.addLabToSection(data);
  };

  const addTestToSection = async (data: ISection.AddTestToSection) => {
    await SectionService.addTestToSection(data);
  };

  const deleteTestFromSection = async (data: ISection.AddTestToSection) => {
    await SectionService.deleteTestFromSection(data);
  };

  const deleteLabFromSection = async (data: ISection.DeleteLabFromSection) => {
    await SectionService.deleteLabFromSection(data);
  };

  const getStudentDisciplineSections = async (id: number) => {
    const res = await SectionService.getStudentDisciplineSections(id);
    if (res.data) {
      sections.value = res.data.ru;
    }
  };

  const getSectionsSeminarian = async (id: number) => {
    const res = await SectionService.fetchSeminarian(id);
    if (res.data) {
      sections.value = res.data.ru;
    }
  };

  const deleteSection = async (id: number) => {
    await SectionService.deleteSection(id);
  };

  return {
    sections,
    getSections,
    createSection,
    changeSection,
    addLabToSection,
    addTestToSection,
    deleteTestFromSection,
    deleteLabFromSection,
    getStudentDisciplineSections,
    getSectionsSeminarian,
    getSectionsReturn,
    deleteSection,
  };
});
