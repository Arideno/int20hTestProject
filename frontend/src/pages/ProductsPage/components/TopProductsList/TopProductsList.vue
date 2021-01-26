<template>
  <div class="top-products-list">
    <Product
      v-for="product of topProducts"
      :key="product.name"
      class="p-col-12 p-md-6 p-lg-3"
      :product="product"
    />
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';

import TopProductsService from 'src/api/services/top-products.service';

import Product from '../Product/Product.vue';

export default {
  name: 'TopProductsList',
  components: {
    Product,
  },
  setup() {
    const topProducts = ref([]);

    onMounted(async () => {
      topProducts.value = await TopProductsService.getTopProducts();
    });

    return {
      topProducts,
    };
  },
};
</script>

<style src="./styles.scss" lang="scss" scoped></style>
