import './globals.css'
import { Inter } from 'next/font/google'
import { ComponentLayout } from "@/components/layouts/ComponentLayout";
import React, { PropsWithChildren } from "react";
import {Metadata} from "next";

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
    title: 'Cyberpunk Attack',
    description: 'Desk game',
    authors: [{ name: "re1nhart", url: "https://localhost:8080" }]
}
export default function RootLayout({
    children,
   }: PropsWithChildren<{}>) {
  return (
  <html lang="en">
    <body>
        {children}
    </body>
  </html>
  )
}
