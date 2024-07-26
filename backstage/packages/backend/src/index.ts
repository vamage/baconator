/*
 * Hi!
 *
 * Note that this is an EXAMPLE Backstage backend. Please check the README.
 *
 * Happy hacking!
 */

import { createBackend } from '@backstage/backend-defaults';
import {
    coreServices,
    createBackendModule,
} from '@backstage/backend-plugin-api';
import { catalogProcessingExtensionPoint } from '@backstage/plugin-catalog-node/alpha';
import { BaconatorProvider } from '../../../plugins/catalog-backend-module-baconator-catalog-connector/src/loader';

export const catalogModuleBaconatorCatalogConnector = createBackendModule({
    pluginId: 'catalog',
    moduleId: 'catalog-backend-module-baconator-catalog-connector',
    register(env) {
        env.registerInit({
            deps: {
                catalog: catalogProcessingExtensionPoint,
                reader: coreServices.urlReader,
                scheduler: coreServices.scheduler,
            },
            async init({catalog, reader, scheduler}) {
                const taskRunner = scheduler.createScheduledTaskRunner({
                    initialDelay: { seconds: 40 },
                    frequency: { minutes: 5 },
                    timeout: { minutes: 10 },
                });
                const baconator = new BaconatorProvider('prod',  reader, taskRunner);
                await catalog.addEntityProvider(baconator);
            },
        });
    },
});


const backend = createBackend();
backend.add(import('@backstage/plugin-app-backend/alpha'));
backend.add(import('@backstage/plugin-proxy-backend/alpha'));
backend.add(import('@backstage/plugin-scaffolder-backend/alpha'));
backend.add(import('@backstage/plugin-techdocs-backend/alpha'));

// auth plugin
backend.add(import('@backstage/plugin-auth-backend'));
// See https://backstage.io/docs/backend-system/building-backends/migrating#the-auth-plugin
backend.add(import('@backstage/plugin-auth-backend-module-guest-provider'));
// See https://backstage.io/docs/auth/guest/provider

// catalog plugin
backend.add(import('@backstage/plugin-catalog-backend/alpha'));
backend.add(
  import('@backstage/plugin-catalog-backend-module-scaffolder-entity-model'),
);

// permission plugin
backend.add(import('@backstage/plugin-permission-backend/alpha'));
backend.add(
  import('@backstage/plugin-permission-backend-module-allow-all-policy'),
);

// search plugin
backend.add(import('@backstage/plugin-search-backend/alpha'));
backend.add(import('@backstage/plugin-search-backend-module-catalog/alpha'));
backend.add(import('@backstage/plugin-search-backend-module-techdocs/alpha'));

backend.add(import('@internal/backstage-plugin-catalog-backend-module-baconator-catalog-connector'));
backend.add(catalogModuleBaconatorCatalogConnector);
backend.start();
