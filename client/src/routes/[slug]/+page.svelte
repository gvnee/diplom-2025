<script lang="ts">
import CodeMirror from "svelte-codemirror-editor";
import {catppuccin} from "codemirror-theme-catppuccin"
import { cpp } from "@codemirror/lang-cpp";

let value = $state("")
let {data} = $props()
</script>

<div class="flex flex-row">
  <div class="flex flex-col gap-3 p-10 w-1/2">
    <data.content />
  </div>

  <div class="w-[1px] bg-[#313244]"></div>

  <!-- <div class="sticky top-0">
    fffff
  </div> -->

  <CodeMirror
    bind:value
    theme={catppuccin('macchiato')}
    lang={cpp()}
    styles={{
      '&': {
        width: "50vw",
        height: "100vh",
        fontSize: "16px"
      }
    }}
  />
</div>