import React, { FC, PropsWithChildren, ReactHTML } from 'react';
import styled from 'styled-components';
import { useTranslation } from 'react-i18next';
import { theme } from '../../theme/theme';
import ua from '../../services/localization/ua';

type ExtensionSelector = Partial<HTMLAnchorElement | HTMLDivElement | HTMLSpanElement> &
    Pick<typographyProps, 'fw' | 'fsz' | 'ff' | 'color'> &
    Required<PropsWithChildren>;

type inferredFsz = keyof typeof theme['fontSizes' | 'fontSizesEm'];

export type typographyProps = PropsWithChildren<{
    color?: keyof typeof theme['colors'];
    fw?: keyof typeof theme['fontWeight'];
    fsz?: keyof typeof theme['fontSizes'];
    ff?: keyof typeof theme['fonts'];
    selector?: keyof ReactHTML;
    href?: string;
    fzType?: 'px' | 'rem';
    className?: string;
    k?: keyof typeof ua | string;
    disabledWrap?: boolean;
}>;

const pickFontSize = (th: typeof theme, fzType: 'px' | 'rem' = 'px', fsz: inferredFsz) => {
  return `${fzType === 'px' ? th.fontSizes[fsz!] : th.fontSizesEm[fsz!] || th.fontSizes.fz14}${fzType}`;
};

const Typography: FC<typographyProps> = ({
  fsz,
  ff,
  fw,
  color,
  children,
  selector,
  href,
  fzType,
  className,
  k,
  disabledWrap,
}) => {
  const Text = createTextComponent(selector, { fsz, ff, fw, color, href, fzType, disabledWrap });
  const { t } = useTranslation();
  const entity = k ? t(k) : children;
  return (
    <Text className={className}>
      { entity as never }
    </Text>
  );
};

export function createTextComponent (selector: keyof ReactHTML = 'span', {
  fsz,
  color,
  fw,
  ff,
  fzType,
  href,
  disabledWrap,
}: Omit<typographyProps, 'selector' | 'children'>) {
  return styled[selector as 'div'].attrs<ExtensionSelector>((props) => ({
    ...(selector === 'a' && { href }),
    ...props,
  }))<ExtensionSelector>`
      color: ${({ theme }) => theme.colors[color!] || theme.colors.main};
      font-size: ${({ theme }) => pickFontSize(theme, fzType, fsz as inferredFsz)};
      font-family: ${({ theme }) => theme.fonts[ff!] || theme.fonts.orbitron};
      font-weight: ${({ theme }) => theme.fontWeight[fw!] || theme.fontWeight.fw400};
      white-space: ${disabledWrap ? 'nowrap' : 'pre-wrap'};
      @media only screen and (max-width: 1024px) {
      }
    `;
}

export default Typography satisfies FC<typographyProps>;
