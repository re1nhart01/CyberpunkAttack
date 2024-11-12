import * as React from 'react';
import { forwardRef, useCallback, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { HomePageStyles } from '../../../styles/pageStyles/home.styles';
import { OverrideTypographyComponents, TypographyComponents } from '../../typography/typography.styles';
import { ImageViewComponents } from '../../images/ImageView/styles';
import { InputStyles } from '../../inputs/styles';
import { ButtonComponents } from '../../buttons/Button/components';

const {
  PageSection,
  PageSectionInner,
  InnerWrapper,
  FormWrapper,
} = HomePageStyles;

const {
  Text48Orbitron700,
  Text24Zekton400,
} = TypographyComponents;

const {
  Text16Zekton400Black,
} = OverrideTypographyComponents;

const {
  AlertImage,
} = ImageViewComponents;

const {
  PrimaryInput,
} = InputStyles;

const { SubmitFormButton } = ButtonComponents;

type userFormProps = {
    subscribeRef: React.MutableRefObject<HTMLDivElement>;
}

export const UserForm = forwardRef<null, userFormProps>(({ subscribeRef }, ref) => {
  const { t } = useTranslation();
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');

  // https://hooks.zapier.com/hooks/catch/20708748/25ngnoi?fullname=zxc&email=asdasdasa&time=123&agent=123121&lang=ua

  const handleSendEmail = useCallback(() => {
    const response = {
      fullname: name,
      email,
      time: new Date().toISOString(),
      agent: window?.navigator?.userAgent,
      lang: window?.navigator?.language,
    };
  }, []);

  return (
    <PageSection>
      <PageSectionInner ref={subscribeRef}>
        <AlertImage />
        <InnerWrapper>
          <Text48Orbitron700>{t('section1.join')}</Text48Orbitron700>
        </InnerWrapper>
        <InnerWrapper>
          <Text24Zekton400>{t('section1.signup')}</Text24Zekton400>
        </InnerWrapper>
        <FormWrapper>
          <PrimaryInput value={name} onChange={({ target }) => setName(target.value)} placeholder="Full name" />
          <PrimaryInput value={email} onChange={({ target }) => setEmail(target.value)} placeholder="Email" />
          <SubmitFormButton onPress={() => {}}>
            <Text16Zekton400Black>
              {t('submit')}
            </Text16Zekton400Black>
          </SubmitFormButton>
        </FormWrapper>
      </PageSectionInner>
    </PageSection>
  );
});
