import { boot } from 'quasar/wrappers';
import type { AxiosInstance, AxiosRequestConfig } from 'axios';
import axios from 'axios';
import { useTokenInterceptors } from 'src/utils/token';
// import { useDatesSerializer } from 'src/utils/dates/dates';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
  }
}

// export const baseURL = process.env.BASE_URL
// export const apiURL = baseURL + process.env.API_PATH
// export const apiPublicURL = apiURL + process.env.PUBLIC_PATH
// export const apiUserURL = apiURL + process.env.USER_PATH
// export const apiAdminURL = apiURL + process.env.ADMIN_PATH
// export const staticURL = apiPublicURL + process.env.STATIC_PATH
// export const staticUserURL = apiUserURL + process.env.STATIC_PATH
// export const staticAdminURL = apiAdminURL + process.env.STATIC_PATH

// export const baseURL = 'http://localhost:8000';
export const baseURL = process.env.BASE_URL;
export const apiURL = baseURL + '/api';
export const apiPublicURL = baseURL;
export const apiStudentURL = apiURL + '/student';
export const apiLecturerURL = apiURL + '/lecturer';
export const apiSeminarianURL = apiURL + '/seminarian';
export const staticURL = apiPublicURL + process.env.STATIC_PATH;
export const staticUserURL = apiStudentURL + process.env.STATIC_PATH;
export const staticAdminURL = apiLecturerURL + process.env.STATIC_PATH;

export const axiosCommonConfig: AxiosRequestConfig = {
  withCredentials: true,
};
export const axiosPublicConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiPublicURL,
};
export const axiosStudentConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiStudentURL,
};
export const axiosLecturerConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiLecturerURL,
};
export const axiosSeminarianConfig: AxiosRequestConfig = {
  ...axiosCommonConfig,
  baseURL: apiSeminarianURL,
};

export const $apiPublic = axios.create(axiosPublicConfig);
// useDatesSerializer($apiPublic);

export const $apiStudent = axios.create(axiosStudentConfig); 
// useDatesSerializer($apiUser);
useTokenInterceptors($apiStudent);

export const $apiLecturer = axios.create(axiosLecturerConfig);
// useDatesSerializer($apiAdmin);
useTokenInterceptors($apiLecturer);

export const $apiSemianrian = axios.create(axiosSeminarianConfig);
// useDatesSerializer($apiAdmin);
useTokenInterceptors($apiSemianrian);

export const $fdHeaders = { 'Content-Type': 'multipart/form-data' };

export default boot(({ app }) => {
  app.config.globalProperties.$axios = axios;
});
