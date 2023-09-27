import React, { FC } from 'react';
import { infoSectionStyles } from './styles';
import { TypographyComponents } from '../../typography/typography.styles';

type infoSectionViewProps = {
    reversed?: boolean;
    headerText: string;
    subText: string;
    contentText: string;
    buttonSection: JSX.Element;
    imageSection: JSX.Element;
}

const { BlackLine, Wrapper, ContentWrapper } = infoSectionStyles;
const { Text18boxed400, Text24bebas400, Text60orbitron800 } = TypographyComponents;

const InfoSectionView: FC<infoSectionViewProps> = ({
  buttonSection,
  imageSection,
  contentText,
  subText,
  headerText,
  reversed }) => {
  return (
    <Wrapper>
      { reversed ? null : imageSection }
      <ContentWrapper reversed={reversed}>
        <Text60orbitron800 k={headerText} />
        <Text24bebas400 k={subText} />
        <BlackLine />
        <Text18boxed400 k={contentText} />
        { buttonSection }
      </ContentWrapper>
      { !reversed ? null : imageSection }
    </Wrapper>
  );
};

export default InfoSectionView;
