// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getFirestore, collection, getDocs } from "firebase/firestore"
import { firebaseConfig } from "../config/config";
import { useState } from "react";
import { REDIRECT_KEY, redirectItemType } from "./firebase.type";

const app = initializeApp(firebaseConfig);

export const analytics = getAnalytics(app);
export const db = getFirestore(app);


export const useFirebase = () => {
    const [redirectList, setRedirectList] = useState<redirectItemType | null>(null);
    async function getRedirectList() {
        try {
            const redirectListRef = collection(db, "redirect-list");
            const querySnapshot = await getDocs(redirectListRef);
            querySnapshot.forEach((doc) => {
                if (doc.id === REDIRECT_KEY) {
                    const data = doc.data() as redirectItemType;
                    if (data) {
                        setRedirectList(data)
                    }
                }
            });
        } catch (error) {
            console.error("Error reading redirect-list collection: ", error);
        }
    }
    
    return {
        getRedirectList,
        redirectList,
        analytics,
        db,
        app,
    }
}