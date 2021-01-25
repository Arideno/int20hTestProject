export default class ShopService {
  static async getShops() {
    return [
      {
        name: 'ATB',
        id: 'atb',
      },
      {
        name: 'Silpo',
        id: 'silpo',
      },
      {
        name: 'Auchan',
        id: 'auchan',
      },
    ];
  }
}
