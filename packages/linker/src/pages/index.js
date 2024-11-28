import * as React from 'react';
import { useEffect } from 'react';
import { Images } from '../constants/images';
import SocialButton from '../components/SocialButton/SocialButton';
import { linksDictionary } from '../constants';
import PunkCarousel from '../components/PunkCarousel/PunkCarousel';
import SocialButtonSmall from '../components/SocialButtonSmall/SocialButtonSmall';

const IndexPage = () => {
  typeof window !== 'undefined' && window.location.replace('https://cyberpunkattack.com');

  return (
    <main className="linker-container">
      <header className="linker-header flex flex-col justify-between items-center pt-[32px] mb-[8px] z-[2]">
        <img className="linker-header__logo" src={Images.Logo} alt="LOGO_IMAGE" />
      </header>
      <PunkCarousel />
      <div className="linker__inner-content">
        <div className="linker__inner_text mb-[12px]">
          <p className="content-header-text mb-[12px]">
            Welcome to the
            {' '}
            <span className="blue-variant">Neon City</span>
          </p>
          <span className="content-text whitespace-break-spaces">
            Try the cyber fighting board game in Cyberpunk universe. Choose your side Cyberpunks vs Corporate.
          </span>
        </div>
        <img
          loading="lazy"
          decoding="async"
          className="linker__inner__versus-img mb-[48px]"
          src={Images.VersusImg}
          alt="VERSUS_IMG"
        />
        <div className="w-full flex-col items-center">
          <p className="w-full mb-[16px] join-to-us">
            Follow us:
          </p>
          <div className="social-button__column">
            <SocialButton {...linksDictionary.instagram} />
            <div className="social-button__row">
              <SocialButtonSmall {...linksDictionary.tiktok} />
              <div className="social-button-row__separator" />
              <SocialButtonSmall {...linksDictionary.discord} />
              <div className="social-button-row__separator" />
              <SocialButtonSmall {...linksDictionary.telegram} />
            </div>
          </div>
          <div className="mt-[32px]">
            <p className="w-100 mb-[16px] join-to-us">
              {linksDictionary.email.label}
            </p>
            <p className="social-button-text text-center opacity-50">
              {linksDictionary.email.url}
            </p>
          </div>
        </div>
      </div>
    </main>
  );
};

export default IndexPage;

export const Head = () => <title>Cyberpunk Attack - futuristic team fight in Cyberpunk universe</title>;
