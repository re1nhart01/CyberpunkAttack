import * as React from 'react';
import { Images } from '../constants/images';
import SocialButton from '../components/SocialButton/SocialButton';
import { linksDictionary } from '../constants';
import PunkCarousel from '../components/PunkCarousel/PunkCarousel';

const IndexPage = () => {
  return (
    <main className="linker-container">
      <header className="linker-header flex flex-col justify-between items-center mt-[32px] mb-[8px]">
        <p className="header-cyberpunk-next">
          <span className="header-cyberpunk-first">C</span>
          yberpunk
        </p>
        <p className="header-attack-text pl-[20px]">Attack</p>
      </header>
      <PunkCarousel />
      <div className="linker__inner-content">
        <span className="content-header-text">
          Welcome to the
          {' '}
          <span className="blue-variant">Neon City</span>
        </span>
        <span className="content-text whitespace-break-spaces">
          Try the cyber fighting board game in Cyberpunk universe. Choose your side Cyberpunks vs Corporate.
          {'\n\n'}
          8 players, 4 roles, 40+ cyber-implants, 100+ actions (guns, bombs, cyber attacks, mercenaries, etc.)
        </span>
        <img className="linker__inner__versus-img" src={Images.VersusImg} alt="VERSUS_IMG" />
        <div className="w-full flex-col items-center">
          <p className="w-full mb-[16px] join-to-us">
            Join to Us:
          </p>
          <div className="social-button__column">
            <SocialButton {...linksDictionary.instagram} />
            <SocialButton {...linksDictionary.tiktok} />
            <SocialButton {...linksDictionary.discord} />
            <SocialButton {...linksDictionary.telegram} />
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
      <div className="blackout" />
    </main>
  );
};

export default IndexPage;

export const Head = () => <title>Cyberpunk Attack - futuristic team fight in Cyberpunk universe</title>;
