<template>
  <!-- Header -->
  <header id="header" class="bg-gray-700">
    <nav class="container mx-auto flex justify-start items-center py-5 px-4">
      <div class="flex items-center justify-between w-full">
        <div>
          <p class="text-white font-bold uppercase text-2xl mr-4">HyperTube</p>
        </div>
        <div class="">
          <!-- Primary Navigation -->
          <ul class="flex flex-row mt-1">
            <!-- Navigation Links -->
            <li>
              <p class="px-2 text-white" to="/about">About</p>
            </li>

            <li>
              <a class="px-2 text-white">Login / Register</a>
            </li>
            <template>
              <li>
                <p class="px-2 text-white">Search</p>
              </li>
              <li>
                <p class="px-2 text-white">Account</p>
              </li>
              <li>
                <button class="px-2 text-white">Logout</button>
              </li>
            </template>
          </ul>
        </div>
      </div>
      <!-- App Name -->
    </nav>
  </header>
</template>

<script setup lang="ts">
import style from "./../style.css";
</script>
