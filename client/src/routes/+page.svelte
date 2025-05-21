<script>
  let {data} = $props()
</script>
<section>
  <div class="flex gap-5">
    {#each data.tutorials as tutorial}
      <div>
        <h1 class="text-3xl font-bold text-[var(--green)]"># {tutorial.name}</h1>
        <ul>
          {#each tutorial.chapters as chapter}
            <li> <a href={chapter.slug}>{chapter.title}</a> </li>
          {/each}
        </ul>
      </div>
    {/each}
  </div>
</section>