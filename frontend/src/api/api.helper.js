import axios from 'axios';
import { stringifyParams } from 'src/helpers/qs.helper';
import { prependHttpIfNotExists } from 'src/helpers/url.helper';

const BASE_URL = prependHttpIfNotExists(process.env.VUE_APP_API_URL) || '/v1';

class Api {
  constructor() {
    console.log(BASE_URL)
    this.instance = axios.create({
      baseURL: BASE_URL,
      headers: {
        'Content-Type': 'application/json',
      },
    });
    this.commonHeaders = { 'Content-Type': 'application/json' };
  }

  async get(url, params = {}) {
    const response = this.instance
      .get(`${url}?${stringifyParams(params)}`, {
        headers: this.commonHeaders,
        data: {},
      })
      .then(({ data }) => data)
      .catch(this.handleError);
    return response;
  }

  async post(url, payload) {
    const response = await this.instance
      .post(url, payload, {
        headers: this.commonHeaders,
      })
      .then(({ data }) => data)
      .catch(this.handleError);
    return response;
  }

  async put(url, payload) {
    const response = await this.instance
      .put(url, payload, {
        headers: this.commonHeaders,
      })
      .then(({ data }) => data)
      .catch(this.handleError);
    return response;
  }

  async delete(url) {
    const response = await this.instance
      .delete(url, {
        headers: this.commonHeaders,
      })
      .then(({ data }) => data)
      .catch(this.handleError);
    return response;
  }

  handleError(error) {
    if (error.response) {
      throw new Error(error.response.data.error);
    } else if (error.request) {
      throw new Error(error.request.responseText);
    } else {
      throw new Error(error.message);
    }
  }
}

export default new Api();