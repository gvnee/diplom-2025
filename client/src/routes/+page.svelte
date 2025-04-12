<script>
  let {data} = $props()
</script>

<section>
  <ul>
    {#each data.posts as post}
    <li>
      <a href={post.slug}>{post.title}</a>
    </li>
    {/each}
  </ul>
</section>