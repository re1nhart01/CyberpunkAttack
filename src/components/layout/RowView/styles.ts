import styled from "styled-components";
import {rowViewProps} from "./index";


export const RowViewStyleComponent = styled.div<Omit<rowViewProps, "children" | "className">>`
  display: flex;
  flex-direction: ${({ pos }) => pos};
  gap: ${({ gap }) => gap}${({ dimension }) => dimension || "px"};
  margin-left: ${({ leftM }) => leftM || 0}%;
  margin-right: ${({ rightM }) => rightM || 0}%;
  justify-content: ${({ justify }) => justify || "flex-start"};
  align-items: ${({ align }) => align || "flex-start"};
`
