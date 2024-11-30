import * as React from 'react';
import { forwardRef, useCallback, useRef, useState } from 'react';
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
  CheckImage,
} = ImageViewComponents;

const {
  PrimaryInput,
} = InputStyles;

const { SubmitFormButton } = ButtonComponents;

type userFormProps = {
    subscribeRef: React.MutableRefObject<HTMLDivElement | null>;
}

export const UserForm = forwardRef<null, userFormProps>(({ subscribeRef }, ref) => {
  const sessionStorage = typeof window !== 'undefined' ? window.sessionStorage : { getItem() {}, setItem() {}, removeItem() {} };
  const { t } = useTranslation();
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');
  const isCompleted = sessionStorage?.getItem('isCompleted');
  const [isLoaded, setIsLoaded] = useState(!!isCompleted);
  const nameInputRef = useRef<HTMLInputElement>(null);
  const emailInputRef = useRef<HTMLInputElement>(null);
  const [isEmailError, setIsEmailError] = useState(false);

  // https://hooks.zapier.com/hooks/catch/20708748/25ngnoi?fullname=zxc&email=asdasdasa&time=123&agent=123121&lang=ua

  const handleSendEmail = useCallback(async () => {
    try {
      setIsEmailError(false);

      const response = {
        fullname: name.trim(),
        email: email.trim(),
        time: new Date().toISOString(),
        agent: window?.navigator?.userAgent,
        lang: window?.navigator?.language,
      };

      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

      if (!emailRegex.test(email)) {
        setIsEmailError(true);
        return;
      }

      let baseUrl = 'https://hooks.zapier.com/hooks/catch/20708748/25ngnoi?';
      let sendpulseAddUser = 'https://red-morning-c59a.attackcyberpunk.workers.dev/?listId=2045';

      let isFirst = true;

      for (const [key, value] of Object.entries(response)) {
        baseUrl += `${!isFirst ? '&' : ''}${key}=${value}`;
        sendpulseAddUser += `&${key}=${value}`;
        isFirst = false;
      }
      await fetch(baseUrl, { method: 'GET', mode: 'no-cors' });

      setIsLoaded(true);
      sessionStorage?.setItem('isCompleted', 'true');

      await fetch(sendpulseAddUser, { method: 'GET', mode: 'no-cors' });
    } catch (e) {
      console.log(e);
    }
  }, [name, email]);

  const handleReload = useCallback(() => {
    setIsLoaded(false);
    setIsEmailError(false);
    sessionStorage?.removeItem('isCompleted');
    setEmail('');
    setName('');
  }, []);

  return (
    <PageSection>
      <PageSectionInner ref={subscribeRef}>
        {isLoaded ? (
          <>
            <CheckImage />
            <InnerWrapper>
              <Text48Orbitron700>{t('section1.thanks')}</Text48Orbitron700>
            </InnerWrapper>
            <InnerWrapper>
              <Text24Zekton400 color="white">{t('section1.thanksText')}</Text24Zekton400>
            </InnerWrapper>
            <FormWrapper $isStretch>
              <SubmitFormButton onPress={handleReload}>
                <Text16Zekton400Black>
                  {t('reload')}
                </Text16Zekton400Black>
              </SubmitFormButton>
            </FormWrapper>
          </>
        ) : (
          <>
            <AlertImage />
            <InnerWrapper>
              <Text48Orbitron700>{t('section1.join')}</Text48Orbitron700>
            </InnerWrapper>
            <InnerWrapper>
              <Text24Zekton400 color="white">{t('section1.signup')}</Text24Zekton400>
            </InnerWrapper>
            <FormWrapper>
              <PrimaryInput ref={nameInputRef} type="text" maxLength={150} value={name} onChange={({ target }) => setName(target.value)} placeholder="Name" />
              <PrimaryInput $isError={isEmailError} ref={emailInputRef} type="text" maxLength={150} value={email} onChange={({ target }) => { setIsEmailError(false); setEmail(target.value); }} placeholder="Email" />
              <SubmitFormButton onPress={handleSendEmail}>
                <Text16Zekton400Black>
                  {t('submit')}
                </Text16Zekton400Black>
              </SubmitFormButton>
            </FormWrapper>
          </>
        )}
      </PageSectionInner>
    </PageSection>
  );
});
