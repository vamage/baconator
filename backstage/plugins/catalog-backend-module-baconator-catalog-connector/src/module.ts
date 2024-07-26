import {
  coreServices,
  createBackendModule,
} from '@backstage/backend-plugin-api';

export const catalogModuleBaconatorCatalogConnector = createBackendModule({
  pluginId: 'catalog',
  moduleId: 'baconator-catalog-connector',
  register(reg) {
    reg.registerInit({
      deps: { logger: coreServices.logger },
      async init({ logger }) {
        logger.info('Hello World!');
      },
    });
  },
});
