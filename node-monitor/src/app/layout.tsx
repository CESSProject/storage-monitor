import type {Metadata} from "next";
import {Inter} from "next/font/google";
import "./globals.css";
import {ThemeModeScript} from "flowbite-react";
import {LeftDrawer} from "@/app/components/sidebar";
import NavBar from "@/app/components/navbar";
import Footer from "@/app/components/footer";
import React from "react";

const inter = Inter({subsets: ["latin"]});

export const metadata: Metadata = {
    title: "Storage Monitor",
    description: "Monitor and manage CESS miners",
};

export default function RootLayout({children,}: Readonly<{ children: React.ReactNode; }>) {
    return (
        <html lang="en" suppressHydrationWarning>
        <head>
            <link rel="icon" href="/favicon.ico" type="image/x-icon"/>
            <ThemeModeScript/>
        </head>
        <body className={inter.className}>
        <div className="w-screen dark:bg-black">
            <LeftDrawer/>
            <div className="w-full flex-grow p-1 md:overflow-y-auto md:p-1 mb-auto">
                <NavBar/>
                {children}
            </div>
        </div>
        <Footer/>
        </body>
        </html>
    );
}
