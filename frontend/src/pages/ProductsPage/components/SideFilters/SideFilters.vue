<template>
  <div class="side-filters p-d-flex p-flex-column">
    <Accordion
      :multiple="true"
      :active-index="[0]"
    >
      <AccordionTab header="Shop">
        <FilterByShop
          :shops="shops"
          :selected-shops="selectedShops"
          :on-shop-select="onShopSelect"
        />
      </AccordionTab>
    </Accordion>
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';

import ShopService from 'src/api/services/shop.service';

import Accordion from 'primevue/accordion';
import AccordionTab from 'primevue/accordiontab';
import FilterByShop from '../FilterByShop/FilterByShop.vue';

export default {
  name: 'SideFilters',
  components: {
    Accordion,
    AccordionTab,
    FilterByShop
  },

  setup() {
    let shops = ref([]);

    onMounted(async () => {
      shops.value = await ShopService.getShops();
    });

    return {
      shops,
      selectedShops: [],
      onShopSelect: () => {},
    };
  }
};
</script>
