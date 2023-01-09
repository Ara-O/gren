<template>
  <h3 class="mt-12 font-medium mb-8">Best Selling Products</h3>
  <section class="flex gap-x-16">
    <div class="flex gap-x-6">
      <best-selling-products-item></best-selling-products-item>
      <best-selling-products-item></best-selling-products-item>
      <best-selling-products-item></best-selling-products-item>
      <best-selling-products-item></best-selling-products-item>
    </div>
    <div class="bg-bestSellingProductItem w-full rounded-xl p-8">
      <div class="flex gap-x-12 items-center">
        <img
          src="../../assets/image-test.png"
          class="h-76 w-64 object-cover rounded-xl"
          alt="Selected item image"
        />
        <article class="w-48">
          <h3 class="font-medium">Apple Watch</h3>
          <h4 class="font-medium mt-2">$3000.00</h4>
          <h5 class="mt-2 text-[14px] font-light">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit. Dolores.
          </h5>
        </article>
      </div>
      <base-button class="mt-6">Add to Cart</base-button>
    </div>
  </section>
</template>
<script lang="ts" setup>
import BestSellingProductsItem from "./BestSellingProductsItem.vue";
</script>
