<template>
  <div class="main flex items-center column q-ma-xs-md">
    <h4>Добро пожаловать</h4>
    <h6 class="text-center">
      Добро пожаловать в систему поддержки обучения по дисциплинам, входящим в блок «Теория принятия решения и мягкие вычисления».
      <br>
      Система доступна для авторизированных пользователей – нажмите на иконку в правом верхнем углу, укажите номер группы и аккуратно введите Вашу фамилию и имя.
    </h6>
    <div class="main__news-container flex items-center column">
      <single-news-container title="Лекции и семинары">
        <template #content>
          <div>
            Видеозаписи лекций и семинаров прошлых лет находятся по <a href="https://disk.yandex.ru/d/cQQGC05bWGPoMA">адресу</a>
          </div>
        </template>
      </single-news-container>
    </div>
  </div>
</template>

<script setup lang="ts">
import SingleNewsContainer from './components/SingleNewsContainer.vue';
</script>

<style scoped lang="scss">
.main {
  &__news-container {
    width: 100%;
    gap: 15px;
  }
}
</style>
