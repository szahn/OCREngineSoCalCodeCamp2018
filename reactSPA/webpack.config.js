const HtmlWebPackPlugin = require("html-webpack-plugin");
const path = require('path');

const htmlPlugin = new HtmlWebPackPlugin({
    template: "./src/index.html",
    filename: "./index.html"
  });

module.exports = {
    output: {
        path: path.join(__dirname, 'wwwroot')
    },
    entry: {
        'app':'./src/index.js'
    },
    module: {
      rules: [
        {
          test: /\.js$/,
          include: [
                path.resolve(__dirname, 'src')
            ],
          exclude: /node_modules/,
          use: {
            loader: "babel-loader"
          }
        },
        {
          test: /src\/\.css$/,
          use: ["style-loader", "css-loader"]
        },
        {
          test: /\.html$/,
          use: [
            {
              loader: "html-loader"
            }
          ]
        }
      ]
    },
    plugins: [htmlPlugin]
  };
  