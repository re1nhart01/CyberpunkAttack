import * as React from 'react';
import type { HeadFC, PageProps } from 'gatsby';
import { useTranslation } from 'react-i18next';
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
import { useEffect } from 'react';
import { service } from '../services';

const {
  FBlockWrapper,
  FBlockContainer,
  FHeaderWrapper,
  Del,
  ContainedRowView,
  VerticalLineWhite2,
  AboutBlockWrapper,
  AboutBlockContent,
  AboutSeparator,
  AboutTheGameWrapper,
  AboutTheGameContainer,
  BorderView,
} = HomePageStyles;
const {
  Text76orbitron700,
  Text18boxed500,
  Text24boxed600,
  Text16boxed500,
  Text16boxed600,
  Text16boxed700,
  Text60orbitron700,
  Text170space700,
} = TypographyComponents;
const {
  FbBlurRowImg,
  CardImageView,
  BackgroundAboutView,
  SeparatorBlack,
  IRLGameView,
  RolesBannerView,
  SocialImage,
} = ImageViewComponents;
const { ShipNovaPoshtaButton, BuyInOneClickButton, SmallSocialButton } = ButtonComponents;
const { Text40orbitron700After } = OverrideTypographyComponents;

const HomePage: React.FC<PageProps> = () => {
  const { t } = useTranslation();

  useEffect(() => {
    service.initServices().then();
  }, []);

  return (
    <MainLayout
      Header={<HeaderView />}
      Footer={<FooterView />}
    >
      <FBlockWrapper>
        <FBlockContainer>
          <FHeaderWrapper>
            <Text76orbitron700 k="firstBlock_HeadingText" />
            <Text18boxed500 k="firstBlock_SubHeading" />
            <BlurRowView>
              <Del>
                <FbBlurRowImg source="fbBoxGame" />
                <RowView rightM={3} gap={8} pos="column">
                  <Text24boxed600 disabledWrap k="gameName" />
                  <Text16boxed500 k="free_ship_ua" />
                </RowView>
              </Del>
              <ShipNovaPoshtaButton
                onPress={() => {}}
                gapInside={10}
              >
                <Text16boxed600 k="buy" />
              </ShipNovaPoshtaButton>
            </BlurRowView>
            <ContainedRowView align="center" dimension="%" gap={6.5} pos="row">
              <RowView align="center" pos="column" gap={4}>
                <Text40orbitron700After k="players_count" />
                <Text16boxed700 k="players" />
              </RowView>
              <VerticalLineWhite2 />
              <RowView align="center" pos="column" gap={4}>
                <Text40orbitron700After k="time_count" />
                <Text16boxed700 k="min_count" />
              </RowView>
              <VerticalLineWhite2 />
              <RowView align="center" pos="column" gap={4}>
                <Text40orbitron700After k="cards_count" />
                <Text16boxed700 k="card" />
              </RowView>
            </ContainedRowView>
          </FHeaderWrapper>
          <CardImageView source="cardView" />
        </FBlockContainer>
      </FBlockWrapper>
      <AboutBlockWrapper>
        <AboutBlockContent>
          <Text60orbitron700 k="game_rules" />
          <AboutSeparator />
          <RowView align="center" pos="row" gap={0}>
            <CardWithText
              leftTextIcon="cardSmallAbout1"
              icon="cardAbout1"
              k="game_mechanics"
              alterBig="CARD_ABOUT_1"
              alterSmall="CARD_SMALL_ABOUT_1"
            />
            <CardWithText
              leftTextIcon="cardSmallAbout2"
              icon="cardAbout2"
              k="role_system"
              alterBig="CARD_ABOUT_2"
              alterSmall="CARD_SMALL_ABOUT_2"
            />
            <CardWithText
              leftTextIcon="cardSmallAbout3"
              icon="cardAbout3"
              k="fight_system"
              alterBig="CARD_ABOUT_3"
              alterSmall="CARD_SMALL_ABOUT_3"
            />
            <CardWithText
              leftTextIcon="cardSmallAbout4"
              icon="cardAbout4"
              k="cyberImplants"
              alterBig="CARD_ABOUT_4"
              alterSmall="CARD_SMALL_ABOUT_4"
            />
            <CardWithText
              leftTextIcon="cardSmallAbout5"
              icon="cardAbout5"
              k="etc"
              alterBig="CARD_ABOUT_5"
              alterSmall="CARD_SMALL_ABOUT_5"
            />
          </RowView>
        </AboutBlockContent>
      </AboutBlockWrapper>
      <SeparatorBlack source="blackSeparator" />
      <AboutTheGameWrapper>
        <AboutTheGameContainer>
          <Text170space700 k="about" />
          <InfoSectionView
            headerText="aboutTheGame"
            subText="spamAboutTheGame"
            contentText="aboutGameContent"
            buttonSection={(
              <BuyInOneClickButton onPress={() => {}}>
                <Text16boxed600 k="buyIn1Click" />
              </BuyInOneClickButton>
            )}
            imageSection={(
              <IRLGameView source="irlGame" />
            )}
          />
        </AboutTheGameContainer>
        <AboutTheGameContainer id="1">
          <Text170space700 k="aiGen" />
          <InfoSectionView
            reversed
            headerText="aboutTheMaterials"
            subText="rolesSeparated"
            contentText="aboutGameContent2"
            buttonSection={(
              <RowView gap={8} pos="row">
                <SmallSocialButton onPress={() => {}}>
                  <SocialImage source="instagram" />
                </SmallSocialButton>
                <SmallSocialButton onPress={() => {}}>
                  <SocialImage source="tiktok" />
                </SmallSocialButton>
                <SmallSocialButton onPress={() => {}}>
                  <SocialImage source="linkedin" />
                </SmallSocialButton>
              </RowView>
            )}
            imageSection={(
              <RolesBannerView source="rolesBanner" />
              )}
          />
        </AboutTheGameContainer>
      </AboutTheGameWrapper>
    </MainLayout>
  );
};

export default HomePage;

export const Head: HeadFC = () => <Html meta="Main Page" title="Cyberpunk Attack" />;
