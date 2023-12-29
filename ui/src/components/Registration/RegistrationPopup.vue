<template>
    <section class="flex flex-col items-center justify-center">
        <div class="absolute left-0 top-0 w-screen h-screen z-10 bg-[#000000af]" @click="emits('closePopup')"></div>
        <section class="absolute top-0 flex items-center w-full h-full justify-center">
            <form class="bg-white z-20 min-h-[700px] w-[500px] box-border px-12 py-12" @submit.prevent="register">
                <div class="flex justify-between">
                    <h3 class="text-3xl font-bold">Sign Up</h3>
                    <img src="../../assets/cancel-icon.svg" alt="Cancel icon" class="w-4 cursor-pointer items-center"
                        @click="emits('closePopup')">
                </div>
                <div class="mt-7">
                    <label for="name" class="text-lg">Full Name</label>
                    <br>
                    <input type="text" placeholder="@gmail.com"
                        class="border-[0.5px] outline-none box-border px-3 mt-3 w-full h-10 border-gridcolor" id="name"
                        required value="Bob Bobbington">
                </div>
                <div class="mt-7">
                    <label for="email" class="text-lg">Email Address</label>
                    <br>
                    <input type="text" placeholder="@gmail.com"
                        class="border-[0.5px] outline-none box-border px-3 mt-3 w-full h-10 border-gridcolor" id="email"
                        required value="bobbington@gmail.com">
                </div>
                <div class="mt-7">
                    <label for="password" class="text-lg">Password</label>
                    <br>
                    <input type="password" placeholder="xxxxxxx"
                        class="border-[0.5px] outline-none box-border px-3 mt-3 w-full h-10 border-gridcolor" id="password"
                        required value="bob">
                </div>

                <h4 class="mt-7 hover:underline">Already have an account?</h4>

                <base-button class="mt-6" type="submit">Sign up</base-button>

                <div class="relative flex items-center justify-center mt-7">
                    <div class="absolute bg-black h-[0.5px] w-full "></div>
                    <h4 class="inline bg-white z-10 w-10 text-center">or</h4>
                </div>

                <div class="flex gap-3 items-center mt-5">
                    <h4>Sign up with: </h4>
                    <img src="../../assets/github-icon.png" class="w-5 mt-0 cursor-pointer" alt="Github icon">
                </div>
            </form>
        </section>
    </section>
</template>

<script lang="ts" setup>
import BaseButton from '../Buttons/BaseButton.vue';
import axios from "axios"

const emits = defineEmits(['closePopup'])

async function register() {
    try {

        console.log('Registering')
        await axios.post(`${import.meta.env.VITE_API_URL}/api/register`)
    } catch (err) {
        console.log(err)
    }
}
</script>