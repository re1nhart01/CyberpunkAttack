import * as React from "react"
import type { HeadFC, PageProps } from "gatsby"
import {Html} from "../components/html";
import MainLayout from "../components/layout/MainLayout";
import HeaderView from "../components/layout/Header";
import FooterView from "../components/layout/Footer";
import {HomePageStyles} from "../styles/pageStyles/home.styles";
import {TypographyComponents} from "../components/typography/typography.styles";
import {useTranslation} from "react-i18next";
import BlurRowView from "../components/layout/BlurRow";
import {ImageViewComponents} from "../components/images/ImageView/styles";
import RowView from "../components/layout/RowView";
import { ButtonComponents } from "../components/buttons/Button/components";


const { FBlockWrapper, FBlockContainer, FHeaderWrapper, Del } = HomePageStyles;
const { Text76orbitron700, Text18boxed500, Text24boxed600, Text16boxed500 } = TypographyComponents;
const { FbBlurRowImg } = ImageViewComponents;
const { ShipNovaPoshtaButton } = ButtonComponents;

const HomePage: React.FC<PageProps> = () => {
    const { t } = useTranslation();
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
                              <Text24boxed600 k="gameName" />
                              <Text16boxed500 k="free_ship_ua" />
                          </RowView>
                          </Del>
                          <ShipNovaPoshtaButton
                              onPress={() => {}}
                              text={t("buy")}
                          />
                      </BlurRowView>
                  </FHeaderWrapper>
              </FBlockContainer>
          </FBlockWrapper>
      </MainLayout>
  )
}

export default HomePage

export const Head: HeadFC = () => <Html meta="Main Page" title="Cyberpunk Attack" />
