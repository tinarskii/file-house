<script lang="ts" setup>
import { MediaConvert } from "../../wailsjs/go/main/App";
import {ref} from "vue";
import {EventsOn} from "../../wailsjs/runtime";

let videoFormat = ref("mp4"),
    status = ref("idle");

async function convert(e: any) {
  e.preventDefault();

  // Convert file
  await MediaConvert({
    convertValues(): any {},
    target_format: videoFormat.value,
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
    <select id="videoFormat" name="videoFormat" v-model="videoFormat">
        <option value="mp4">MP4</option>
        <option value="avi">AVI</option>
        <option value="wmv">WMV</option>
        <option value="mov">MOV</option>
        <option value="flv">FLV</option>
        <option value="mkv">MKV</option>
        <option value="webm">WEBM</option>
    </select>
    <button type="submit">Convert</button>
    <div class="loading">
      <div class="spinner" v-if="status === 'loading'"></div>
    </div>
  </form>
</template>

<style scoped>
.loading {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}
</style>