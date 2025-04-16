<script lang="ts">
  import Editor from '$lib/components/editor.svelte';

  let {data} = $props()
</script>

<div class="cntnr">

  <div class="markdown">
    <data.content />
  </div>

  <div class="c flex flex-col overflow-scroll">
    <div>submit</div>
    <Editor />
  </div>

</div>

<style>
  .cntnr, .markdown, .c {
    display: flex;
    height: calc(100vh - var(--header-height));
  }
  .markdown {
    overflow: scroll;
    flex-direction: column;
    gap: 12px;
    padding: 2.5rem;
    width: 50%;
    scrollbar-width: thin;
  }
</style>