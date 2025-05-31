import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const organizationName = "ptrvsrg";
const projectName = "casdoor-operator";
const baseUrl = `/${projectName}/`;

const config: Config = {
    title: 'Casdoor Operator',
    favicon: 'img/favicon.ico',
    tagline: 'A Kubernetes Operator for Casdoor',
    projectName,
    organizationName,
    baseUrl,
    url: `https://${organizationName}.github.io`,
    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'warn',
    i18n: {
        defaultLocale: 'en',
        locales: ['en'],
    },
    future: {
        experimental_storage: {
            type: 'localStorage',
            namespace: 'casdoor-operator',
        }
    },
    presets: [
        [
            'classic',
            {
                docs: {
                    sidebarPath: './sidebars.ts',
                    editUrl: `https://github.com/${organizationName}/${projectName}/tree/main/`,
                },
                theme: {
                    customCss: './src/css/custom.css',
                },
                svgr: {
                    svgrConfig: {},
                },
                sitemap: {
                    lastmod: 'datetime',
                    changefreq: 'weekly',
                    priority: 0.5,
                    ignorePatterns: [],
                    filename: 'sitemap.xml',
                },
            } satisfies Preset.Options,
        ],
    ],

    themes: ['@docusaurus/theme-live-codeblock'],

    themeConfig: {
        image: 'img/docusaurus-social-card.jpg',
        navbar: {
            title: 'Casdoor Operator',
            logo: {
                alt: 'Casdoor Operator Logo',
                src: 'img/logo.svg',
            },
            items: [
                {
                    type: 'docSidebar',
                    sidebarId: 'docSidebar',
                    position: 'left',
                    label: 'Documentation',
                },
                {
                    href: 'https://github.com/ptrvsrg/casdoor-operator',
                    label: 'GitHub',
                    position: 'right',
                },
            ],
        },
        footer: {
            links: [
                {
                    title: 'Learn about Casdoor Operator',
                    items: [
                        {
                            label: 'GitHub',
                            href: 'https://github.com/ptrvsrg/casdoor-operator',
                        },
                    ]
                },
                {
                    title: 'Learn about Casdoor',
                    items: [
                        {
                            label: 'Documentation',
                            to: 'https://casdoor.org/docs/overview',
                        },
                        {
                            label: 'Demo',
                            href: 'https://demo.casdoor.com',
                        },
                        {
                            label: 'GitHub',
                            href: 'https://github.com/casdoor/casdoor',
                        },
                    ]
                },
                {
                    title: 'Learn about Operator SDK',
                    items: [
                        {
                            label: 'Documentation',
                            to: 'https://sdk.operatorframework.io/docs',
                        },
                        {
                            label: 'GitHub',
                            href: 'https://github.com/operator-framework/operator-sdk',
                        },
                    ]
                }
            ],
            logo: {
                alt: 'Casdoor Operator Logo',
                src: 'img/logo.svg',
                width: 30,
                href: '/'
            },
            copyright: `Copyright © ${new Date().getFullYear()} ptrvsrg. Built with Docusaurus.`,
        },
        prism: {
            theme: prismThemes.github,
            darkTheme: prismThemes.oneDark,
        },
        plugins: [
            [require.resolve("@cmfcmf/docusaurus-search-local"), {
                indexPages: true,
                includeParentCategoriesInPageTitle: true,
            }],
        ]
    } satisfies Preset.ThemeConfig,
};

export default config;
