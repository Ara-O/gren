<template>
  <button class="text-white px-8 py-2 font-normal bg-black" :class="class" :type="props.type">
    <slot></slot>
  </button>
</template>

<script setup lang="ts">
const props = withDefaults(defineProps<{
  class: string,
  type: "button" | "submit" | "reset" | undefined
}>(), {
  class: "",
  type: "submit"
});

</script>

