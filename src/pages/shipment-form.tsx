import { Html } from "../components/html";
import { ImageViewComponents } from "../components/images/ImageView/styles";
import { InputStyles } from "../components/inputs/styles";
import FooterView from "../components/layout/Footer";
import { FullScreenMenuComponent } from "../components/layout/FullScreenMenu/FullScreenMenu";
import HeaderView from "../components/layout/Header";
import MainLayout from "../components/layout/MainLayout";
import Preloader from "../components/layout/Preloader/Preloader";
import {
  OverrideTypographyComponents,
  TypographyComponents,
} from "../components/typography/typography.styles";
import {
  contactUs,
  discordLink,
  instagramLink,
  kickstarter,
} from "../constant/constants";
import { svgs } from "../constant/svgs";
import { service } from "../services";
import { HomePageStyles } from "../styles/pageStyles/home.styles";
import { shipmentFormStyles } from "../styles/pageStyles/shipment-form.styles";
import type { HeadFC, PageProps } from "gatsby";
import * as React from "react";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";

const { Container, FormContainer, InputContainer, TextContainer } =
  shipmentFormStyles;

const {
  FBlockWrapper,
  PageContainer,
  PageSection,
  PageSectionInner,
  Spacer,
  SocialButtonInner,
  Separator,
  KickstarterContainer,
} = HomePageStyles;
const {
  Text24Zekton400To32,
  Text16Zekton400NoColor,
  Text24Zekton700,
  Text56Bangers400,
  Text18Zekton400,
  Text26Space400,
} = TypographyComponents;
const {
  HeaderIllustration,
  SeparatorBlue,
  VersusIllustration,
  CyberpunkText,
  CyberbodyImage,
  CardsList,
  FormImage1,
  FormImage2,
} = ImageViewComponents;
const { PrimaryInput } = InputStyles;

const HomePage: React.FC<PageProps> = () => {
  const { t } = useTranslation();
  const isMobile =
    typeof window !== "undefined" ? window.innerWidth < 1024 : false;
  const [isOpen, setIsOpen] = useState(false);

  const [firstName, setFirstName] = useState("");

  const onScrollIntoView = (
    arg: "subscribe" | "about" | "trailer" | "start"
  ) => {
    switch (arg) {
      case "start":
      case "subscribe":
      case "about":
      case "trailer":
        window.location.href = "/";
        break;
      default:
    }
  };

  const goTo = (url: string) => {
    window.open(url, "_blank");
  };

  useEffect(() => {
    service.initServices().then();
  }, []);

  return (
    <MainLayout>
      <HeaderView
        setIsOpen={setIsOpen}
        isOpen={isOpen}
        onScrollIntoView={onScrollIntoView}
      />
      {isMobile && (
        <FullScreenMenuComponent
          setIsOpen={setIsOpen}
          onScrollIntoView={onScrollIntoView}
          isOpen={isOpen}
        />
      )}
      <Preloader />
      <PageContainer>
        <Container>
          <FormImage1 />
          <FormContainer>
            <TextContainer paddingTop={0}>
              <Text26Space400 color="white">Full Name</Text26Space400>
            </TextContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <TextContainer paddingTop={32}>
              <Text26Space400 color="white">
                Email & Phone Number
              </Text26Space400>
            </TextContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <TextContainer paddingTop={32}>
              <Text26Space400 color="white">Shipping Address</Text26Space400>
            </TextContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
            <InputContainer>
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
              <PrimaryInput
                type="text"
                maxLength={150}
                value={firstName}
                onChange={({ target }) => setFirstName(target.value)}
                placeholder="Name"
              />
            </InputContainer>
          </FormContainer>
          <FormImage2 />
        </Container>
        <Spacer height={100} />
        <FooterView />
      </PageContainer>
    </MainLayout>
  );
};

export default HomePage;

export const Head: HeadFC = () => (
  <Html
    meta="This board game is a cooperative team techno fight game where 2 - 8 players can clash in a battle. Join the battle as a resistance hacker or take control of the corporation as its boss."
    title="Cyberpunk Attack - cooperative team techno fight game"
  />
);
