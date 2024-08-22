export type ISection = ISection.Section;

export namespace ISection {
  export interface Section {
    name: string;
    section_id: number;
  }

  export interface GetSections {
    en: ISection[];
    ru: ISection[];
  }

  export interface CreateSection {
    name: string;
    name_en: string;
    discipline_id: number;
  }

  export interface AddLabToSection {
    external_lab_id: number;
    section_id: number;
    default_mark: number;
  }

  export interface AddTestToSection {
    test_id: number;
    section_id: number;
  }

  export interface DeleteLabFromSection {
    laboratory_id: number;
    section_id: number;
  }
}
