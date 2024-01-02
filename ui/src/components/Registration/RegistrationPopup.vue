<template>
    <section class="flex flex-col items-center justify-center">
        <div class="absolute left-0 top-0 w-screen h-screen z-10 bg-[#000000af]" @click="emits('closePopup')"></div>
        <section class="absolute top-0 flex items-center w-full h-full justify-center">
            <Form :validation-schema="schema" class="bg-white z-20 min-h-[700px] w-[500px] box-border px-12 py-12"
                @submit="register">
                <div class="flex justify-between">
                    <h3 class="text-3xl font-bold">Sign Up</h3>
                    <img src="../../assets/cancel-icon.svg" alt="Cancel icon" class="w-4 cursor-pointer items-center"
                        @click="emits('closePopup')">
                </div>
                <div class="mt-7">
                    <label for="name" class="text-lg">Full Name</label>
                    <br>
                    <Field name="full_name"
                        class="border-[0.5px] outline-none box-border px-3 mt-3 mb-2 w-full h-10 border-gridcolor" id="name"
                        value="Ara" />
                    <ErrorMessage name="full_name" class=" text-red-600"></ErrorMessage>
                </div>
                <div class="mt-5">
                    <label for="email" class="text-lg">Email Address</label>
                    <br>
                    <Field name="email_address"
                        class="border-[0.5px] outline-none box-border px-3 mt-3 mb-2 w-full h-10 border-gridcolor"
                        id="email" value="ara@gmail.com" />
                    <ErrorMessage name="email_address" class=" text-red-600"></ErrorMessage>

                </div>
                <div class="mt-5">
                    <label for="password" class="text-lg">Password</label>
                    <br>
                    <Field type="password" name="password"
                        class="border-[0.5px] outline-none box-border px-3 mt-3 mb-2 w-full h-10 border-gridcolor"
                        id="password" value="bob" />
                    <ErrorMessage name="password" class=" text-red-600"></ErrorMessage>

                </div>

                <h4 class="mt-4 hover:underline">Already have an account?</h4>

                <base-button type="submit" class="mt-6">Sign up</base-button>
                <h4 class="text-red-600 mt-5">{{ errorMessage }}</h4>
                <div class="relative flex items-center justify-center mt-3">
                    <div class="absolute bg-black h-[0.5px] w-full "></div>
                    <h4 class="inline bg-white z-10 w-10 text-center">or</h4>
                </div>

                <div class="flex gap-3 items-center mt-5">
                    <h4>Sign up with: </h4>
                    <img @click="test" src="../../assets/github-icon.png" class="w-5 mt-0 cursor-pointer" alt="Github icon">
                </div>
            </Form>
        </section>
    </section>
</template>

<script lang="ts" setup>
import BaseButton from '../Buttons/BaseButton.vue';
import { Form, Field, ErrorMessage } from 'vee-validate'
import { object, string } from "yup"
import { ref } from 'vue';
import axios from "axios"

const emits = defineEmits(['closePopup'])

let errorMessage = ref<string>("")

const schema = object({
    "full_name": string().trim().required('Full name is required').min(2, "Full name must be at least 2 characters"),
    "email_address": string().email("Please enter a valid email address").required("Please enter a valid email address"),
    "password": string().min(8, "Password must be at least 8 characters")

})

async function register(values: any) {
    try {
        errorMessage.value = ""
        console.log('Registering...', values)
        let res = await axios.post(`${import.meta.env.VITE_API_URL}/api/register`, values)
        console.log(res)
    } catch (err: any) {
        console.log(err)
        errorMessage.value = err?.response?.data || "An error occured"

    }
}

function test() {
    axios.get("http://localhost:8080/hi")
}
</script>