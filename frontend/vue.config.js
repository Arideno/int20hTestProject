var path = require('path');

module.exports = {
  configureWebpack: (config) => {
    return {
      resolve: {
        alias: {
          src: path.resolve(__dirname, 'src'),
        },
      },
      module: {
        rules: [
          {
            test: /\.m?jsx?$/,
            loader: 'babel-loader',
            options: {
              plugins: ["@babel/plugin-proposal-optional-chaining"]
            },
            include: /node_modules\/primevue/
          }
        ]
      }
    };
  },
};
