<template>
  <div class="products-page">
    <ProductsList :products="products" />
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
  data() {
    return {
      products: []
    }
  },
  async mounted() {
    this.products = await ProductService.getProducts();
  }
}
</script>