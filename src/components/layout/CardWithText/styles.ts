import styled from 'styled-components';
import ImageView from '../../images/ImageView';

export const cardWithTextStyles = {
  Wrapper: styled.div`
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      justify-content: space-between;
      gap: 24px;
    `,
  Image: styled(ImageView).attrs({})`
      width: 100%;
      height: 100%;
    `,
  SmallImage: styled(ImageView).attrs({})`
      width: 32px;  
      height: 32px;
    `,
  RedLine: styled.div`
      width: 64px;
      height: 1px;
      background-color: ${({ theme }) => theme.colors.purple};
    `,
};
