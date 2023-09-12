import * as React from "react";
import {FC, PropsWithChildren} from "react";


type seoHtmlProps = PropsWithChildren<{
    title: string;
    meta: string;
}>
export const Html: FC<seoHtmlProps> = ({ meta, title, children }) => {
    return (
        <>
            <meta name="description" content={meta} />
            <title>{ title }</title>
            {children}
        </>
    )
}
