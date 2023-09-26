import styled from 'styled-components';
import { HEADER_HEIGHT } from '../../../services/constant/icons';

export const HeaderStyles = {
  Wrapper: styled.header`
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      height: ${HEADER_HEIGHT}px;
      padding-left: 3.15%;
      padding-right: 3.15%;
      position: absolute;
      width: 92%;
      flex-wrap: wrap;
      top: 0;
      z-index: 999;
    `,
  LogoWrapper: styled.div`
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      align-items: center;
    `,
  ButtonsWrapper: styled.div`
      
    `,
};
