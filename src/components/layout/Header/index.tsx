import React, { FC } from "react";
import { useTranslation } from "react-i18next";
import { ShrinkContainer } from "../MainLayout/styles";

import { HeaderStyles } from "./styles";
import { TypographyComponents } from "../../typography/typography.styles";
import ButtonRowView from "../../buttons/ButtonsRow";
import { ImageViewComponents } from "../../images/ImageView/styles";
import { donateLink, instagramLink } from "../../../constant/constants";
import BurgerMenu from "../../buttons/BurgerMenu/BurgerMenu";

type headerViewProps = {
  onScrollIntoView: (v: "subscribe" | "about" | "trailer" | "start") => void;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  isOpen: boolean;
  hideButtons?: boolean;
};

const { Wrapper, LogoWrapper, ButtonsWrapper, HeaderButton, Separator, HeaderButtonRound } = HeaderStyles;

const { Text16Zekton700 } = TypographyComponents;
const { LaidlonLogo, CALogo, InstaLogo } = ImageViewComponents;

const HeaderView: FC<headerViewProps> = ({ onScrollIntoView, hideButtons, setIsOpen, isOpen }) => {
  const { t } = useTranslation();
  const isMobile = typeof window !== "undefined" ? window.innerWidth < 1024 : false;
  const onDonatePress = () => {
    window.location.href = donateLink;
  };

  const onInstagramPress = () => {
    window.location.href = instagramLink;
  };

  const onGoToDelivery = () => {
    window.open("/shipment-form", "_blank", "rel=noopener noreferrer");
  };

  return (
    <Wrapper>
      <LogoWrapper onClick={() => onScrollIntoView("start")}>
        <LaidlonLogo />
        <CALogo />
      </LogoWrapper>
      {isMobile ? <BurgerMenu setIsOpen={setIsOpen} isOpen={isOpen} /> : (
        <ButtonsWrapper>
          {!hideButtons && (
          <>
            <HeaderButton onPress={() => onScrollIntoView("subscribe")}>
              <Text16Zekton700>{t("header.subscribe")}</Text16Zekton700>
            </HeaderButton>
            <Separator />
            <HeaderButton onPress={() => onScrollIntoView("about")}>
              <Text16Zekton700>{t("header.about")}</Text16Zekton700>
            </HeaderButton>
            <Separator />
            <HeaderButton onPress={() => onScrollIntoView("trailer")}>
              <Text16Zekton700>{t("header.trailer")}</Text16Zekton700>
            </HeaderButton>
            <Separator />
          </>
          )}
          <HeaderButton onPress={onDonatePress}>
            <Text16Zekton700>{t("header.donate")}</Text16Zekton700>
          </HeaderButton>
          <Separator />
          <HeaderButton onPress={onInstagramPress}>
            <InstaLogo />
          </HeaderButton>
          <Separator />
          <HeaderButtonRound onPress={() => onGoToDelivery()}>
            <Text16Zekton700 color="black">{t("header.delivery")}</Text16Zekton700>
          </HeaderButtonRound>
        </ButtonsWrapper>
      )}
    </Wrapper>
  );
};

export default HeaderView;
