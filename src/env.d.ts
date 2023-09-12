declare global {
    namespace NodeJS {
        interface ProcessEnv {
            GATSBY_NAME: string;
            GATSBY_FIREBASE_API_KEY: string;
            GATSBY_FIREBASE_AUTH_DOMAIN: string;
            GATSBY_FIREBASE_PROJECT_ID: string;
            GATSBY_FIREBASE_STORAGE_BUCKET: string;
            GATSBY_FIREBASE_MESSAGING_ID: string;
            GATSBY_FIREBASE_APP_ID: string;
            GATSBY_FIREBASE_MEASURING_ID: string;
        }
    }
}

export {};
