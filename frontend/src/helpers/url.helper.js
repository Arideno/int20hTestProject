export const prependHttpIfNotExists = url => {
  if (!/^https?:\/\//i.test(url)) {
    return 'http://' + url;
  } else {
    return url;
  }
};
