import React, { FC } from 'react';
import { ICONS } from '../../../constant/icons';
import { cardWithTextStyles } from './styles';
import RowView from '../RowView';
import { TypographyComponents } from '../../typography/typography.styles';

type cardWithTextViewProps = {
    k: string;
    icon: keyof typeof ICONS;
    leftTextIcon: keyof typeof ICONS;
    alterSmall: string;
    alterBig: string;
}

const { Wrapper, Image, SmallImage, RedLine } = cardWithTextStyles;
const { Text20boxed400 } = TypographyComponents;

const CardWithTextView: FC<cardWithTextViewProps> = ({
  alterBig,
  alterSmall,
  leftTextIcon,
  icon,
  k }) => {
  return (
    <Wrapper>
      <Image alterText={alterBig} source={icon} />
      <RowView align="center" pos="row" dimension="px" gap={4}>
        <SmallImage source={leftTextIcon} alterText={alterSmall} />
        <Text20boxed400 k={k} />
      </RowView>
      <RedLine />
    </Wrapper>
  );
};

export default CardWithTextView;
