<script lang="ts">
  import Editor from '$lib/components/editor.svelte';

  let {data} = $props()
  
  let editorDiv: HTMLDivElement

  const submit = async () => {
    editorDiv.scroll({top: editorDiv.scrollHeight, behavior: 'smooth'})
  }
</script>

<div class="cntnr">

  <div class="markdown">
    <h1>{data.meta.title}</h1>
    <data.content />
  </div>

  <div bind:this={editorDiv} class="overflow-scroll">

    <Editor />

    <div class="h-[var(--submit-height)] flex items-center justify-end p-3">
      <button type="button" onclick={submit} class="rounded-lg bg-[var(--green)] cursor-pointer px-3 py-1 text-[var(--crust)]">
        submit
      </button>
    </div>


    <div class="terminal h-52">
      output
    </div>

  </div>

</div>

<style>
  .cntnr, .markdown {
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
    border-right: 1px solid var(--base);
  }
</style>