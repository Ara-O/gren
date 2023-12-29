<template>
  <button
    class="bg-gren text-white text-xs px-9 py-3.5 rounded-3xl font-normal baseButtonClass"
  >
    <slot></slot>
  </button>
</template>

<script setup lang="ts">
import { defineProps } from "vue";

const props = defineProps({
  bgColor: {
    type: String,
    required: false,
    default: "#0F2D14",
  },
});

let buttonBackgroundColor: String = props.bgColor;
let buttonTextColor: String;

//if the background color is green
if (buttonBackgroundColor === "#0F2D14") {
  buttonTextColor = "white";
} else {
  buttonTextColor = "black";
}
</script>

<style>
.baseButtonClass {
  background-color: v-bind(buttonBackgroundColor);
  color: v-bind(buttonTextColor);
}
</style>
