<template>
  <div class="bg-[#02192E] text-white min-h-screen">
    <div class="container mx-auto py-8">
      <div class="container mx-auto py-8">
        <h1 class="text-2xl text-center font-bold mb-4">
          Welcome to HyperTube
        </h1>
      </div>
      <div>
        <label for="Search">
          <span class="text-sm font-medium text-white"> Search </span>

          <div class="relative">
            <input
              type="text"
              id="Search"
              placeholder="Enter Keywords"
              class="p-2 mt-0.5 w-full rounded border-gray-300 shadow-sm text-gray-700 sm:text-sm bg-white"
            />

            <span
              class="absolute inset-y-0 right-2 grid w-8 place-content-center text-white"
            >
              <button
                type="button"
                aria-label="Submit"
                class="rounded-full p-1.5 text-gray-700 transition-colors hover:bg-gray-100"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  class="size-4"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"
                  />
                </svg>
              </button>
            </span>
          </div>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>
