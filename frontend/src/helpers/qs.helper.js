import qs from 'qs';

export const stringifyParams = params =>
  qs.stringify(params, {
    arrayFormat: 'comma',
    encodeValuesOnly: true,
    encode: false,
  });

export const parseUrlParams = params => {
  const parsed = qs.parse(params, {
    comma: true,
    ignoreQueryPrefix: true,
    decoder(value) {
      if (/^(\d+|\d*\.\d+)$/.test(value)) {
        return parseFloat(value);
      }

      let keywords = {
        true: true,
        false: false,
        null: null,
        undefined: undefined,
      };
      if (value in keywords) {
        return keywords[value];
      }

      return value;
    },
  });
  return parsed;
};

export const updateQuerystringParams = params => {
  const newUrlParams = stringifyParams(params);
  window.history.replaceState(null, '', `?${newUrlParams}`);
};

export const updateQuerystringParam = (key, value) => {
  const oldParams = parseUrlParams(window.location.search);
  const newUrlParams = stringifyParams({ ...oldParams, [key]: value });
  window.history.replaceState(null, '', `?${newUrlParams}`);
};
