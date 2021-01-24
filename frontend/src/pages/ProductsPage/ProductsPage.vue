<template>
  <div class="products-page p-grid">
    <SideFilters class="p-xl-2" />
    <ProductsGrid
      class="p-xl-10"
      :products="products"
    />
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';

import ProductService from 'src/api/services/product.service';

import ProductsGrid from './components/ProductsGrid/ProductsGrid.vue';
import SideFilters from './components/SideFilters/SideFilters.vue';

export default {
  name: 'ProductsPage',
  components: {
    ProductsGrid,
    SideFilters,
  },
  setup() {
    let products = ref([]);

    onMounted(async () => {
      products.value = await ProductService.getProducts();
    });

    return {
      products,
    };
  },
};
</script>
