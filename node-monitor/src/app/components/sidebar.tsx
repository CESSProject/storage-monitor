"use client";

import {Drawer, Sidebar} from "flowbite-react";
import {useState} from "react";
import {FaBarsStaggered} from "react-icons/fa6";
import {HiChartPie,} from "react-icons/hi";
import {HiComputerDesktop, HiSquaresPlus} from "react-icons/hi2";

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
                className="p-2 w-50"
                theme={{root: {edge: "left-12"}}}
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
                                <Sidebar.Items>
                                    <Sidebar.ItemGroup>
                                        <Sidebar.Item href="/dashboard" icon={HiChartPie}>
                                            Dashboard
                                        </Sidebar.Item>
                                        <Sidebar.Item href="/system" icon={HiComputerDesktop}>
                                            Configuration
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
