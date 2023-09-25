import styled from "styled-components";
import ButtonView from "../Button";
import {ICONS} from "../../../services/constant/icons";


const HeaderButton = styled(ButtonView)`
      padding: 12px 24px;
  position: relative;
  &:not(:last-child)::after {
    content:"";
    position:absolute;
    top:39%;
    right:0%;
    width:1px;
    z-index: 999;
    background-color: ${({ theme }) => theme.colors.main};
    height:12px;
  }
    `;

export const ButtonRowStyles = {
    Outer: styled.div`
      flex-direction: row;
      display: flex;
    `,
    Wrapper: styled.nav`
      background: url(${ICONS.headerBlurBG}) no-repeat;
      fill: rgba(0, 0, 0, 0.20);
      stroke-width: 1px;
      stroke: var(--Main, #01FFFF);
      backdrop-filter: blur(12px);
      flex-direction: row;
      display: flex;
    `,
    HeaderRightButton: styled(HeaderButton)`
      background: url(${ICONS.headerRightBlurBG}) no-repeat;
      background-position-x: 100%;
    `,
    HeaderButton,
}

