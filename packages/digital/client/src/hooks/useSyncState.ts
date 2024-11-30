import { useCallback, useEffect, useState } from "react";

export const useSyncState = (cb: () => void) => {
  const [update, setUpdate] = useState(false);

  const handleUpdate = useCallback((flag = true) => {
    setUpdate(flag);
  }, []);

  useEffect(() => {
    if (update) {
      cb();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [update]);

  return {
    handleUpdate,
  };
};
