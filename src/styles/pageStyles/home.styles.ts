import styled from 'styled-components';
import { HEADER_HEIGHT, ICONS } from '../../services/constant/icons';
import RowView from '../../components/layout/RowView';

export const HomePageStyles = {
  Del: styled.div`
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-right: 5%;
    `,
  FBlockWrapper: styled.div`
      display: flex;
      background-image: url(${ICONS.headlinerBG});
      background-repeat: no-repeat;
      background-size: cover;
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
};
