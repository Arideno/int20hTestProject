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

  static async getProducts(filter = {}) {
    const minPrice = filter.price?.min;
    const maxPrice = filter.price?.max;
    const { Products: products } = await api.get(endpoint, {
      store: filter.shop?.id || 'atb',
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

    return products.map(ProductService.mapToProductModel).filter(isPriceValid);
  }
}
