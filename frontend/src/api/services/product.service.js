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
    const { Products: products } = await api.get(endpoint, {
      store: filter.shop?.id || 'atb',
    });
    return products.map(ProductService.mapToProductModel);
  }
}
