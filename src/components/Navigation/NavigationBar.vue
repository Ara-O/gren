<template>
  <nav class="flex justify-between text-sm items-center font-normal mb-7">
    <div class="flex gap-x-32 items-center">
      <div class="h-8 w-8 bg-gren rounded"></div>
      <router-link to="/">Home</router-link>
      <router-link to="/">Deals</router-link>
      <!-- <router-link to="/">Deals</router-link> -->
      <img
        src="../../assets/search-icon.png"
        alt="Search icon"
        class="w-3.5 opacity-60"
      />
    </div>
    <!-- <div class="relative">
      <input
        type="text"
        class="w-80 rounded-2xl h-9 pl-5"
        placeholder="Search"
      />
     
    </div> -->
    <div class="flex gap-x-32 items-center">
      <span class="flex place-items-center gap-x-3">
        <img
          src="../../assets/account-icon.png"
          alt="Account Icon"
          class="w-3.5"
        />
        <router-link to="/">Account</router-link>
      </span>
      <span class="flex place-items-center gap-x-3">
        <img src="../../assets/cart-icon.png" alt="Cart Icon" class="w-3.5" />
        <router-link to="/">Cart</router-link>
      </span>
      <baseButton>Sign Up</baseButton>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
onMounted(() => {
  console.log("onmoeunted");
});
</script>
