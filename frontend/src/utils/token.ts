import type { AxiosError, AxiosInstance, AxiosRequestConfig } from 'axios';
import { TokenService } from './tokenService';

export function useTokenInterceptors($api: AxiosInstance) {
  useAccessInterceptor($api);
  useRefreshInterceptor($api);
}

function useAccessInterceptor($api: AxiosInstance) {
  $api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${TokenService.token}`;
    return config;
  });
}

function useRefreshInterceptor($api: AxiosInstance) {
  $api.interceptors.response.use(
    undefined,
    async (error: AuthInterceptorError) => {
      if (error.response?.status === 401) {
        TokenService.token = null;
        TokenService.type = null;
        window.history.pushState({}, '', window.location.origin + '/#/auth');
        window.location.reload();
      }
      throw error;
    }
  );
}

//eslint-disable-next-line
//@ts-ignore
interface AuthInterceptorError extends AxiosError {
  config: AxiosRequestConfig & { _isRetry: boolean | undefined };
}
