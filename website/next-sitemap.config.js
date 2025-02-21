const { getIntegrationsPaths } = require("./utils/integrations");

/** @type {import('next-sitemap').IConfig} */
module.exports = {
  siteUrl: "https://www.cloudquery.io",
  generateRobotsTxt: true,
  exclude: [
    "/buy/*",
    "/landing/*",
    "/docs/plugins/sources/_*",
    "/docs/plugins/destinations/_*",
    "/docs/plugins/sources/*/_*",
    "/docs/plugins/destinations/*/_*",
  ],
  robotsTxtOptions: {
    policies: [
      {
        userAgent: '*',
        allow: '/',
        disallow: [
          // disallow buy redirection pages
          '/buy/*',
          // disallow pages that start with underscores
          "/docs/plugins/sources/_*",
          "/docs/plugins/destinations/_*",
          '/docs/plugins/sources/*/_*',
          '/docs/plugins/destinations/*/_*',
        ],
      },
    ],
  },
  additionalPaths: async () => {
    const integrationsData = await getIntegrationsPaths();
    const integrationsPaths = integrationsData.map(({ params: { slug } }) => {
      return { loc: `/integrations/${slug.join("/")}`};
    })
    return integrationsPaths;
  },
};
