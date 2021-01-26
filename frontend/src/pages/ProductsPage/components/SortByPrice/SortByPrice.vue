<template>
  <div class="sort-by-price">
    <Dropdown
      :model-value="selectedSortOption.value"
      :options="sortOptions"
      option-label="name"
      placeholder="Sort by Price"
      @change="onPriceSortOptionChange"
    />
  </div>
</template>

<script>
import { onMounted, reactive } from 'vue';

import { parseUrlParams, updateQuerystringParam } from 'src/helpers/qs.helper';

import Dropdown from 'primevue/dropdown';

export default {
  name: 'SortByPrice',
  components: {
    Dropdown,
  },
  emits: ['on-price-sort-option-change'],

  setup(_, { emit }) {
    const sortOptions = [
      {
        name: 'From low to high',
        type: 'asc',
        id: 'price-asc',
      },
      {
        name: 'From high to low',
        type: 'desc',
        id: 'price-desc',
      },
    ];

    const selectedSortOption = reactive({
      value: {
        name: null,
        type: null,
        id: null,
      },
    });

    function onPriceSortOptionChange(priceSortOption) {
      const priceSortOptionType = priceSortOption?.value?.type;
      updateQuerystringParam(['sort', 'price'], priceSortOptionType);
      emit('on-price-sort-option-change', priceSortOptionType);
      selectedSortOption.value = priceSortOption.value;
    };

    onMounted(async () => {
      const { sort: { price: priceSortType } = {} } = parseUrlParams();
      if (!priceSortType) {
        selectedSortOption.value = sortOptions[1];
      } else {
        const selectedSortOptionFromUrl = sortOptions.find(option => option.type === priceSortType) || {};
        selectedSortOption.value = selectedSortOptionFromUrl;
      }
    });

    return {
      sortOptions,
      selectedSortOption,
      onPriceSortOptionChange,
    };
  },
};
</script>
