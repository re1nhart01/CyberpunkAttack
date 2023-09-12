
export abstract class ServiceBuilder {
    protected constructor() {}

    public abstract init(): Promise<void>
}
