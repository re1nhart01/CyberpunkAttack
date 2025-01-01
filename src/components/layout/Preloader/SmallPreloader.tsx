import React from "react";
import styled, { keyframes } from "styled-components";

// Keyframes for animation
const spin = keyframes`
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
`;

// Styled Preloader Container
const PreloaderContainer = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh; /* Full screen height */
  width: 100vw;  /* Full screen width */
  background-color: rgba(0, 0, 0, 0.8); /* Semi-transparent background */
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9999; /* Ensure it appears above all content */
`;

// Styled Loader (Spinner)
const Spinner = styled.div`
  width: 20px;
  height: 20px;
  border: 4px solid rgba(255, 255, 255, 0.3); /* Light border */
  border-top: 4px solid ${({ theme }) => theme.colors.dark}; /* White top border */
  border-radius: 50%; /* Circular shape */
  animation: ${spin} 1s linear infinite; /* Infinite rotation */
`;

const SmallPreloader = () => {
  return (
    <Spinner />
  );
};

export default SmallPreloader;
