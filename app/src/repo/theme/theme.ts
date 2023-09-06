import {useCallback, useContext} from "react";
import {themeContext} from "@/repo/theme/ThemeProvider";


export enum themes {
    dark,
    light,
}


export const useTheme = () => {
    const {theme, setTheme} = useContext(themeContext);

    const updateTheme = useCallback((v: themes) => {
        if (theme === v) return;
        setTheme(v);
    }, [theme])

    return {
        theme,
        updateTheme,
    }
}
