import React, {useEffect} from 'react';
import { useFirebase } from './core/firebase';
import preloader from './assets/preload_sun_Loading_wait_website-512.webp';
import cornerSeparator from "./assets/section seperator.svg"
import { logEvent } from 'firebase/analytics';

function App() {
  const { getRedirectList, redirectList, analytics } = useFirebase();
  const params = new URLSearchParams(window.location.search);
  
  useEffect(() => {
    const redirectValue = params.get("redirect");
    if (redirectList && redirectValue) {
      const existedRedirect = redirectList[redirectValue];
      existedRedirect && setTimeout(() => {
        window.location.replace(existedRedirect);
      }, 1000)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [redirectList])
  
  useEffect(() => {
    getRedirectList().then();
    logEvent(analytics, 'page_view', {
      page_path: window.location.pathname + window.location.search,
      page_title: `Redirect page at time: ${new Date().toISOString()}`
    })
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])
  
  return (
    <div className="app">
      <img className="cornerSeparator-top" src={cornerSeparator} alt="cornerSeparator" />
      <main className="app-header">
        <img src={preloader} className="app-logo" alt="logo" />
        <p className="redirect-text">
          Redirect
        </p>
      </main>
      <img className="cornerSeparator-bottom" src={cornerSeparator} alt="cornerSeparator" />
    </div>
  );
}

export default App;
