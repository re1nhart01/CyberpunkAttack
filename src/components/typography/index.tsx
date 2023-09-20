import React, { FC, PropsWithChildren, ReactHTML } from 'react';
import { theme } from "../../theme/theme";
import styled from "styled-components";

type ExtensionSelector = Partial<HTMLAnchorElement | HTMLDivElement | HTMLSpanElement> &
    Pick<typographyProps, "fw" | "fsz" | "ff" | "color"> &
    Required<PropsWithChildren>;

export type typographyProps = PropsWithChildren<{
    color?: keyof typeof theme["colors"];
    fw?: keyof typeof theme["fontWeight"];
    fsz?: keyof typeof theme["fontSizes"];
    ff?: keyof typeof theme["fonts"];
    selector?: keyof ReactHTML;
    href?: string;
    className?: string;
}>;


const Typography: FC<typographyProps> = ({
     fsz,
     ff,
     fw,
     color,
     children,
     selector,
     href,
     className,
     }) => {
    const Text = createTextComponent(selector, { fsz, ff, fw, color, href });

    return (
        <Text className={className}>
            { children as never }
        </Text>
    );
};

export function createTextComponent (selector: keyof ReactHTML = "span", {
    fsz,
    color,
    fw,
    ff,
    href }: Omit<typographyProps, "selector" | "children">) {
    return styled[selector as "div"].attrs<ExtensionSelector>((props) => ({
        ...(selector === "a" && { href }),
        ...props,
    }))<ExtensionSelector>
        `
      color: ${({ theme }) => theme.colors[color!] || theme.colors.main};
      font-size: ${({ theme }) => theme.fontSizes[fsz!] || theme.fontSizes.fz14}px;
      font-family: ${({ theme }) => theme.fonts[ff!] || theme.fonts.orbitron};
      font-weight: ${({ theme }) => theme.fontWeight[fw!] || theme.fontWeight.fw400};
    `
}

export default Typography satisfies FC<typographyProps>;
