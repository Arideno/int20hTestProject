import api from '../api.helper';

const endpoint = '/info';

export default class ProductService {
  constructor() {}

  static mapToProductModel(serverResponseModel) {
    const { Name, Price, Image, Url, Store } = serverResponseModel;
    return {
      name: Name,
      price: Price,
      image: Image,
      url: Url,
      store: Store,
    };
  }

  static async getProducts(options = {}) {
    const minPrice = options.price?.min;
    const maxPrice = options.price?.max;
    const priceSortOption = options.sort?.price;
    const { Products: products } = await api.get(endpoint, {
      store: options.shop?.id || 'atb',
    });

    const isPriceValid = product => {
      if (minPrice && product.price < minPrice) {
        return false;
      }
      if (maxPrice && product.price > maxPrice) {
        return false;
      }
      return true;
    };

    return products
      .map(ProductService.mapToProductModel)
      .filter(isPriceValid)
      .sort((product1, product2) => {
        if (priceSortOption === 'asc') return product1.price - product2.price;
        else return product2.price - product1.price;
      });
  }
}
