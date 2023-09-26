import React, { FC, PropsWithChildren } from 'react';
import { ICONS } from '../../../services/constant/icons';

interface imageViewProps extends PropsWithChildren {
    source: keyof typeof ICONS;
    className?: string;
    alterText?: string;
    append?: this['children'] extends undefined ? undefined : 'left' | 'right';
}

const ImageView: FC<imageViewProps> = ({
  source,
  alterText,
  className,
  children,
  append,
}) => {
  return (
    <>
      { append === 'left' ? children : null }
      <img className={className} src={ICONS[source]} alt={alterText} />
      { append === 'right' ? children : null }
    </>
  );
};

export default ImageView;
