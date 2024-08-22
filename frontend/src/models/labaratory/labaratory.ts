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

  export interface Labaratory {
    laboratory_id: number;
    name: string;
    task_description: string;
    minutes_duration: number;
    linc: string;
    default_mark: number;
    day_fine: number;
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

  export interface GetLabsFromSection {
    labs: ExternalLaboratorySection[];
  }

  export interface GetExternalLabs {
    Ru: ExternalLabaratory[];
    En: ExternalLabaratory[];
  }
}
