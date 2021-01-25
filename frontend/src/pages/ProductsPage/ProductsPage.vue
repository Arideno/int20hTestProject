<template>
  <div class="products-page">
    <Section class="p-grid">
      <SideFilters
        class="p-xl-2"
        @on-filter-change="onFilterChange"
      />
      <ProductsGrid
        class="p-xl-10"
        :products="products"
      />
    </Section>
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';

import ProductService from 'src/api/services/product.service';

import ProductsGrid from './components/ProductsGrid/ProductsGrid.vue';
import SideFilters from './components/SideFilters/SideFilters.vue';
import Section from 'src/components/Section/Section.vue';

export default {
  name: 'ProductsPage',
  components: {
    ProductsGrid,
    SideFilters,
    Section,
  },
  setup() {
    const products = ref([]);

    onMounted(async () => {
      products.value = await ProductService.getProducts();
    });

    return {
      products,
    };
  },

  methods: {
    async onFilterChange(filter) {
      const products = await ProductService.getProducts(filter);
      this.products = products;
    },
  },
};
</script>
