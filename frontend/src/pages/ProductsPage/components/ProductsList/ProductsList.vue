<template>
  <div class="products-grid p-grid">
    <Product
      class="p-col-12 p-md-3 p-lg-3"
      v-for="product of products"
      :product="product"
      :key="product.name"
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
    products: Array
  }
}
</script>