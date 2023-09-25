import styled from "styled-components";
import ImageView from "./index";


export const ImageViewComponents = {
    FbBlurRowImg: styled(ImageView).attrs({
        source: "fbBoxGame",
        alterText: "fbBoxGame"
    })`
      margin-right: 16px;
      max-width:100px;
      width:100%;
    `,
    NovaPoshtaIcon: styled(ImageView).attrs({
        source: "npLogo",
        alterText: "NP_LOGO_IMAGE"
    })`
      width: 24px;
      height: 24px;
    `,
    CardImageView: styled(ImageView).attrs({
        source: "cardView",
        alterText: "NP_LOGO_IMAGE"
    })``
};
