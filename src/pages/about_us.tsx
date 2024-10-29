import * as React from 'react';
import type { HeadFC, PageProps } from 'gatsby';
import { useTranslation } from 'react-i18next';
import { useEffect } from 'react';
import { Html } from '../components/html';
import MainLayout from '../components/layout/MainLayout';
import HeaderView from '../components/layout/Header';
import FooterView from '../components/layout/Footer';
import { HomePageStyles } from '../styles/pageStyles/home.styles';
import { OverrideTypographyComponents, TypographyComponents } from '../components/typography/typography.styles';
import BlurRowView from '../components/layout/BlurRow';
import { ImageViewComponents } from '../components/images/ImageView/styles';
import RowView from '../components/layout/RowView';
import { ButtonComponents } from '../components/buttons/Button/components';
import CardWithText from '../components/layout/CardWithText';
import InfoSectionView from '../components/layout/InfoSection';
import { service } from '../services';
import { AboutUsStyles } from '../styles/pageStyles/about_us.styles';

const { Wrapper, InnerWrapper, BlockLeft, BlockRight } = AboutUsStyles;
const { AboutUsImage } = ImageViewComponents;
const { Text18boxed500 } = TypographyComponents;

const AboutUsPage: React.FC<PageProps> = () => {
  const { t } = useTranslation();

  useEffect(() => {
    service.initServices().then();
  }, []);

  return (
    <MainLayout
      Header={<HeaderView />}
      Footer={<FooterView />}
    >
      <Wrapper>
        <BlockLeft>
          <InnerWrapper>
            <AboutUsImage source="igor" />
            <AboutUsImage source="me" />
          </InnerWrapper>
        </BlockLeft>
        <BlockRight>
          <Text18boxed500 k="aboutUs_info" />
        </BlockRight>
      </Wrapper>
    </MainLayout>
  );
};

export default AboutUsPage;

export const Head: HeadFC = () => <Html meta="About US Page" title="Cyberpunk Attack" />;
