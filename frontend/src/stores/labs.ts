import { defineStore } from 'pinia';
import { ILabaratory } from 'src/models/labaratory/labaratory';
import { LabsService } from '../services/labs';
import { ref } from 'vue';

export const useLabsStore = defineStore('labs', () => {
  const labs = ref<ILabaratory.ExternalLabaratory[] | null>(null);
  const sectionLabs = ref<ILabaratory.ExternalLaboratorySection[] | null>(null);
  const studentsOpenedLab = ref<ILabaratory.StudentOpenLab[]>([]);
  const allSectionsLabs = ref<Record<number, ILabaratory.GetLabsFromSection>>({});

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
    console.log('res: ', res)
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

  const editLab = async (data: ILabaratory.EditLabaratory) => {
    await LabsService.editLab(data);
  };

  const getStudentsOpenLabs = async (id: number) => {
    const res = await LabsService.getStudentsOpenLab(id);
    if (res.data) {
      studentsOpenedLab.value = res.data.students;
    }
  };

  const openLab = async (data: ILabaratory.OpenLab) => {
    await LabsService.openLab(data);
  };

  const closeLab = async (data: ILabaratory.CloseLab) => {
    await LabsService.closeLab(data);
  };

  const getSectionLabs = async (id: number) => {
    const res = await LabsService.getLabsFromSection(id);
    console.log('got res http: ', res.data)
    if (res.data) {
      sectionLabs.value = res.data.labs;
    }
    return { data: res.data, id };
  };

  const getAllSectionsLabs = async (id: number[]) => {
    console.log('in getAllSectionsLabs')
    const promises: Promise<{data: ILabaratory.GetLabsFromSection | undefined; id: number;}>[] = [];
    id.forEach((id) => {
      promises.push(getSectionLabs(id));
    });
    const res = await Promise.all(promises);
    allSectionsLabs.value = {};
    if (res) {
      console.log('got result: ', res)
      res.forEach((section) => {
        if (section.data){
          allSectionsLabs.value[section.id] = section.data;
        }
      });
      console.log('res res: ', res)
    }
  };



  return {
    addLab,
    getLabs,
    labs,
    getLabsFromSection,
    sectionLabs,
    deleteLab,
    getLabsFromSectionReturn,
    editLab,
    getStudentsOpenLabs,
    studentsOpenedLab,
    openLab,
    closeLab,
    getAllSectionsLabs,
    allSectionsLabs,
  };
});
