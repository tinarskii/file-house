<script setup lang="ts">
import AudioConvert from "./components/AudioConvert.vue";
import VideoConvert from "./components/VideoConvert.vue";
import ImageConvert from "./components/ImageConvert.vue";
import {ref} from "vue";

let currentPage = ref("home")
</script>

<template>
  <div v-if="currentPage === 'home'">
    <h1>FileHouse</h1>
    <p>Convert your files without pain™️</p>
    <button @click="currentPage = 'toAudio'">🔉 Audio Convert</button>
    <button @click="currentPage = 'toVideo'">🎥 Video Convert</button>
    <button @click="currentPage = 'toImage'">🖼️ Image Convert</button>
  </div>
    <AudioConvert v-if="currentPage === 'toAudio'" />
    <VideoConvert v-if="currentPage === 'toVideo'" />
    <ImageConvert v-if="currentPage === 'toImage'" />
  <button v-if="currentPage !== 'home'" @click="currentPage = 'home'">🏠 Home</button>
</template>

<style scoped>

</style>