<script>
  import ProductCard from './ProductCard.svelte';

  let { title, products } = $props();
</script>

<section class="mb-20">
  <div class="inline-block mb-10">
    <h2 class="text-2xl font-semibold text-gray-900">{title}</h2>
    <div class="mt-2 w-12 h-0.5 bg-gradient-to-r from-[#013373] to-[#3b8fce] rounded"></div>
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
    {#each products as product (product.title)}
      <ProductCard {product} />
    {/each}
  </div>
</section>
