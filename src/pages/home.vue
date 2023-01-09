<template>
  <navigation-bar></navigation-bar>
  <section
    class="rounded-3xl h-[28rem] bg-hero bg-[size:796px] bg-[-193px] text-white flex items-center pl-44"
  >
    <div>
      <h3 class="text-white text-4xl font-semibold mb-4">
        Deals, Deals, Deals!
      </h3>
      <h5 class="mb-8 text-[15px]">Lorem, ipsum dolor amet!</h5>
      <base-button bgColor="white">Add to Cart</base-button>
    </div>
  </section>

  <!-- Best Selling Products -->
  <best-selling-products></best-selling-products>
</template>

<script setup lang="ts">
import navigationBar from "../components/Navigation/NavigationBar.vue";
import BestSellingProducts from "../components/Pages/BestSellingProducts.vue";
</script>
