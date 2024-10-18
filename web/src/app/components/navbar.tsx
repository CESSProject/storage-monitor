"use client";
import {Layout, Menu, Switch} from "antd";
import {MenuOutlined} from "@ant-design/icons";
import Image from "next/image";
import Link from "next/link";
import {useEffect, useState} from "react";

const {Header} = Layout;

export default function NavBar() {
    const [isDarkMode, setIsDarkMode] = useState(false);

    useEffect(() => {
        const isDark = localStorage.getItem('darkMode') === 'true';
        setIsDarkMode(isDark);
        if (isDark) {
            document.documentElement.classList.add('dark');
        } else {
            document.documentElement.classList.remove('dark');
        }
    }, []);

    const toggleDarkMode = () => {
        const newDarkMode = !isDarkMode;
        setIsDarkMode(newDarkMode);
        localStorage.setItem('darkMode', newDarkMode.toString());
        if (newDarkMode) {
            document.documentElement.classList.add('dark');
        } else {
            document.documentElement.classList.remove('dark');
        }
    };

    return (
        <Header className="flex justify-between items-center bg-white dark:bg-gray-800">
            <div className="flex items-center">
                <Link href="/dashboard" className="flex items-center">
                    <Image
                        src="/favicon.ico"
                        className="mr-3 h-6 sm:h-9 auto"
                        alt="Storage Monitor Logo"
                        width={36}
                        height={24}
                    />
                    <span className="text-xl font-semibold dark:text-white">
                        Storage Miner Monitor
                    </span>
                </Link>
            </div>

            <div className="flex items-center">
                <Menu mode="horizontal" className="border-0 ">
                    <Menu.Item key="docs" className="dark:bg-gray-800">
                        <a
                            href="https://doc.cess.network"
                            target="_blank"
                            rel="noopener noreferrer"
                            className="flex items-center text-black bg-white dark:bg-gray-800"
                        >
                            <svg
                                className="w-6 h-6 text-gray-800 dark:text-white"
                                aria-hidden="true"
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 16 20"
                            >
                                <path
                                    stroke="currentColor"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth="2"
                                    d="M1 17V2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a2 2 0 0 0-2 2Zm0 0a2 2 0 0 0 2 2h12M5 15V1m8 18v-4"
                                />
                            </svg>
                        </a>
                    </Menu.Item>
                </Menu>
                <Switch
                    checked={isDarkMode}
                    onChange={toggleDarkMode}
                    checkedChildren="🌙"
                    unCheckedChildren="☀️"
                />
            </div>

            <div className="md:hidden">
                <MenuOutlined className="text-2xl"/>
            </div>
        </Header>
    );
}