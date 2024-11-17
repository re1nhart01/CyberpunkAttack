import type { GatsbyConfig } from 'gatsby';

require('dotenv').config({
  path: '.env',
});

const config: GatsbyConfig = {
  siteMetadata: {
    title: 'Cyberpunk Attack - cooperative team techno fight game',
    siteUrl: 'https://cyberpunkattack.com/',
  },
  graphqlTypegen: false,
  plugins: [
    'gatsby-plugin-styled-components',
    'gatsby-plugin-image',
    'gatsby-plugin-sitemap',
    {
      resolve: 'gatsby-plugin-manifest',
      options: {
        icon: 'src/images/icon.png',
      },
    },
    'gatsby-plugin-mdx',
    'gatsby-plugin-sharp',
    'gatsby-transformer-sharp',
    {
      resolve: 'gatsby-source-filesystem',
      options: {
        name: 'images',
        path: './src/images/',
      },
      __key: 'images',
    },
    {
      resolve: 'gatsby-plugin-preload-fonts',
    },
    {
      resolve: 'gatsby-plugin-offline',
    },
    {
      resolve: 'gatsby-source-filesystem',
      options: {
        name: 'pages',
        path: './src/pages/',
      },
      __key: 'pages',
    },
    {
      resolve: 'gatsby-plugin-purgecss',
      options: {
        printRejected: true, // Показать удаленные стили
        develop: false, // Включить только в продакшн-сборке
      },
    },
    {
      resolve: 'gatsby-plugin-google-gtag',
      options: {
        trackingIds: ['G-LSKCKS8X2J'], // Replace with your GA Measurement ID
        gtagConfig: {
          anonymize_ip: true,
          cookie_expires: 0,
        },
        pluginConfig: {
          head: true,
          respectDNT: true,
        },
      },
    },
  ],
};

export default config;
