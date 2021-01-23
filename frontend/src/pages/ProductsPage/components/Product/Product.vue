<template>
  <div class="product">
    Name: {{ product.name }}
    Price: {{ product.price }}
  </div>
</template>

<script>
export default {
  name: 'Product',
  props: {
    product: Object
  }
}
</script>