<template>
  <nav class="flex justify-between px-10 text-sm h-20 items-center font-normal border-gridcolor border-b-[0.5px]">
    <div class="flex">
      <img src="../../assets/user-icon.svg" alt="User Icon" class="w-5 cursor-pointer" tabindex="0"
        @click="openRegistrationPopup" />
      <div class="bg-gridcolor h-20 w-[0.5px] ml-10"></div>
    </div>
    <h1 class="text-2xl">Gren</h1>
    <div class="flex">
      <div class="bg-gridcolor h-20 w-[0.5px] mr-10"></div>
      <img src="../../assets/cart-icon.svg" alt="Cart Icon" class="w-5 cursor-pointer opacity-80" tabindex="0" />
    </div>
  </nav>
  <registration-popup v-if="registrationPopupIsOpen" @close-popup="closeRegistrationPopup"
    @user-already-has-account="changeToLoginPopup"></registration-popup>
  <login-popup v-if="loginPopupIsOpen" @change-to-register-popup="changeToRegisterPopup"
    @close-popup="closeLoginPopup"></login-popup>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import registrationPopup from "../Registration/RegistrationPopup.vue"
import loginPopup from "../Login/LoginPopup.vue"
let registrationPopupIsOpen = ref<boolean>(false)
let loginPopupIsOpen = ref<boolean>(false)

// trying to make everything as obvious as possible
function openRegistrationPopup() {
  registrationPopupIsOpen.value = true
}

function changeToLoginPopup() {
  registrationPopupIsOpen.value = false
  loginPopupIsOpen.value = true
}

function changeToRegisterPopup() {
  registrationPopupIsOpen.value = true
  loginPopupIsOpen.value = false
}

function closeRegistrationPopup() {
  registrationPopupIsOpen.value = false
}

function closeLoginPopup() {
  loginPopupIsOpen.value = false
}
</script>
