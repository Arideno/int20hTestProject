<template>
  <div class="products-page">
    <ProductsList :products="products" />
  </div>
</template>

<script>
import ProductsList from './components/ProductsList/ProductsList.vue';
import ProductService from 'src/api/services/product.service';

import { onMounted, ref } from 'vue';

export default {
  name: 'ProductsPage',
  components: {
    ProductsList,
  },
  setup() {
    let products = ref([]);

    onMounted(async () => {
      products.value = await ProductService.getProducts();
    });

    return {
      products
    }
  }
}
</script>