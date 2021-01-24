<template>
  <div class="product-card-wrapper">
    <Card>
      <template #header>
        <img
          :src="product.image"
          :alt="product.name"
        >
      </template>
      <template #title>
        <span class="product-store"> {{ product.store }} </span>
        -
        <span class="product-price"> ₴{{ product.price }} </span>
      </template>
      <template #content>
        <div class="product-name">
          {{ product.name }}
        </div>
      </template>
    </Card>
  </div>
</template>

<script>
import Card from 'primevue/card';

export default {
  name: 'Product',
  components: {
    Card,
  },
  props: {
    product: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style src="./styles.scss" lang="scss" scoped></style>
