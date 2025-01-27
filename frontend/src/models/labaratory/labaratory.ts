export namespace ILabaratory {
  export interface AddLabaratory {
    name: string;
    name_en: string;
    task_description: string;
    task_description_en: string;
    minutes_duration: number;
    linc: string;
    default_mark: number;
    day_fine: number;
    token: string;
  }

  export interface EditLabaratory {
    name: string;
    description: string;
    name_en: string;
    description_en: string;
    laboratory_id: string;
    // token: string | null;
    // linc: string  | null;
  }

  export interface Labaratory {
    laboratory_id: number;
    name: string;
    task_description: string;
    minutes_duration: number;
    linc: string;
    default_mark: number;
    day_fine: number;
  }

  export interface LabWithClosedDate {
    laboratory_id: number;
    name: string;
    task_description: string;
    default_mark: number;
    closed_date: number;
  }

  export interface ExternalLaboratorySection {
    laboratory_id: number;
    external_laboratory_id: number;
    name: string;
    task_description: string;
    link: string;
    default_mark: number;
  }

  export interface ExternalLabaratory {
    external_laboratory_id: number;
    link: string;
    name: string;
    task_description: string;
  }

  export interface GetLabs {
    Ru: Labaratory[];
    En: Labaratory[];
  }

  export interface GetLabsWithClosedDate {
    ru: LabWithClosedDate[];
    en: LabWithClosedDate[];
  }

  export interface GetLabsFromSection {
    labs: ExternalLaboratorySection[];
  }

  export interface GetExternalLabs {
    Ru: ExternalLabaratory[];
    En: ExternalLabaratory[];
  }

  export interface StudentOpenLab {
    student_id: number;
    name: string;
    surname: string;
    group_name: string;
    closed_date: number
  }

  export interface GetStudentsOpenLab {
    students: StudentOpenLab[];
  }

  export interface OpenLab {
    date: number;
    laboratory_id: number;
    user_id: number;
  }

  export interface CloseLab {
    laboratory_id: number;
    user_id: number;
  }

}
