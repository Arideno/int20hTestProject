<template>
  <div>
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
  data() {
    return {
      products: [
        {
          name: 'name1',
          price: 50.0,
        },
        {
          name: 'name2',
          price: 100.0,
        }
      ]
    }
  },
  components: {
    Product,
  }
}
</script>