import { HEADER_HEIGHT } from "../../constant/constants";
import { styled } from "styled-components";

export const shipmentFormStyles = {
  Container: styled.div`
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    flex: 1;
    padding-top: ${HEADER_HEIGHT + 30}px;
    padding-left: 32px;
    padding-right: 32px;
    margin-bottom: 80px;
    flex: 1;
    height: 100%;
  `,
  FormContainer: styled.div<{ $isStretch?: boolean }>`
    padding-top: 24px;
    display: flex;
    flex-direction: column;

    align-items: center;
    gap: 8px;
    width: ${({ $isStretch }) => ($isStretch ? "initial" : "100%")};
    max-width: 650px;

    @media only screen and (max-width: 1024px) {
      flex-direction: column;
      margin-left: auto;
      margin-right: auto;
      & > * {
        width: 100%;
      }
    }
  `,
  InputListContainer: styled.div``,
  InputContainer: styled.div`
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    gap: 8px;
    max-width: 650px;
    flex-wrap: wrap;
    flex: 1;
    & > input {
      display: flex;
      min-width: 120px;
    }
  `,
  TextContainer: styled.div<{ paddingTop: number }>`
    margin-top: ${({ paddingTop }) => paddingTop}px;
    width: 100%;
  `,
};
