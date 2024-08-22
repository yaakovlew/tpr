import { defineStore } from 'pinia';
import { ReportService } from '../services/report';

export const useReportStore = defineStore('report', () => {
  const getReport = async (
    groupId: number,
    disciplineId: number,
    isExam: boolean
  ) => {
    const res = await ReportService.getGroupDisciplineReport({
      group_id: groupId,
      discipline_id: disciplineId,
      is_exam: isExam,
    });
    return res;
  };

  const getReportSeminarian = async (
    groupId: number,
    disciplineId: number,
    isExam: boolean
  ) => {
    const res = await ReportService.getGroupDisciplineReportSeminarian({
      group_id: groupId,
      discipline_id: disciplineId,
      is_exam: isExam,
    });
    return res;
  };

  return { getReport, getReportSeminarian };
});
