<template>
  <div class="products-list">
    <Product
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