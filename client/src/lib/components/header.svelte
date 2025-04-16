<script>
  import logo from "$lib/assets/glorp.webp";
</script>

<nav>
  <div>
    <a href="/" class="link">
      <img class="link" src={logo} alt="logo" width="48" height="48" />
    </a>
  </div>
  <ul>
    <iconify-icon class="link" icon="ic:baseline-discord" width="32" height="32"></iconify-icon>
    <iconify-icon class="link" icon="mdi:github" width="32" height="32"></iconify-icon>
  </ul>
</nav>

<style>
  nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    align-items: center;
    height: var(--header-height);
    width: 100%;
    position: fixed;
    top: 0;
    z-index: 100;
    background-color: var(--crust);
    border-bottom: 1px solid var(--base);
  }
  ul {
    display: flex;
    gap: 10px;
  }
  .link {
    transition-duration: 300ms;
    cursor: pointer;
  }
  .link:hover {
    transform: scale(1.2);
  }
</style>
