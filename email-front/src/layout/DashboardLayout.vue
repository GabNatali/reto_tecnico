<template>
  <div class="flex flex-col w-screen min-h-screen bg-[#f5f8fa] box-border">
    <div class="relative flex flex-0 items-center w-full h-16 px-4 md:px-6 z-49 shadow bg-white">
      <p class="text-xl font-bold">Dashboard</p>
    </div>
    <div class="flex flex-col lg:h-[calc(100vh-7rem)] py-4 px-8 w-[calc(100)]">
      <RouterView></RouterView>
    </div>
  </div>
</template>
