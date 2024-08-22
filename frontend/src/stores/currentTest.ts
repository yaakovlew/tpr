import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { IQuestion } from 'src/models/test/question';
import { TestService } from 'src/services/test';
import { ITest } from 'src/models/test/test';

export const useCurrentTestStore = defineStore('current-test', () => {
  const currentTest: Ref<number | null> = ref(
    localStorage.getItem('current-test') !== 'null'
      ? Number(localStorage.getItem('current-test'))
      : null
  );
  const testDuration: Ref<number | null> = ref(
    localStorage.getItem(`test-duration-${currentTest.value}`) !== 'null'
      ? Number(localStorage.getItem(`test-duration-${currentTest.value}`))
      : null
  );
  const testStart: Ref<number | null> = ref(
    localStorage.getItem(`test-start-${currentTest.value}`) !== 'null'
      ? Number(localStorage.getItem(`test-start-${currentTest.value}`))
      : null
  );
  const closeDate: Ref<number | null> = ref(
    localStorage.getItem(`test-close-date-${currentTest.value}`) !== 'null'
      ? Number(localStorage.getItem(`test-close-date-${currentTest.value}`))
      : null
  );
  const test: Ref<IQuestion.StudentTest | null> = ref(null);

  const setCurrentTest = async (newTest: number | null) => {
    currentTest.value = newTest;
    localStorage.setItem('current-test', String(newTest));
  };

  const setTestDuration = (newDuration: number | null) => {
    if (localStorage.getItem(`test-${currentTest.value}`)) return;
    testDuration.value = newDuration;
    localStorage.setItem(
      `test-duration-${currentTest.value}`,
      String(newDuration)
    );
  };

  const setTestStart = (newStart: number | null) => {
    if (localStorage.getItem(`test-${currentTest.value}`)) return;
    testStart.value = newStart;
    localStorage.setItem(`test-start-${currentTest.value}`, String(newStart));
  };

  const setCloseDate = (newDate: number | null) => {
    if (localStorage.getItem(`test-${currentTest.value}`)) return;
    closeDate.value = newDate;
    localStorage.setItem(
      `test-close-date-${currentTest.value}`,
      String(newDate)
    );
  };

  const isInRange = () =>
    testStart.value !== null && testDuration.value !== null
      ? Math.floor(new Date().getTime() / 1000) > testStart.value &&
        new Date().getTime() < testStart.value + testDuration.value * 120
      : false;

  const getTest = async () => {
    if (currentTest.value) {
      const res = await TestService.getTest(currentTest.value);
      if (res.data) {
        test.value = res.data.ru;
      }
    }
  };

  const passTest = async (data: ITest.PassTest) => {
    const res = await TestService.passTest(data);
    if (res.data) {
      localStorage.removeItem(`test-duration-${currentTest.value}`);
      localStorage.removeItem(`test-start-${currentTest.value}`);
      localStorage.removeItem(`test-close-date-${currentTest.value}`);
    }
    return res;
  };

  return {
    test,
    currentTest,
    setCurrentTest,
    getTest,
    passTest,
    testDuration,
    setTestDuration,
    testStart,
    setTestStart,
    closeDate,
    setCloseDate,
    isInRange,
  };
});
