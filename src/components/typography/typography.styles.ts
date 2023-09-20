import Typography, {typographyProps} from "./index";
import styled from "styled-components";
import {theme} from "../../theme/theme";
import {ReactHTML} from "react";


export const TypographyComponents = {
    LogoText: styled(Typography).attrs<typographyProps>(({ theme }) => ({
        color: "white",
        fw: "fw400",
        fsz: "fz22",
        ff: "space",
        selector: "span",
    }))`
      line-height: 120%;
      letter-spacing: 3.242px;
    `,
    SubLogoText: styled(Typography).attrs<typographyProps>(({ theme }) => ({
        color: "purple",
        fw: "fw400",
        fsz: "fz11",
        ff: "aquantix",
        selector: "span",
    }))`
      line-height: 100%;
      letter-spacing: 18.526px;
    `,
}
