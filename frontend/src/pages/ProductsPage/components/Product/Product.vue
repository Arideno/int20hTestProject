<template>
  <div class="product-card-wrapper">
    <Card style="position: relative;">
      <template #header>
        <a :href="product.url">
          <img
            :src="product.image"
            :alt="product.name"
          >
        </a>
      </template>
      <template #title >
        <div class="p-d-flex title">
        <span class="product-store"> {{ product.store }} </span>
        <span class="product-price"> {{ product.price }} ₴ </span>
      </div></template>
      <template #content>
        <a :href="product.url" class="product-name">
          {{ product.name }}
        </a>
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
