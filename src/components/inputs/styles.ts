import styled from 'styled-components';

export const InputStyles = {
  PrimaryInput: styled.input`
    background-color: rgba(255, 255, 255, 0.1);
    border: 1px solid rgba(255, 255, 255, 0.2);
    border-radius: 60px;
    padding-left: 24px;
    padding-top: 16px;
    padding-bottom: 16px;
    font-family: Zekton,serif;
    font-size: 16px;
    font-weight: 400;
    line-height: 16px;
    text-align: left;
    text-underline-position: from-font;
    text-decoration-skip-ink: none;
    color: ${({ theme }) => theme.colors.white};
    flex: 1;
    
    &::placeholder {
      color: rgba(255, 255, 255, 0.6);
    }

  `,
};
