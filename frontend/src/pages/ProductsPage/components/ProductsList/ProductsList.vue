<template>
  <div class="products-grid p-grid">
    <Product
      v-for="product of products"
      :key="product.name"
      class="p-col-12 p-md-3 p-lg-3"
      :product="product"
    />
  </div>
</template>

<script>
import Product from '../Product/Product.vue';

export default {
  name: 'ProductsList',
  components: {
    Product,
  },
  props: {
    products: {
      type: Array,
      required: true,
    },
  },
};
</script>
