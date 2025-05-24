# Casdoor Operator - Documentation

This website is built using [Docusaurus](https://docusaurus.io/), a modern static website generator.

```bash
$ npm run help

docusaurus:
    script: docusaurus
 
start:
    script: docusaurus start
    desciption: Launches a local server for development
    usage: npm start
 
build:
    script: docusaurus build
    desciption: Collects a production version of the site
    usage: npm run build
 
swizzle:
    script: docusaurus swizzle
    desciption: Allows you to change the components of docusaurus (themes, layouts)
    usage: npm run swizzle
 
deploy:
    script: docusaurus deploy
    desciption: Deploys the assembled site (depending on the configuration)
    usage: npm run deploy
 
clear:
    script: docusaurus clear
    desciption: Cleans cache and generated files
    usage: npm run clear
 
serve:
    script: docusaurus serve
    desciption: Launches a local server for testing production-assembly
    usage: npm run serve
 
write-translations:
    script: docusaurus write-translations
    desciption: Creates translation files for internationalization
    usage: npm run write-translations
 
write-heading-ids:
    script: docusaurus write-heading-ids
    desciption: Adds ID to the headings to MDX/MarkDown files
    usage: npm run write-heading-ids
 
typecheck:
    script: tsc
    desciption: Checks Typescript types in the project
    usage: npm run typecheck
 
help:
    script: npm-scripts-help
    desciption: Shows this description of the scripts
    usage: npm run help
```