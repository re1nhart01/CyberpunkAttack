import styled from 'styled-components';

export const infoSectionStyles = {
  Wrapper: styled.div`
    padding-left: 9.3%;
    padding-right: 4.2%;
    padding-top: 90px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    flex-wrap: wrap;
  `,
  ContentWrapper: styled.div<{ reversed?: boolean }>`
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: space-between;
    gap: 25px;
    max-width: ${({ reversed }) => (reversed ? 30 : 44)}%;
  `,
  BlackLine: styled.div`
    width: 100%;
    height: 2px;
    background-color: ${({ theme }) => theme.colors.black};
    position: relative;
    &:after {
      content: "";
      position: absolute;
      width: 31px;
      height: 5px;
      background-color: black;
      bottom: -5px;
      left: 0;
      clip-path: polygon(0 0, 100% 0%, 85% 100%, 0 100%);
    }
  `,
};
