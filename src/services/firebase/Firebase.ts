import { FirebaseApp } from "@firebase/app";
import {initializeApp} from "firebase/app";
import { getAnalytics, Analytics } from "firebase/analytics";
import { ServiceBuilder } from "../service.d";


export class Firebase extends ServiceBuilder{
    private impl: FirebaseApp | null;
    private analyticsImpl: Analytics | null;
    constructor(private config: { [key: string]: string }) {
        super();
        this.impl = null;
        this.analyticsImpl = null;
    }

    public async init() {
        this.impl = initializeApp(this.config);
        this.analyticsImpl = getAnalytics(this.impl);
    }
}
