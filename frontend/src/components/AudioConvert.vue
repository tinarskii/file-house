<script lang="ts" setup>
import { MediaConvert } from "../../wailsjs/go/main/App";
import {ref} from "vue";
import {EventsOn} from "../../wailsjs/runtime";

let audioFormat = ref("mp3");

async function convert(e: any) {
  e.preventDefault();

  // Convert file
  await MediaConvert({
    convertValues(): any {},
    target_format: audioFormat.value,
    target_bitrate: {
      audio: "128k",
      video: "128k",
    }
  });
}
EventsOn("success", () => {
  alert("File converted successfully");
});

EventsOn("failed", () => {
  alert("Failed to convert file");
});
</script>

<template>
  <form @submit="convert">
    <select id="audioFormat" name="audioFormat" v-model="audioFormat">
      <option value="mp3" >MP3</option>
      <option value="ogg" >OGG</option>
      <option value="wav" >WAV</option>
      <option value="aac" >AAC</option>
      <option value="wma" >WMA</option>
      <option value="flac">FLAC</option>
      <option value="m4a" >M4A</option>
      <option value="opus">OPUS</option>
    </select>
    <button type="submit">Convert</button>
  </form>
</template>
