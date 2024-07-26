
import { Entity } from '@backstage/catalog-model';

import {
    EntityProvider,
    EntityProviderConnection,
} from '@backstage/plugin-catalog-node';
import {SchedulerServiceTaskRunner, UrlReaderService} from '@backstage/backend-plugin-api';


/**
 * Provides entities from fictional frobs service.
 */
export class BaconatorProvider implements EntityProvider {
 //   private readonly logger: LoggerService;
    private readonly env: string;
    private readonly reader: UrlReaderService;
    private connection?: EntityProviderConnection;
    private taskRunner: SchedulerServiceTaskRunner;

    /** [1] */
    constructor(
        env: string,
        reader: UrlReaderService,
        taskRunner: SchedulerServiceTaskRunner,
    ) {
        this.env = env;
        this.reader = reader;
        this.taskRunner = taskRunner;

    }

    /** [2] */
    getProviderName(): string {
        return `baconator-${this.env}`;
    }

    /** [3] */
    async connect(connection: EntityProviderConnection): Promise<void> {
        this.connection = connection;
        await this.taskRunner.run({
            id: this.getProviderName(),
            fn: async () => {
                await this.run();
            },
        });
    }

    /** [4] */
    async run(): Promise<void> {
        if (!this.connection) {
            throw new Error('Not initialized');
        }


        const response = await this.reader.readUrl(
            `http://baconator:8081/users`,
        );
        console.info(`Fetched ${ (await response.buffer()).toString()} entities from baconator`);
        let data:  Entity[]  = JSON.parse((await response.buffer()).toString());
      //  this.logger.info(`Fetched ${data} entities from baconator`);

        for (const entity of data) {
            entity.metadata.annotations = {
                ...entity.metadata.annotations,
                'baconator.com/created-by': 'baconator-provider',
            };
        }





        /** [6] */
        await this.connection.applyMutation({
            type: 'full',
            entities: data.map(entity => ({
                entity,
                locationKey: `baconator-provider`,

            })),
        });
    }
}