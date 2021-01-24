<template>
  <div class="product-grid-item card">
    <div class="product-grid-item-top">
      <div>
        <span class="product-store">
          {{product.store}}
        </span>
      </div>
    </div>
    <div class="product-grid-item-content">
      <img :src="product.image" :alt="product.name"/>
      <div class="product-name">
        {{product.name}}
      </div>
    </div>
    <div class="product-grid-item-bottom">
      <span class="product-price">
        ₴{{product.price}}
      </span>
    </div>
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