import Typography, {typographyProps} from "./index";
import styled from "styled-components";


export const TypographyComponents = {
    Text22space400: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw400",
        fsz: "fz22",
        ff: "space",
        selector: "span",
    })`
      line-height: 120%;
      letter-spacing: 3.242px;
    `,
    Text11aquantix400: styled(Typography).attrs<typographyProps>({
        color: "purple",
        fw: "fw400",
        fsz: "fz11",
        ff: "aquantix",
        selector: "span",
    })`
      line-height: 100%;
      letter-spacing: 18.526px;
    `,
    // FIRST BLOCK INDEX.TSX
    Text76orbitron700: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw700",
        fsz: "fz76",
        ff: "orbitron",
        selector: "span",
        fzType: "rem",
    })`
      text-shadow: -2px -2px 0px ${({ theme }) => theme.colors.main}, 2px 2px 0px ${({ theme }) => theme.colors.purple};
      line-height: 100%;
      text-transform: uppercase;
    `,
    Text18boxed500: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw500",
        fsz: "fz18",
        ff: "boxedRound",
        selector: "span",
        fzType: "px",
    })`
      line-height: 150%;
      opacity: 0.8;
    `,
    Text24boxed600: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw600",
        fsz: "fz24",
        ff: "boxedRound",
        selector: "span",
        fzType: "px",
    })`
      line-height: 100%;
    `,
    Text16boxed500: styled(Typography).attrs<typographyProps>({
        color: "main",
        fw: "fw500",
        fsz: "fz16",
        ff: "boxedRound",
        selector: "span",
        fzType: "px",
    })`
      line-height: 110%;
    `,
    Text16boxed600: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw600",
        fsz: "fz16",
        ff: "boxedRound",
        selector: "span",
        fzType: "px",
    })``,
    Text40orbitron700: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw700",
        fsz: "fz40",
        ff: "orbitron",
        selector: "span",
        fzType: "px",
    })``,
    Text16boxed700: styled(Typography).attrs<typographyProps>({
        color: "white",
        fw: "fw700",
        fsz: "fz16",
        ff: "boxedRound",
        selector: "span",
        fzType: "px",
    })`
    line-height: 150%;
    opacity: 0.5;
    `
}


export const OverrideTypographyComponents = {
    Text40orbitron700After: styled(TypographyComponents.Text40orbitron700)``
}
