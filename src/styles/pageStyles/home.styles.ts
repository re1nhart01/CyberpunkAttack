import styled from "styled-components";
import {HEADER_HEIGHT, ICONS} from "../../services/constant/icons";
import RowView from "../../components/layout/RowView";


export const HomePageStyles = {
    Del: styled.div`
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-right: 5%;
    `,
    FBlockWrapper: styled.div`
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      background: url(${ICONS.headlinerBG}) lightgray 50% / cover no-repeat;
    `,
    FBlockContainer: styled.div`
      width: 100vw;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
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
    ContainedRowView: styled(RowView)`
      width: 100%;
    `,
    VerticalLineWhite2: styled.div`
      width:2px;
      height:40px;
      background-color: ${({ theme }) => theme.colors.white};
    `,
    GradientRowView: styled(RowView)`
      width: 100%;
      height: 160px;
      background: linear-gradient(180deg, #000 0%, rgba(0, 0, 0, 0.00) 100%);
      z-index: 999;
      position: relative;
      top: -220px;
    `,
    AboutBlockWrapper: styled.div`
      width: 100vw;
      height: 100vh;
      background: url(${ICONS.backgroundAbout}) no-repeat;
      background-size: 100% 100%;
    `,
    AboutBlockContent: styled.div`
      padding-left: 4.16%;
      padding-right: 4.16%;
      padding-top: 80px;
      z-index: 10;
    `,
    AboutCarousel: styled.div`
      display: flex;
      flex-direction: row;
      justify-content: space-between;
    `,
    AboutSeparator: styled.div`
      width: 22%;
      height: 1px;
      background-color: ${({ theme }) => theme.colors.white};
      margin-top: 17px;
      margin-bottom: 40px;
    `
}
