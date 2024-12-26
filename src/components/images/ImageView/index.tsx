import { ICONS } from "../../../constant/icons";
import React, { FC, PropsWithChildren } from "react";

interface imageViewProps extends PropsWithChildren {
  source?: keyof typeof ICONS;
  className?: string;
  alterText?: string;
  append?: this["children"] extends undefined ? undefined : "left" | "right";
  loading?: "lazy" | "eager";
}

const ImageView: FC<imageViewProps> = ({
  source,
  alterText,
  className,
  children,
  append,
  loading,
}) => {
  return (
    <>
      {append === "left" ? children : null}
      <img
        loading={loading}
        className={className}
        src={ICONS[source as keyof typeof ICONS]}
        alt={alterText}
      />
      {append === "right" ? children : null}
    </>
  );
};

export default ImageView;
