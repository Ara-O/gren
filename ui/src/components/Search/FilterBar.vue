<template>
    <div class="flex justify-between items-center px-[101px] h-20 border-b-[0.5px] border-gridcolor">
        <h4 class="text-lg cursor-pointer hidden md:block" tabindex="0">New Arrivals</h4>
        <h4 class="text-lg cursor-pointer hidden md:block" tabindex="0">Categories</h4>
        <h4 class="text-lg cursor-pointer hidden md:block" tabindex="0">Size</h4>
        <h4 class="text-lg cursor-pointer hidden md:block" tabindex="0">Material</h4>
        <h4 class="text-lg cursor-pointer hidden md:block" tabindex="0">Price Range</h4>
        <div class="relative flex items-center w-full md:w-auto">
            <input type="text"
                class="border-[0.5px] border-gridcolor px-2 box-border w-full pr-10 outline-none h-9 md:w-52">
            <img src="../../assets/search-icon.svg" alt="Search icon" class="w-3.5 absolute right-4" tabindex="0">
        </div>
    </div>
</template>
