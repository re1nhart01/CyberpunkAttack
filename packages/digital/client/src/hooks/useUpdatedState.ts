import type { SetStateAction } from "react";
import { useEffect, useState } from "react";

export const useUpdatedState = <T>(
  defaultValue: T
): [T, React.Dispatch<SetStateAction<T>>] => {
  const [state, setState] = useState(defaultValue);

  useEffect(() => {
    if (state !== defaultValue) {
      setState(defaultValue);
    }

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [defaultValue]);

  return [state, setState];
};
