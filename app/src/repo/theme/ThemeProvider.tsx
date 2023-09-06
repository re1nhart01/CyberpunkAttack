import {createContext, memo, PropsWithChildren, useState} from "react";
import {themes} from "@/repo/theme/theme";

let themeContext = createContext({ theme: themes.dark, setTheme: (v: themes) => {} })



const ThemeProvider = ({ children }: PropsWithChildren<{}>) => {
    let [theme , setTheme] = useState<themes>(themes.dark);

    return (
        <themeContext.Provider value={{ theme , setTheme }}>
            {children}
        </themeContext.Provider>
)
}

export {themeContext};
export default memo(ThemeProvider)
