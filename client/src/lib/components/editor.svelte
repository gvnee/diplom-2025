<script lang="ts">
  import CodeMirror from "svelte-codemirror-editor";
  import {catppuccin} from "codemirror-theme-catppuccin"
  import { cpp } from "@codemirror/lang-cpp";

  let {value = $bindable(), ...props} = $props()
</script>

<CodeMirror
  bind:value={value} {...props}
  class="editor"
  theme={catppuccin("macchiato")}
  lang={cpp()}
  styles={{
    "&": {
      width: "50vw",
      height: "calc(100vh - var(--header-height) - var(--submit-height))",
      overflow: "scroll",
      fontSize: "16px",
    },
  }}
/>