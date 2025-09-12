<script>
  /** @type {{children?: import('svelte').Snippet, [key: string]: any}} */
  const { ...props } = $props();
</script>

<main {...props} class="container mx-auto px-5 {props.class}">
  {@render props.children?.()}
</main>
