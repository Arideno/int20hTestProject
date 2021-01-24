<template>
  <div class="filter-by-shop">
    <div
      v-for="shop of shops"
      :key="shop.name"
      class="p-field-checkbox"
    >
      <Checkbox
        :id="shop.name"
        name="shop"
        :value="shop"
        @select="onShopSelect"
      />
      <label :for="shop.name">{{ shop.name }}</label>
    </div>
  </div>
</template>

<script>
import Checkbox from 'primevue/checkbox';

export default {
  name: 'FilterByShop',
  components: {
    Checkbox,
  },
  props: {
    shops: {
      type: Array,
      required: true,
    },
    selectedShops: {
      type: Array,
      required: true,
    },
    onShopSelect: {
      type: Function,
      required: true,
    }
  },
};
</script>
