<script>
  let {data} = $props()
</script>
<section>
  <div class="flex gap-12">
    {#each data.tutorials as tutorial}
      <div>
        <h1 class="text-3xl font-bold text-[var(--green)] mb-5"># {tutorial.name}</h1>
        <ul class="ml-10 space-y-1">
          {#each tutorial.chapters as chapter}
            <li class="list-disc text-xl"> <a href={chapter.slug}>{chapter.title}</a> </li>
          {/each}
        </ul>
      </div>
    {/each}
  </div>
</section>

<style>
  h1 {
    margin-bottom: 20px;
  }
  ul {
    margin-left: 40px;
  }
</style>