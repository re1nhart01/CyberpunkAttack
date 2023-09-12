import React, { FC, PropsWithChildren, ReactHTML } from 'react';
import { theme } from "../../theme/theme";
import styled from "styled-components";


type typographyProps = PropsWithChildren<{
    color: keyof typeof theme["colors"];
    fw: keyof typeof theme["fontWeight"];
    fsz: keyof typeof theme["fontSizes"];
    ff: keyof typeof theme["fonts"];
    selector: keyof ReactHTML;
    href?: string;
}>;
const Typography: FC<typographyProps> = ({
     fsz,
     ff,
     fw,
     color,
     children,
     selector,
     href,
     }) => {
    const Text = createTextComponent(selector);
    return (
        <Text
            href={href}
            ff={ff}
            fsz={fsz}
            fw={fw}
            color={color}
        >
            { children as never }
        </Text>
    );
};

const createTextComponent = (selector: keyof ReactHTML) => {
    return styled[selector as "div"]<
        Partial<HTMLAnchorElement | HTMLDivElement | HTMLSpanElement> &
        Pick<typographyProps, "fw" | "fsz" | "ff" | "color"> &
        Required<PropsWithChildren>>`
      color: ${({ color = "main", theme }) =>
              theme.colors[color]};
      font-size: ${({ fsz = "fz14", theme }) => theme.fontSizes[fsz]}px;
      font-family: ${({ ff = "orbitron", theme }) => theme.fonts[ff]};
      font-weight: ${({ fw = "fw400", theme }) => theme.fontWeight[fw]};
    `
}

export default Typography satisfies FC<typographyProps>;
