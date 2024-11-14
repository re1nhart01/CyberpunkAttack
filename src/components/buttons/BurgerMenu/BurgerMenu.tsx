import React, { useState } from 'react';
import styled from 'styled-components';

const BurgerMenu = ({ setIsOpen, isOpen }: { setIsOpen: React.Dispatch<React.SetStateAction<boolean>>; isOpen: boolean }) => {
  return (
    <Wrapper>
      <Burger onClick={() => setIsOpen((prev) => !prev)} $isOpen={isOpen}>
        <Line $isOpen={isOpen} />
        <Line $isOpen={isOpen} />
        <Line $isOpen={isOpen} />
      </Burger>
    </Wrapper>
  );
};

export default BurgerMenu;

// Styled Components

const Wrapper = styled.div`
  position: relative;
`;

const Burger = styled.div<{ $isOpen: boolean; }>`
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  width: 30px;
  height: 30px;
  cursor: pointer;
`;

const Line = styled.div<{ $isOpen: boolean; }>`
  width: 100%;
  height: 3px;
  background-color: #ffffff;
  transition: transform 0.3s, opacity 0.3s;
  
  &:nth-child(1) {
    transform: ${({ $isOpen }) =>
    ($isOpen ? 'rotate(45deg) translate(7px, 5px)' : 'none')};
  }

  &:nth-child(2) {
    opacity: ${({ $isOpen }) => ($isOpen ? 0 : 1)};
  }

  &:nth-child(3) {
    transform: ${({ $isOpen }) =>
    ($isOpen ? 'rotate(-45deg) translate(9px, -7px)' : 'none')};
  }
`;
