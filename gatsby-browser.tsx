import * as React from "react"
import type { GatsbyBrowser } from "gatsby"
import {useEffect} from "react";
import {I18nextProvider} from "react-i18next";
import styled, {ThemeProvider} from "styled-components";
//CSS
import "./src/styles/normalize.css";
import "./src/styles/fonts.css";
import "./src/styles/main.css";

import {theme} from "./src/theme/theme";
import { service } from "./src/services";
import i18n from "./src/services/localization/i18n";


const { ApplicationWrapper } = {
    ApplicationWrapper: styled.div`
      min-height: 100%;
    `,
}

export const wrapPageElement: GatsbyBrowser["wrapPageElement"] = ({
    element,
  }) => {
    console.log('zxc2');

    useEffect(() => {
       service.initServices().then();
    }, [])

    return (
    <I18nextProvider i18n={i18n}>
        <ThemeProvider theme={theme}>
                {element}
        </ThemeProvider>
    </I18nextProvider>
    )
}


