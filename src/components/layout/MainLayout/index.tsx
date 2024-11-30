import React, { FC, PropsWithChildren } from 'react';
import { ThemeProvider } from 'styled-components';
import { I18nextProvider } from 'react-i18next';
import { MainLayoutStyles } from './styles';
import { theme } from '../../../theme/theme';
import i18n from '../../../services/localization/i18n';

type mainLayoutProps = PropsWithChildren<{
    Header?: JSX.Element;
    Footer?:JSX.Element;
}>;

const { ContentWrapper, MainWrapper, ShrinkContainer, FullContainer } = MainLayoutStyles;

const MainLayout: FC<mainLayoutProps> = ({ Footer, Header, children }) => {
  return (
    <I18nextProvider i18n={i18n}>
      <ThemeProvider theme={theme}>
        <MainWrapper>
          { Header || null }
          { children }
        </MainWrapper>
      </ThemeProvider>
    </I18nextProvider>
  );
};

export default MainLayout;
