import React, { FC, PropsWithChildren } from 'react';
import { RowViewStyleComponent } from './styles';

export type rowViewProps = PropsWithChildren<{
    pos: 'row' | 'column';
    className?: string;
    gap: number | string;
    leftM?: number | string;
    rightM?: number | string;
    align?: string;
    justify?: string;
    dimension?: 'px' | '%';
}>;

const RowView: FC<rowViewProps> = ({
  leftM,
  rightM,
  dimension,
  gap,
  pos,
  className,
  children,
  justify,
  align,
}) => {
  return (
    <RowViewStyleComponent
      leftM={leftM}
      rightM={rightM}
      gap={gap}
      className={className}
      pos={pos}
      justify={justify}
      align={align}
      dimension={dimension}
    >
      { children }
    </RowViewStyleComponent>
  );
};

export default RowView;
