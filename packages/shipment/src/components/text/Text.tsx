import { FC } from "react";

type textProps = {
  text: string;
  className: string;
};

export const TextView: FC<textProps> = ({ text, className }) => {
  return <div className={className}>{text}</div>;
};
