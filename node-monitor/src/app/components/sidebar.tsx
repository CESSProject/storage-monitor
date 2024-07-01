"use client";

import { Drawer, Sidebar, TextInput } from "flowbite-react";
import { useState } from "react";
import { FaBarsStaggered } from "react-icons/fa6";
import {
  HiChartPie,
  HiClipboard,
  HiCollection,
  HiSearch,
} from "react-icons/hi";
import { HiComputerDesktop, HiInformationCircle, HiSquaresPlus } from "react-icons/hi2";

export function LeftDrawer() {
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = () => setIsOpen(false);

  return (
    <>
      <Drawer
        edge
        open={isOpen}
        onClose={handleClose}
        position="left"
        className="p-0"
        theme={{ root: { edge: "left-12" } }}
      >
        <Drawer.Header
          closeIcon={FaBarsStaggered}
          title="Storage Monitor"
          titleIcon={HiSquaresPlus}
          onClick={() => setIsOpen(!isOpen)}
          className="cursor-pointer px-4 pt-4 hover:bg-gray-50 dark:hover:bg-gray-700"
        />
        <Drawer.Items>
          <Sidebar
            aria-label="Sidebar"
            className="[&>div]:bg-transparent [&>div]:p-0"
          >
            <div className="flex h-full flex-col justify-between py-2">
              <div>
                <form className="pb-3 md:hidden">
                  <TextInput
                    icon={HiSearch}
                    type="search"
                    placeholder="Search"
                    required
                    size={32}
                  />
                </form>
                <Sidebar.Items>
                  <Sidebar.ItemGroup>
                    <Sidebar.Item href="/dashboard" icon={HiChartPie}>
                      Dashboard
                    </Sidebar.Item>
                    <Sidebar.Item href="/dashboard/system" icon={HiComputerDesktop}>
                      System
                    </Sidebar.Item>
                  </Sidebar.ItemGroup>
                  <Sidebar.ItemGroup>
                    <Sidebar.Item
                      href="https://github.com/themesberg/flowbite-react/"
                      icon={HiClipboard}
                    >
                      Docs
                    </Sidebar.Item>
                    <Sidebar.Item
                      href="https://flowbite-react.com/"
                      icon={HiCollection}
                    >
                      Components
                    </Sidebar.Item>
                    <Sidebar.Item
                      href="https://github.com/themesberg/flowbite-react/issues"
                      icon={HiInformationCircle}
                    >
                      Help
                    </Sidebar.Item>
                  </Sidebar.ItemGroup>
                </Sidebar.Items>
              </div>
            </div>
          </Sidebar>
        </Drawer.Items>
      </Drawer>
    </>
  );
}
