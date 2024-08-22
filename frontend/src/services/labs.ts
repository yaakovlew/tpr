import { $apiLecturer, $apiSemianrian } from 'src/boot/axios';
import { ILabaratory } from 'src/models/labaratory/labaratory';
import { useServiceAction } from 'src/utils/service/action';

export const LabsService = {
  addLabaratory: useServiceAction((data: ILabaratory.AddLabaratory) =>
    $apiLecturer.post('/laboratory-work/external', data)
  ),
  getExternalLabs: useServiceAction(() =>
    $apiLecturer.get<ILabaratory.GetExternalLabs>('/laboratory-work/external')
  ),
  getLabsFromSection: useServiceAction((id: number) =>
    $apiLecturer.get<ILabaratory.GetLabsFromSection>(
      `/discipline/section/laboratory-work/${id}`
    )
  ),
  deleteLab: useServiceAction((id: number) =>
    $apiLecturer.delete(`/laboratory-work/${id}`)
  ),
  getLabsFromSectionSeminarian: useServiceAction((id: number) =>
    $apiSemianrian.get<ILabaratory.GetLabs>(`/discipline/laboratory-work/${id}`)
  ),
};
