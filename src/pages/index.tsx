import type { HeadFC, PageProps } from "gatsby";
import * as React from "react";
import { useEffect, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { ButtonComponents } from "../components/buttons/Button/components";
import { SocialButton } from "../components/buttons/SocialButton";
import { UserForm } from "../components/forms/UserForm/UserForm";
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
} = ImageViewComponents;

const HomePage: React.FC<PageProps> = () => {
  const { t } = useTranslation();
  const isMobile =
    typeof window !== "undefined" ? window.innerWidth < 1024 : false;
  const subscribeRef = useRef<HTMLDivElement | null>(null);
  const aboutGameRef = useRef<HTMLDivElement>(null);
  const trailerRef = useRef<HTMLDivElement>(null);
  const [isOpen, setIsOpen] = useState(false);

  const onScrollIntoView = (
    arg: "subscribe" | "about" | "trailer" | "start",
  ) => {
    switch (arg) {
      case "start":
        window.scroll({
          top: 0,
          left: 0,
          behavior: "smooth",
        });
        break;
      case "subscribe":
        subscribeRef.current?.scrollIntoView({
          behavior: "smooth",
          block: "center",
        });
        break;
      case "about":
        aboutGameRef.current?.scrollIntoView({
          behavior: "smooth",
          block: "nearest",
        });
        break;
      case "trailer":
        trailerRef.current?.scrollIntoView({
          behavior: "smooth",
          block: "center",
        });
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
        <FBlockWrapper>
          <HeaderIllustration isMobile={isMobile} />
          <KickstarterContainer>
            <div id="kickstarter-block">
              <div className="kick-content">
                <Text24Zekton400To32 color="white">
                  {t("cyberpunk")}
                </Text24Zekton400To32>
                <Text16Zekton400NoColor color="main">
                  {t("available")}
                </Text16Zekton400NoColor>
              </div>
            </div>
            <div style={{ boxSizing: "border-box" }}>
              <SocialButton goTo={() => goTo(kickstarter)}>
                <Text24Zekton700>{t("getNow")}</Text24Zekton700>
              </SocialButton>
            </div>
          </KickstarterContainer>
        </FBlockWrapper>
        <UserForm subscribeRef={subscribeRef} />
        <Spacer height={isMobile ? 88 : 120} />
        <PageSection ref={aboutGameRef}>
          <PageSectionInner>
            <Text56Bangers400 color="rgbaw09">
              {t("diveInto")}
              {" "}
              <Text56Bangers400 color="main">
                {t("cyberFuture")}
              </Text56Bangers400>
            </Text56Bangers400>
            <SeparatorBlue />
            <Spacer height={34} />
            <Text18Zekton400 color="rgbaw09">{t("firstText")}</Text18Zekton400>
            <Spacer height={34} />
            <Text18Zekton400 color="main">
              {t("gameMechanics")}
              {" "}
              <Text18Zekton400 color="rgbaw09">
                {t("gameMechanicsText")}
              </Text18Zekton400>
            </Text18Zekton400>
            <Spacer height={34} />
            <VersusIllustration loading="lazy" />
            <Spacer height={34} />
            <Text56Bangers400 color="rgbaw09">
              {t("about")}
              {" "}
              <Text56Bangers400 color="main">{t("lore")}</Text56Bangers400>
            </Text56Bangers400>
            <Spacer height={34} />
            <Text18Zekton400 color="rgbaw09">
              {t("aboutLoreText")}
            </Text18Zekton400>
            <Spacer height={67} />
            <CyberpunkText />
            <Spacer height={43} />
            <CyberbodyImage loading="lazy" />
            <Spacer height={34} />
            <Text26Space400>
              {t("join")}
              {" "}
              <Text26Space400 color="white">{t("toTheBattle")}</Text26Space400>
            </Text26Space400>
            <Spacer height={16} />
            <CardsList loading="lazy" />
            <Spacer height={34} />
          </PageSectionInner>
        </PageSection>
        <PageSection ref={trailerRef}>
          <PageSectionInner>
            <iframe
              loading="lazy"
              width="100%"
              style={{ width: "100%", aspectRatio: 16 / 9 }}
              src="https://www.youtube.com/embed/W4agXL6aVow?si=mMg4nYVPUDCl3PHR"
              title="YouTube video player"
              frameBorder="0"
              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
              referrerPolicy="strict-origin-when-cross-origin"
              allowFullScreen
            />
          </PageSectionInner>
        </PageSection>
        <Spacer height={40} />
        <PageSection>
          <PageSectionInner>
            <Separator />
            <SocialButtonInner>
              <SocialButton goTo={() => goTo(instagramLink)}>
                <svgs.instagram />
                <Text24Zekton700>{t("socials.instagram")}</Text24Zekton700>
              </SocialButton>
              <SocialButton goTo={() => goTo(discordLink)}>
                <svgs.discord />
                <Text24Zekton700>{t("socials.discord")}</Text24Zekton700>
              </SocialButton>
              <SocialButton goTo={() => goTo(contactUs)}>
                <svgs.contactUs />
                <Text24Zekton700>{t("socials.contactUs")}</Text24Zekton700>
              </SocialButton>
            </SocialButtonInner>
          </PageSectionInner>
        </PageSection>
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
