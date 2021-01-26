<template>
  <div class="products-page">
    <Section>
      <Panel class="p-grid">
        <template
          #header
        >
          <SortByPrice
            @on-price-sort-option-change="onPriceSortOptionChange"
          />
        </template>
        <div class="p-d-flex">
          <SideFilters
            class="p-xl-2"
            @on-filter-change="onFilterChange"
          />
          <ProductsGrid
            class="p-xl-10"
            :products="products"
          />
        </div>
      </Panel>
    </Section>
  </div>
</template>

<script>
import { ref, reactive } from 'vue';

import ProductService from 'src/api/services/product.service';

import Panel from 'primevue/panel';
import ProductsGrid from './components/ProductsGrid/ProductsGrid.vue';
import SideFilters from './components/SideFilters/SideFilters.vue';
import Section from 'src/components/Section/Section.vue';
import SortByPrice from './components/SortByPrice/SortByPrice.vue';

export default {
  name: 'ProductsPage',
  components: {
    ProductsGrid,
    SideFilters,
    Section,
    SortByPrice,
    Panel,
  },
  setup() {
    const products = ref([]);
    const filters = reactive({
      value: {
        shop: null,
        price: null,
      },
    });

    const sort = reactive({
      value: {
        price: null,
      },
    });

    return {
      products,
      filters,
      sort,
    };
  },

  methods: {
    getProductsSearchOptions() {
      return {
        shop: {
          id: this.filters?.value?.shop?.id,
        },
        price: {
          min: this.filters?.value?.price?.min,
          max: this.filters?.value?.price?.max,
        },
        sort: {
          price: this.sort?.value?.price,
        },
      };
    },
    async updateProducts() {
      const products = await ProductService.getProducts(this.getProductsSearchOptions());
      this.products = products;
    },
    async onFilterChange(filters) {
      this.filters.value = filters;
      await this.updateProducts();
    },
    async onPriceSortOptionChange(priceSortOption) {
      this.sort.value = { price: priceSortOption };
      await this.updateProducts();
    },
  },
};
</script>

<style src="./styles.scss" lang="scss" scoped></style>
