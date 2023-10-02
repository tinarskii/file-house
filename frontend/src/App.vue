<script lang="ts" setup>
import { MediaConvert, WriteFile, RemoveFile } from "../wailsjs/go/main/App";

async function convert(e: any) {
  e.preventDefault();

  // Write file to local
  const file = e.target.files[0] as File
  const fileText = await (file).text();
  await WriteFile(fileText, `\$TMP/${file.name}`);

  // Convert file
  const convertedFile = await MediaConvert({
    source: `\$TMP\\${file.name}`,
    target_dir: `\$TMP`,
    target_name: `filehouse_${file.name}`,
    target_format: "ogg",
    target_audio_codec: "libopus",
    target_video_codec: "copy"
  });

  // Download file
  const blob = new Blob([convertedFile as any], { type: "audio/mp3" });
  const link = document.createElement("a");
  link.href = window.URL.createObjectURL(blob);
  link.download = "output.mp3";
  link.click();

  // Remove file from local
  // await RemoveFile(`\$TMP/${file.name}`);
}
</script>

<template>
  <input type="file" @change="convert">
</template>
