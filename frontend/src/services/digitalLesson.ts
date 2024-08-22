import { $apiLecturer, $apiStudent } from 'src/boot/axios';
import { IDigitalLesson } from 'src/models/digitalLesson/digital-lesson';
import { useServiceAction } from 'src/utils/service/action';

export const DigitalLessonService = {
  getDigitalLessons: useServiceAction(() =>
    $apiLecturer.get<IDigitalLesson.GetDigitalLessons>('/material')
  ),
  addDigitalLesson: useServiceAction((data: IDigitalLesson.AddDigitalLesson) =>
    $apiLecturer.post<IDigitalLesson.DigitalLesson[]>('/material', data)
  ),
  addDigitalLessonToDiscipline: useServiceAction(
    (data: IDigitalLesson.AddDigitalLessonToDiscipline) =>
      $apiLecturer.post('/material/digital-lesson', data)
  ),
  changeDigitalLesson: useServiceAction(
    (data: IDigitalLesson.ChangeDigitalLesson) => {
      console.log(data);
      return $apiLecturer.put<IDigitalLesson.DigitalLesson>(
        `/material?digital_guide_id=${data.digital_guide_id}&description=${data.description}&description_en=${data.description_en}&name=${data.name}&name_en=${data.name_en}`
      );
    }
  ),
  getDisciplineLessons: useServiceAction((disciplineId: number) =>
    $apiLecturer.get<
      Record<'digital_guides', IDigitalLesson.AddDigitalLessonToDiscipline[]>
    >(`/material/digital-lesson/${disciplineId}`)
  ),
  getDigitalLesson: useServiceAction((id: number) =>
    $apiLecturer.get<IDigitalLesson.GetDigitalLessons>(`/material/${id}`)
  ),
  removeDigitalLessonFromDiscipline: useServiceAction(
    (data: IDigitalLesson.AddDigitalLessonToDiscipline) =>
      $apiLecturer.delete(
        `/material/digital-lesson?digital_guide_id=${data.digital_material_id}&discipline_id=${data.discipline_id}`
      )
  ),
  removeDigitalLesson: useServiceAction((id: number) =>
    $apiLecturer.delete(`/material/${id}`)
  ),
  getStudentMaterials: useServiceAction((id: number) =>
    $apiStudent.get<IDigitalLesson.GetDigitalLessonsStudent>(`/material/${id}`)
  ),
  uploadDigitalLesson: useServiceAction(
    (data: IDigitalLesson.AddDigitalMaterial) =>
      $apiLecturer.post(`/material/upload/${data.id}`, data.file)
  ),
  downloadDigitalLesson: useServiceAction((id: number) =>
    $apiLecturer.get(`/material/download/${id}`, {
      responseType: 'blob',
    })
  ),
  getFilesId: useServiceAction((id: number) =>
    $apiLecturer.get<IDigitalLesson.GetFileIds>(`/material/${id}`)
  ),
};
