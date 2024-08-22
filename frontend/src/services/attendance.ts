import { $apiLecturer, $apiSemianrian } from 'src/boot/axios';
import {
  IAttendance,
  ILesson,
  ISeminar,
} from 'src/models/attendance/attendance';
import { IVisiting } from 'src/models/attendance/visiting';
import { useServiceAction } from 'src/utils/service/action';

export const AttendanceService = {
  fetch: useServiceAction((id: string) =>
    $apiLecturer.get<IAttendance.GetLessons>(`/attendance/lesson/${id}`)
  ),
  createLesson: useServiceAction((data: IAttendance.CreateLesson) =>
    $apiLecturer.post<ILesson[]>('/attendance/lesson', data)
  ),
  changeLessonName: useServiceAction((data: ILesson) =>
    $apiLecturer.put<ILesson[]>('/attendance/lesson', data)
  ),
  getLessonVisiting: useServiceAction((data: IVisiting.FetchData) =>
    $apiLecturer.get<IVisiting.GetLessonsVisiting>(
      `/attendance/lesson/visiting?group_id=${data.group_id}&lesson_id=${data.lesson_id}`
    )
  ),
  changeLessonVisiting: useServiceAction(
    (data: IVisiting.FetchStudentVisiting) =>
      $apiLecturer.put<ILesson[]>('/attendance/lesson/visiting', data)
  ),
  addLessonVisiting: useServiceAction((data: IVisiting.FetchStudentVisiting) =>
    $apiLecturer.post<ILesson[]>('/attendance/lesson/visiting', data)
  ),
  deleteLesson: useServiceAction((id: string) =>
    $apiLecturer.delete<ILesson[]>(`/attendance/lesson/${id}`)
  ),
  fetchSeminars: useServiceAction((data: IAttendance.FecthSeminars) =>
    $apiLecturer.get<IAttendance.GetSeminars>(
      `/attendance/seminar?group_id=${data.group_id}&discipline_id=${data.discipline_id}`
    )
  ),
  deleteSeminar: useServiceAction((id: number) =>
    $apiLecturer.delete<ISeminar[]>(`/attendance/seminar/${id}`)
  ),
  createSeminar: useServiceAction((data: IAttendance.CreateSeminar) =>
    $apiLecturer.post<ISeminar[]>('/attendance/seminar', data)
  ),
  changeSeminarName: useServiceAction((data: IAttendance.ChangeSeminarName) =>
    $apiLecturer.put<ISeminar[]>(
      `/attendance/seminar?seminar_id=${data.seminar_id}&name=${data.name}&date=${data.date}`,
      data
    )
  ),
  changeSeminarDate: useServiceAction((data: IAttendance.ChangeSeminarDate) => {
    return $apiLecturer.put<ISeminar[]>('/attendance/seminar/date', data);
  }),
  getSeminarVisiting: useServiceAction((id: number) =>
    $apiLecturer.get<IVisiting.GetSeminarsVisiting>(
      `/attendance/seminar/visiting/${id}`
    )
  ),
  addSeminarVisiting: useServiceAction(
    (data: IVisiting.FetchStudentSeminarVisiting) =>
      $apiLecturer.post<ISeminar[]>('/attendance/seminar/visiting', data)
  ),
  changeSeminarVisiting: useServiceAction(
    (data: IVisiting.FetchStudentSeminarVisiting) =>
      $apiLecturer.put<ISeminar[]>('/attendance/seminar/visiting', data)
  ),
  getLessonDate: useServiceAction((data: IAttendance.GetLessonDate) =>
    $apiLecturer.get<IAttendance.LessonDate>(
      `/attendance/lesson/date?lesson_id=${data.lesson_id}&group_id=${data.group_id}`
    )
  ),
  changeLessonDate: useServiceAction((data: IAttendance.ChangeLessonDate) =>
    $apiLecturer.put<IAttendance.ChangeLessonDate>(
      '/attendance/lesson/date',
      data
    )
  ),
  addLessonDate: useServiceAction((data: IAttendance.AddLessonDate) =>
    $apiLecturer.post<IAttendance.AddLessonDate>(
      '/attendance/lesson/date',
      data
    )
  ),
  deleteLessonDate: useServiceAction((data: IAttendance.DeleteLessonDate) =>
    $apiLecturer.delete<IAttendance.AddLessonDate>(
      `/attendance/lesson/date?lesson_id=${data.lesson_id}&group_id=${data.group_id}`
    )
  ),
  getDisciplineLessons: useServiceAction((disciplineId: number) =>
    $apiLecturer.get<IAttendance.GetLessons>(
      `/attendance/lesson/${disciplineId}`
    )
  ),
  getGroupLessons: useServiceAction((data: IAttendance.FecthSeminars) =>
    $apiLecturer.get<IAttendance.GetLessonsDate>(
      `/attendance/lesson?group_id=${data.group_id}&discipline_id=${data.discipline_id}`
    )
  ),
  //TODO: в тг фотка висит
  getLessonVisitingSeminarian: useServiceAction((data: IVisiting.FetchData) =>
    $apiSemianrian.get<IVisiting.GetLessonsVisiting>(
      `/attendance/lesson/visiting?group_id=${data.group_id}&lesson_id=${data.lesson_id}`
    )
  ),
  changeLessonVisitingSeminarian: useServiceAction(
    (data: IVisiting.FetchStudentVisiting) =>
      $apiSemianrian.put<ILesson[]>('/attendance/lesson/visiting', data)
  ),
  addLessonVisitingSeminarian: useServiceAction(
    (data: IVisiting.FetchStudentVisiting) =>
      $apiSemianrian.post<ILesson[]>('/attendance/lesson/visiting', data)
  ),
  fetchSeminarsSeminarian: useServiceAction((data: IAttendance.FecthSeminars) =>
    $apiSemianrian.get<IAttendance.GetSeminars>(
      `/attendance/seminar?group_id=${data.group_id}&discipline_id=${data.discipline_id}`
    )
  ),
  deleteSeminarSeminarian: useServiceAction((id: number) =>
    $apiSemianrian.delete<ISeminar[]>(`/attendance/seminar/${id}`)
  ),
  createSeminarSeminarian: useServiceAction((data: IAttendance.CreateSeminar) =>
    $apiSemianrian.post<ISeminar[]>('/attendance/seminar', data)
  ),
  changeSeminarNameSeminarian: useServiceAction(
    (data: IAttendance.ChangeSeminarName) =>
      $apiSemianrian.put<ISeminar[]>(
        `/attendance/seminar?seminar_id=${data.seminar_id}&name=${data.name}&date=${data.date}`,
        data
      )
  ),
  changeSeminarDateSeminarian: useServiceAction(
    (data: IAttendance.ChangeSeminarDate) => {
      return $apiSemianrian.put<ISeminar[]>('/attendance/seminar/date', data);
    }
  ),
  getSeminarVisitingSeminarian: useServiceAction((id: number) =>
    $apiSemianrian.get<IVisiting.GetSeminarsVisiting>(
      `/attendance/seminar/visiting/${id}`
    )
  ),
  fetchSeminarian: useServiceAction((id: string) =>
    $apiSemianrian.get<IAttendance.GetLessons>(`/attendance/lesson/${id}`)
  ),
  addSeminarVisitingSeminarian: useServiceAction(
    (data: IVisiting.FetchStudentSeminarVisiting) =>
      $apiSemianrian.post<ISeminar[]>('/attendance/seminar/visiting', data)
  ),
  changeSeminarVisitingSeminarian: useServiceAction(
    (data: IVisiting.FetchStudentSeminarVisiting) =>
      $apiSemianrian.put<ISeminar[]>('/attendance/seminar/visiting', data)
  ),
  getLessonDateSeminarian: useServiceAction((data: IAttendance.GetLessonDate) =>
    $apiSemianrian.get<IAttendance.LessonDate>(
      `/attendance/lesson/date?lesson_id=${data.lesson_id}&group_id=${data.group_id}`
    )
  ),
  getDisciplineLessonsSeminarian: useServiceAction((disciplineId: number) =>
    $apiSemianrian.get<IAttendance.GetLessons>(
      `/attendance/lesson/${disciplineId}`
    )
  ),
  getGroupLessonsSeminarian: useServiceAction(
    (data: IAttendance.FecthSeminars) =>
      $apiSemianrian.get<IAttendance.GetLessonsDate>(
        `/attendance/lesson/table?discipline_id=${data.discipline_id}&group_id=${data.group_id}`
      )
  ),
};
