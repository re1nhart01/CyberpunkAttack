import styled from "styled-components";
import {HEADER_HEIGHT, ICONS} from "../../services/constant/icons";


export const HomePageStyles = {
    Del: styled.div`
      display: flex;
      align-items: center;
      flex: 1;
    `,
    FBlockWrapper: styled.div`
      display: flex;
      background-image: url(${ICONS.headlinerBG});
    `,
    FBlockContainer: styled.div`
      width: 100vw;
      display: flex;
      height: 100vh;
      margin-top: ${HEADER_HEIGHT}px;
      padding-left: 8.3%;
      padding-right: 7.6%;
    `,
    FBlockImage: styled.img`
      object-fit: fill;
    `,
    FHeaderWrapper: styled.div`
      gap: 32px;
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      padding-top: 10%;
      max-width: 625px;
    `,
}