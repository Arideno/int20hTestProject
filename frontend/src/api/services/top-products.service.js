import api from '../api.helper';

const endpoint = '/top';

export default class TopProductsService {
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

  static async getTopProducts() {
    const products = await api.get(endpoint);
    return products.map(TopProductsService.mapToProductModel);
  }
}
