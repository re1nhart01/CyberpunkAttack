import * as React from "react"
import type { HeadFC, PageProps } from "gatsby"
import {Html} from "../components/html";
import MainLayout from "../components/layout/MainLayout";
import HeaderView from "../components/layout/Header";
import FooterView from "../components/layout/Footer";
import Typography from "../components/typography";
import {ICONS} from "../services/constant/icons";

const HomePage: React.FC<PageProps> = () => {
  return (
      <MainLayout
          Header={<HeaderView />}
          Footer={<FooterView />}
      >
          <div>
              <img src={ICONS.gatsbyIcon} alt=""/>

          </div>
      </MainLayout>
  )
}

export default HomePage

export const Head: HeadFC = () => <Html meta="Main Page" title="Cyberpunk Attack" />
