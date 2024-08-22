export namespace ITheme {
  export interface Theme {
    theme_id: number;
    name: string;
    weight: number;
  }

  export interface ThemeTest {
    theme_id: number;
    name: string;
    weight: number;
    count: number;
  }

  export interface ThemeTestWithTotalCount {
    theme_id: number;
    name: string;
    weight: number;
    count: number;
    total_count: number | undefined;
  }

  export interface CreateTheme {
    name: string;
    weight: number;
  }

  export interface ChangeThemeName {
    name: string;
    theme_id: number;
    weight: number;
  }

  export interface ChangeThemeWeight {
    weight: number;
    theme_id: number;
  }

  export interface GetThemes {
    themes: ITheme.Theme[];
  }
  export interface GetThemesTest {
    themes: ITheme.ThemeTest[];
  }
}
