export namespace IDigitalLesson {
  export interface DigitalLesson {
    description: string;
    name: string;
    study_guide_id: number;
  }

  export interface ChangeDigitalLesson {
    digital_guide_id: number;
    description: string;
    description_en: string;
    name: string;
    name_en: string;
  }

  export interface StudentMaterial {
    description: string;
    digital_material_id: number;
    name: string;
  }

  export interface GetDigitalLessons {
    en: DigitalLesson[];
    ru: DigitalLesson[];
  }

  export interface GetFileId {
    file_id: number;
  }

  export interface GetFileIds {
    files: GetFileId[];
  }

  export interface GetDigitalLessonsStudent {
    en: StudentMaterial[];
    ru: StudentMaterial[];
  }

  export interface AddDigitalLesson {
    description: string;
    description_en: string;
    name: string;
    name_en: string;
  }

  export interface AddDigitalMaterial {
    file: FormData;
    id: number;
  }

  export interface AddDigitalLessonToDiscipline {
    discipline_id: number;
    digital_material_id: number;
  }
}
