import { HEADER_HEIGHT } from "../../constant/constants";
import { styled } from "styled-components";

export const shipmentFormStyles = {
  Container: styled.div`
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    flex: 1;
    padding-top: ${HEADER_HEIGHT + 30}px;
    padding-left: 32px;
    padding-right: 32px;
    padding-bottom: 80px;
  `,
  FormContainer: styled.div<{ $isStretch?: boolean }>`
    padding-top: 24px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    width: ${({ $isStretch }) => ($isStretch ? "initial" : "100%")};
    flex-wrap: wrap;
    max-width: 650px;

    & > div span {
      align-self: start;
      padding-left: 10px;
    }

    @media only screen and (max-width: 1024px) {
      width: ${({ $isStretch }) => ($isStretch ? "initial" : "100%")};
      flex-direction: column;
      & > * {
        box-sizing: border-box;
        width: 100%;
      }
    }
  `,
  InputContainer: styled.div`
    max-height: 53px;
    width: 100%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 8px;
    & > input {
    }
  `,
  TextContainer: styled.div<{ paddingTop: number }>`
    margin-bottom: 16px;
    margin-top: ${({ paddingTop }) => paddingTop}px;
    width: 100%;
  `,
};
