import React from "react";
import { LeftDrawer } from "../components/sidebar";
import NavBar from "../components/navbar";
import Footer from "@/app/components/footer";

export default function Layout({ children }: { children: React.ReactNode }) {

  return (
    <div className="w-screen h-screen dark:bg-black">
      <LeftDrawer />
      <div className="w-full flex-grow p-1 md:overflow-y-auto md:p-1">
        <NavBar />
        {children}
      </div>
        {/*<Footer />*/}
    </div>
  );
}
