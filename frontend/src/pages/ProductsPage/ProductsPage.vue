<template>
  <div class="products-page">
    <Section>
      <Panel class="p-list">
        <template
          #header
        >
          3 cheapest products
        </template>
        <TopProductsList />
      </Panel>
      <Panel class="p-grid">
        <template
          #header
        >
          <SortByPrice
            @on-price-sort-option-change="onPriceSortOptionChange"
          />
        </template>
        <div class="p-d-flex products-page-content">
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
import { ref, reactive, onMounted } from 'vue';

import { parseUrlParams } from 'src/helpers/qs.helper';
import ProductService from 'src/api/services/product.service';

import Panel from 'primevue/panel';
import ProductsGrid from './components/ProductsGrid/ProductsGrid.vue';
import TopProductsList from './components/TopProductsList/TopProductsList.vue';
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
    TopProductsList,
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
    
    function getProductsSearchOptions() {
      return {
        shop: {
          id: filters?.value?.shop?.id,
        },
        price: {
          min: filters?.value?.price?.min,
          max: filters?.value?.price?.max,
        },
        sort: {
          price: sort?.value?.price,
        },
      };
    };

    async function updateProducts() {
      products.value = await ProductService.getProducts(getProductsSearchOptions());
    };

    onMounted(() => {
      const { shop, price, sort: sortFromUrl } = parseUrlParams();
      filters.value = { shop, price };
      sort.value = sortFromUrl;
      updateProducts(); 
    });

    return {
      products,
      filters,
      sort,
      updateProducts,
      getProductsSearchOptions,
    };
  },

  methods: {
    async onFilterChange(filters) {
      this.filters.value.shop = filters.shop;
      this.filters.value.price = filters.price;
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
