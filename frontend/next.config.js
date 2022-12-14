module.exports = {
  webpack(config) {
    config.module.rules.push({
      test: /\.svg$/,
      use: [
        {
          loader: '@svgr/webpack',
          options: {
            options: { svgoConfig: { plugins: [{ removeViewBox: false }] } },
          },
        },
      ],
    })

    config.module.rules.push({
      test: /\.(eot|woff|woff2|ttf|svg|png|jpg|gif)$/,
      use: {
        loader: 'url-loader',
        options: {
          limit: 100000,
          name: '[name].[ext]',
        },
      },
    })
    return config
  },
  i18n: {
    locales: ['en', 'th'],
    defaultLocale: 'en'
  },

  basePath: process.env.NEXT_PUBLIC_BASE_URL,
  assetPrefix: process.env.NEXT_PUBLIC_BASE_URL,

  images: {
    domains: ['picsum.photos']
  }
}
