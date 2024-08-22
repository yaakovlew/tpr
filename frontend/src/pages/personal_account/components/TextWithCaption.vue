<template>
  <div>
    <div class="caption">
      <slot name="caption" />
      <span v-if="required" class="text-red">*</span>
    </div>
    <div class="text">
      <slot />
    </div>
  </div>
</template>

<script lang="ts" setup>
withDefaults(
  defineProps<{
    required?: boolean;
  }>(),
  {
    required: false,
  }
);
</script>

<style lang="scss" scoped>
.caption {
  font-size: 14px;
  color: #6666;
}

.text {
  font-size: 20px;
  color: rgb(82, 82, 82);
  font-weight: 600;
}
</style>
