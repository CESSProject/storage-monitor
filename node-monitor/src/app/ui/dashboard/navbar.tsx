"use client";
import { DarkThemeToggle, Navbar } from "flowbite-react";
import Image from "next/image";
import Link from "next/link";
import { HiRefresh } from "react-icons/hi";

// interface ThemeToggleProp {
//   // refreshData: () => void;
//   callRefreshData: (refreshData: () => void) => void;
// }

export default function NavBar() {
  return (
    <Navbar fluid rounded>
      <Navbar.Brand as={Link} href="/dashboard" className="pl-12">
        <Image
          src="/favicon.ico"
          className="mr-3 h-6 sm:h-9"
          alt="Storage Monitor Logo"
          width={36}
          height={24}
          style={{ width: '36', height: '24' }}
        />
        <span className="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
          Storage Monitor
        </span>
      </Navbar.Brand>
      <Navbar.Toggle />

      <Navbar.Collapse className="pl-10 pr-4">
        <Navbar.Link href="#">
          <div className="inline-flex rounded-md shadow-sm" role="group">
            <DarkThemeToggle />
          </div>
        </Navbar.Link>
      </Navbar.Collapse>
    </Navbar>
  );
}
