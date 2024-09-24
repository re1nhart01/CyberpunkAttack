import {
  ChangeEvent,
  ChangeEventHandler,
  DetailedHTMLProps,
  FC,
  InputHTMLAttributes,
} from "react";

type cyberInputProps = {
  onChange: (text: string) => void;
  inputProps: DetailedHTMLProps<
    InputHTMLAttributes<HTMLInputElement>,
    HTMLInputElement
  >;
};

export const FormInputView: FC<cyberInputProps> = ({
  onChange,
  inputProps,
}) => {
  const onChangeInput = (event: ChangeEvent<HTMLInputElement>) => {
    onChange?.(event.currentTarget.value);
  };

  return (
    <div className="h-[64px] w-[70%] text-black">
      <input
        className="h-[64px] w-[100%] p-[24px_12px] gap-[8px] rounded-[8px] bg-[#0000000D] font-orbitron text-[16px] font-medium leading-[16px] text-left"
        {...inputProps}
        onChange={onChangeInput}
        type="text"
      />
    </div>
  );
};
