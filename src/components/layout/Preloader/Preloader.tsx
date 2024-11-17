import React, { useEffect, useState } from 'react';
import styled, { keyframes } from 'styled-components';
import { ICONS } from '../../../constant/icons';
import ImageView from '../../images/ImageView';
import { TypographyComponents } from '../../typography/typography.styles';

// Keyframes for spinner animation
const spin = keyframes`
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
`;

const fadeOut = keyframes`
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
  }
`;

const fadeIn = keyframes``;

const { Text26Space400 } = TypographyComponents;

// Styled components
const PreloaderWrapper = styled.div<{ $isExiting?: boolean; }>`
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: #11172A repeat;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 40px;
  z-index: 9999; /* Ensure it's on top of everything */
  animation: ${({ $isExiting }) => ($isExiting ? fadeOut : fadeIn)} 0.5s ease-in-out;
  opacity: ${({ $isExiting }) => ($isExiting ? 0 : 1)};
  pointer-events: ${({ $isExiting }) => ($isExiting ? 'none' : 'auto')};
`;

const Spinner = styled.div`
  animation: ${spin} infinite 3s ease-in-out;
`;

const ImagePreloader = styled(ImageView).attrs({
  source: 'preloader',
  alterText: 'PRELOADER_IMG',
})`
    width: 40vw;
    height: 40vw;
`;

const Preloader = () => {
  const [isLoading, setIsLoading] = useState(true);
  const [isExiting, setIsExiting] = useState(false);

  const subscribe = () => {
    // Trigger the fade-out animation
    const timer = setTimeout(() => {
      setIsExiting(true);
      const innerTimer = setTimeout(() => {
        setIsLoading(false);
        clearTimeout(innerTimer);
      }, 700);
      clearTimeout(timer);
    }, 1500); // Remove preloader after animation
  };

  useEffect(() => {
    subscribe();
  }, []);

  if (!isLoading) return null;

  return (
    <PreloaderWrapper $isExiting={isExiting}>
      <Spinner>
        <ImagePreloader />
      </Spinner>
      <Text26Space400 color="white">
        Loading
      </Text26Space400>
    </PreloaderWrapper>
  );
};

export default Preloader;
