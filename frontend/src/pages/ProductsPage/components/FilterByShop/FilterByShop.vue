<template>
  <div class="filter-by-shop">
    <div
      v-for="shop of shops"
      :key="shop.id"
      class="p-field-radiobutton"
    >
      <RadioButton
        :id="shop.id"
        :model-value="selectedShop"
        name="shop"
        :value="shop"
        @change="onShopSelect(shop)"
      />
      <label :for="shop.id">{{ shop.name }}</label>
    </div>
  </div>
</template>

<script>
import RadioButton from 'primevue/radiobutton';

export default {
  name: 'FilterByShop',
  components: {
    RadioButton,
  },
  props: {
    shops: {
      type: Array,
    },
    selectedShop: {
      type: [Object, null],
    },
  },
  emits: ['on-shop-select'],

  methods: {
    onShopSelect(shop) {
      this.$emit('on-shop-select', shop);
    },
  },
};
</script>
