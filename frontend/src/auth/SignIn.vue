<template>
    <div class="min-h-screen bg-[#02192E] text-white flex items-center justify-center px-4">
      <div class="w-full max-w-sm">
        <h2 class="text-center text-2xl/9 font-bold tracking-tight text-white mb-6">
          Sign in to your account
        </h2>
  
        <form class="space-y-6" action="#" method="POST">
          <div>
            <label for="email" class="block text-sm/6 font-medium text-gray-100">Email address</label>
            <div class="mt-2">
              <input
                type="email"
                name="email"
                id="email"
                autocomplete="email"
                class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
              />
            </div>
          </div>
  
          <div>
            <div class="flex items-center justify-between">
              <label for="password" class="block text-sm/6 font-medium text-gray-100">Password</label>
              <div class="text-sm">
                <a href="#" class="font-semibold text-indigo-400 hover:text-indigo-300">Forgot password?</a>
              </div>
            </div>
            <div class="mt-2">
              <input
                type="password"
                name="password"
                id="password"
                autocomplete="current-password"
                class="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
              />
            </div>
          </div>
  
          <div>
            <button
              type="submit"
              class="flex w-full justify-center rounded-md bg-indigo-500 px-3 py-1.5 text-sm/6 font-semibold text-white hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
            >
              Sign in
            </button>
          </div>
        </form>
  
        <p class="mt-10 text-center text-sm/6 text-gray-400">
        Don't have an account?
          {{ ' ' }}
          <a href="#" class="font-semibold text-indigo-400 hover:text-indigo-300">Sign Up</a>
        </p>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  </script>