<template>
  <div class="side-filters p-d-flex p-flex-column">
    <Accordion
      :multiple="true"
      :active-index="[0, 1]"
    >
      <AccordionTab header="Shop">
        <FilterByShop
          :shops="shops"
          :selected-shop="selectedShop"
          @on-shop-select="onShopSelect"
        />
      </AccordionTab>
      <AccordionTab header="Price">
        <FilterByPrice
          :min="price.min"
          :max="price.max"
          @on-price-select="onPriceSelect"
        />
      </AccordionTab>
    </Accordion>
  </div>
</template>

<script>
import { onMounted, ref, reactive, toRefs, watchEffect } from 'vue';

import ShopService from 'src/api/services/shop.service';
import { parseUrlParams, updateQuerystringParam } from 'src/helpers/qs.helper';

import Accordion from 'primevue/accordion';
import AccordionTab from 'primevue/accordiontab';
import FilterByShop from '../FilterByShop/FilterByShop.vue';
import FilterByPrice from '../FilterByPrice/FilterByPrice.vue';

export default {
  name: 'SideFilters',
  components: {
    Accordion,
    AccordionTab,
    FilterByShop,
    FilterByPrice,
  },
  emits: ['on-filter-change'],

  setup(_, { emit }) {
    const shops = ref([]);
    const filters = reactive({
      shop: null,
      price: {
        min: null,
        max: null,
      },
    });

    function onFilterChange() {
      emit('on-filter-change', filters);
    };

    function onShopSelect(shop) {
      updateQuerystringParam(['shop', 'id'], shop.id);
      filters.shop = shop;
      onFilterChange();
    };

    function onPriceSelect({ min, max }) {
      updateQuerystringParam(['price', 'min'], min);
      updateQuerystringParam(['price', 'max'], max);
      filters.price = { min, max };
      onFilterChange();
    };

    onMounted(async () => {
      shops.value = await ShopService.getShops();

      let { shop: { id: shopID } = {}, price } = parseUrlParams();
      if (!shopID) {
        shopID = shops.value[0]?.id;
      }
      if (!price) {
        price = { min: 0, max: 100 };
      }
      filters.price = price;
      filters.shop = shops.value.find(shop => shop.id === shopID);

      onShopSelect(filters.shop);
      onPriceSelect(filters.price);

      emit('on-filter-change', filters);
    });

    const { shop: selectedShop, price } = toRefs(filters);

    return {
      shops,
      selectedShop,
      price,
      filters,
      onShopSelect,
      onPriceSelect,
      onFilterChange,
    };
  },
};
</script>
