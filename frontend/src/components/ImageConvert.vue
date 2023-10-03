<script lang="ts" setup>
import { ImageConvert } from "../../wailsjs/go/main/App";
import {ref} from "vue";
import {EventsOn} from "../../wailsjs/runtime";

let imageFormat = ref("jpg");

async function convert(e: any) {
  e.preventDefault();

  await ImageConvert({
    target_format: imageFormat.value,
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
    <select id="imageFormat" name="imageFormat" v-model="imageFormat">
      <option value="jpg" >JPG</option>
      <option value="png" >PNG</option>
      <option value="gif" >GIF</option>
      <option value="webp">WEBP</option>
      <option value="bmp" >BMP</option>
      <option value="tiff">TIFF</option>
    </select>
    <button type="submit">Convert</button>
  </form>
</template>
