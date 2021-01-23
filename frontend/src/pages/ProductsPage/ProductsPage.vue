<template>
  <div class="products-page">
    <ProductsList :products="[]" />
  </div>
</template>

<script>
import ProductsList from './components/ProductsList/ProductsList.vue';

import ProductService from 'src/api/services/product.service';

export default {
  name: 'ProductsPage',
  components: {
    ProductsList,
  },

  async setup() {
    const products = await ProductService.getProducts();

    return {
      products,
    } 
  }
}
</script>