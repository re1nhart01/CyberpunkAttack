import styled from 'styled-components';

export const AboutUsStyles = {
  BlockLeft: styled.div`
    height: 100%;
    max-width: 50vw;
    width: 50vw;
    padding-left: 16px;
    display: flex;
    flex-direction: row;
    align-items: flex-end;
  `,
  BlockRight: styled.div`
  height: 100%;
  max-width: 50vw;
  width: 50vw;
  box-sizing: border-box;
  padding-left: 16px;
  padding-right: 16px;
  padding-top: 60px;
`,
  Wrapper: styled.div`
    padding-top: 120px;
    background-color: #191D33;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
    height: 100%;
    `,
  InnerWrapper: styled.div`
    width: 50%;
    height: 100%;
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
  `,
};
