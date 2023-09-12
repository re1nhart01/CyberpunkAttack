import { Firebase } from "./firebase/Firebase";
import {Localization} from "./localization/Localization";

export class Services {

    private readonly _firebase_module: Firebase;
    private readonly _localization_module: Localization;
    constructor() {
        this._firebase_module = new Firebase(this.firebaseConfig);
        this._localization_module = new Localization();
    }
    public get firebaseConfig() {
        return {
            apiKey: process.env.GATSBY_FIREBASE_API_KEY,
            authDomain: process.env.GATSBY_FIREBASE_AUTH_DOMAIN,
            projectId: process.env.GATSBY_FIREBASE_PROJECT_ID,
            storageBucket: process.env.GATSBY_FIREBASE_STORAGE_BUCKET,
            messagingSenderId: process.env.GATSBY_FIREBASE_MESSAGING_ID,
            appId: process.env.GATSBY_FIREBASE_APP_ID,
            measurementId: process.env.GATSBY_FIREBASE_MEASURING_ID
        };
    }

    public get firebase_module(): Firebase {
        return this._firebase_module;
    }

    public async initServices() {
        await this._localization_module.init();
        await this._firebase_module.init();
    }
}


export const service = new Services();
