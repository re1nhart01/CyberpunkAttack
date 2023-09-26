import React, { FC, PropsWithChildren } from 'react';
import { MainLayoutStyles } from './styles';

type mainLayoutProps = PropsWithChildren<{
    Header?: JSX.Element;
    Footer?:JSX.Element;
}>;

const { ContentWrapper, MainWrapper, ShrinkContainer, FullContainer } = MainLayoutStyles;

const MainLayout: FC<mainLayoutProps> = ({ Footer, Header, children }) => {
  return (
    <MainWrapper>
      { Header || null }
      <ContentWrapper>
        <FullContainer>
          { children }
        </FullContainer>
      </ContentWrapper>
      { Footer || null }
    </MainWrapper>
  );
};

export default MainLayout;
