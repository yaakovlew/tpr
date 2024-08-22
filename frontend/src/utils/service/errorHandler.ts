import { IResourceError, type IAxiosError } from 'src/services/resource/error';
import type { IResource } from 'src/services/resource/resource';
import axios, { type AxiosError } from 'axios';

export type IResourceErrorAction = (err: IResourceError) => void;

export function handleAxiosError<T>(
  err: unknown,
  callback?: IResourceErrorAction | null
): IResource<T> {
  if (isAxiosError(err)) {
    const error = IResourceError.build(err);

    // if (err.request.status === 401) {
    //   // const url = new URL('', 'http://soft-computing-mephi.ru.na4u.ru/#/auth');
    //   window.history.pushState({}, '', window.location.origin + '/#/auth');
    //   console.log(window.location.origin);
    // }
    callback?.(error);
    return { error };
  } else throw err;
}

function isAxiosError(err: unknown): err is AxiosError<IAxiosError> {
  return axios.isAxiosError(err);
}
