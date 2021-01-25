<template>
  <div class="filter-by-price">
    <div class="p-d-flex p-flex-column">
      <div class="p-d-flex price-input-group p-pb-4 p-jc-between p-ai-center">
        <InputNumber
          v-model="minPrice"
          input-class="price-input"
          mode="decimal"
          :min-fraction-digits="2"
          :max-fraction-digits="2"
        />
        <span class="divider">-</span>
        <InputNumber
          v-model="maxPrice"
          input-class="price-input"
          mode="decimal"
          :min-fraction-digits="2"
          :max-fraction-digits="2"
        />
        <Button
          class="p-text-center"
          label="OK"
          @click="onPriceSelect"
        />
      </div>
      <Slider
        :model-value="[minPrice,maxPrice]"
        :range="true"
        :step="0.1"
        @change="onSliderUpdate"
      />
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

import Slider from 'primevue/slider';
import InputNumber from 'primevue/inputnumber';
import Button from 'primevue/button';

import { parseUrlParams } from 'src/helpers/qs.helper';

export default {
  name: 'FilterByPrice',
  components: {
    Slider,
    InputNumber,
    Button,
  },
  emits: ['on-price-select'],

  setup() {
    const { price: { min, max } = {} } = parseUrlParams();

    const minPrice = ref(min || 0);
    const maxPrice = ref(max || 100);

    return {
      minPrice,
      maxPrice,
    };
  },
  data() {
    return {
      priceRange: [this.minPrice, this.maxPrice],
    };
  },

  methods: {
    onPriceSelect() {
      this.$emit('on-price-select', { min: this.minPrice, max: this.maxPrice });
    },
    onSliderUpdate([min, max]) {
      this.minPrice = min;
      this.maxPrice = max;
    },
  },
};
</script>

<style src="./styles.scss" lang="scss" scoped></style>
