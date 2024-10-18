"use client";

import { Drawer, Menu } from "antd";
import { useState } from "react";
import { MenuOutlined, PieChartOutlined, DesktopOutlined, AppstoreAddOutlined } from "@ant-design/icons";
import Link from "next/link";

export function LeftDrawer() {
    const [isOpen, setIsOpen] = useState(false);

    const handleClose = () => setIsOpen(false);

    return (
        <>
            <Drawer
                placement="left"
                onClose={handleClose}
                open={isOpen}
                width={250}
                closeIcon={<MenuOutlined />}
                className="bg-white dark:bg-gray-800 transition-colors duration-200"
                title={
                    <div
                        onClick={() => setIsOpen(!isOpen)}
                        className="cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-700"
                    >
                        <AppstoreAddOutlined /> Storage Monitor
                    </div>
                }
            >
                <Menu
                    mode="inline"
                    theme="light"
                    className="h-full border-r-0 bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300"
                    items={[
                        {
                            key: "dashboard",
                            icon: <PieChartOutlined />,
                            label: <Link href="/dashboard">Dashboard</Link>,
                        },
                        {
                            key: "system",
                            icon: <DesktopOutlined />,
                            label: <Link href="/system">Configuration</Link>,
                        },
                    ]}
                />
            </Drawer>
            <MenuOutlined
                onClick={() => setIsOpen(true)}
                className="fixed top-4 left-4 cursor-pointer text-gray-700 dark:text-gray-300"
            />
        </>
    );
}