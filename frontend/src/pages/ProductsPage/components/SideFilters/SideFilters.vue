<template>
  <div class="side-filters p-d-flex p-flex-column">
    <Accordion
      :multiple="true"
      :active-index="[0]"
    >
      <AccordionTab header="Shop">
        <FilterByShop
          :shops="shops"
          :selected-shop="selectedShop"
          @on-shop-select="onShopSelect"
        />
      </AccordionTab>
    </Accordion>
  </div>
</template>

<script>
import { onMounted, ref, reactive, toRefs } from 'vue';

import ShopService from 'src/api/services/shop.service';
import { parseUrlParams, updateQuerystringParam } from 'src/helpers/qs.helper';

import Accordion from 'primevue/accordion';
import AccordionTab from 'primevue/accordiontab';
import FilterByShop from '../FilterByShop/FilterByShop.vue';

export default {
  name: 'SideFilters',
  components: {
    Accordion,
    AccordionTab,
    FilterByShop,
  },
  emits: ['on-filter-change'],

  setup() {
    const shops = ref([]);
    const filters = reactive({
      shop: null,
      price: {
        min: 0,
        max: 100,
      },
    });

    onMounted(async () => {
      shops.value = await ShopService.getShops();

      const { shop: { id: shopID } = {}, price = {} } = parseUrlParams();
      // console.log('parseUrlParams', parseUrlParams(), shopID,shops.value.find(shop => shop.id === shopID));
      // filters = {
      //   shop: shops.value.find(shop => shop.id === shopID),
      //   price,
      // };
      filters.shop = shops.value.find(shop => shop.id === shopID);
      console.log(filters, shops.value.find(shop => shop.id === shopID));
    });

    const { shop: selectedShop, price } = toRefs(filters);

    return {
      shops,
      selectedShop,
      price,
      filters,
    };
  },

  methods: {
    onShopSelect: function(shop) {
      updateQuerystringParam(['shop', 'id'], shop.id);
      this.filters.shop = shop;
      this.onFilterChange();
    },
    onPriceChange: function({ min, max }) {
      updateQuerystringParam(['price', 'min'], min);
      updateQuerystringParam(['price', 'max'], max);
      this.filters.price = { min, max };
      this.onFilterChange();
    },
    onFilterChange() {
      this.$emit('on-filter-change', this.filters);
    },
  },
};
</script>
