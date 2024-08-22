import { defineStore } from 'pinia';
import { ILabaratory } from 'src/models/labaratory/labaratory';
import { LabsService } from '../services/labs';
import { ref } from 'vue';

export const useLabsStore = defineStore('labs', () => {
  const labs = ref<ILabaratory.ExternalLabaratory[] | null>(null);
  const sectionLabs = ref<ILabaratory.ExternalLaboratorySection[] | null>(null);

  const addLab = async (lab: ILabaratory.AddLabaratory) => {
    await LabsService.addLabaratory(lab);
  };

  const getLabs = async () => {
    const res = await LabsService.getExternalLabs();
    if (res.data) {
      labs.value = res.data.Ru;
    }
  };

  const getLabsFromSection = async (id: number) => {
    const res = await LabsService.getLabsFromSection(id);
    if (res.data) {
      sectionLabs.value = res.data.labs;
    }
  };

  const getLabsFromSectionReturn = async (id: number) => {
    const res = await LabsService.getLabsFromSection(id);
    if (res.data) {
      return { labs: res.data.labs, id };
    }
  };

  const deleteLab = async (id: number) => {
    await LabsService.deleteLab(id);
  };

  return {
    addLab,
    getLabs,
    labs,
    getLabsFromSection,
    sectionLabs,
    deleteLab,
    getLabsFromSectionReturn,
  };
});
