"use client";
import {DarkThemeToggle, Navbar} from "flowbite-react";
import Image from "next/image";
import Link from "next/link";

export default function NavBar() {
    return (
        <Navbar fluid rounded>
            <Navbar.Brand as={Link} href="/dashboard" className="pl-12">
                <Image
                    src="/favicon.ico"
                    className="mr-3 h-6 sm:h-9 auto"
                    alt="Storage Monitor Logo"
                    width={36}
                    height={24}
                    style={{width: '36', height: '24'}}
                />
                <span
                    className="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Storage Miner Monitor</span>
            </Navbar.Brand>
            <Navbar.Toggle/>

            <Navbar.Collapse className="pl-10 pr-4">
                <div className="inline-flex rounded-md shadow-sm" role="group">
                    <ul className="flex flex-row justify-center items-center font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                        <li>
                            <a href="https://docs.cess.cloud/core"
                               target="_blank"
                               className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">
                                <svg className="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true"
                                     xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 20">
                                    <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round"
                                          strokeWidth="2"
                                          d="M1 17V2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a2 2 0 0 0-2 2Zm0 0a2 2 0 0 0 2 2h12M5 15V1m8 18v-4"/>
                                </svg>
                            </a>
                        </li>
                        <li>
                            <a className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
                               href="https://github.com/CESSProject/storage-monitor" target="_blank">
                                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                                    <mask id="lineMdGithubLoop0" width="24" height="24" x="0" y="0">
                                        <g fill="#fff">
                                            <ellipse cx="9.5" cy="9" rx="1.5" ry="1"/>
                                            <ellipse cx="14.5" cy="9" rx="1.5" ry="1"/>
                                        </g>
                                    </mask>
                                    <g fill="none" stroke="currentColor" strokeLinecap="round"
                                       strokeLinejoin="round" strokeWidth="2">
                                        <path strokeDasharray="30" strokeDashoffset="30"
                                              d="M12 4C13.6683 4 14.6122 4.39991 15 4.5C15.5255 4.07463 16.9375 3 18.5 3C18.8438 4 18.7863 5.21921 18.5 6C19.25 7 19.5 8 19.5 9.5C19.5 11.6875 19.017 13.0822 18 14C16.983 14.9178 15.8887 15.3749 14.5 15.5C15.1506 16.038 15 17.3743 15 18C15 18.7256 15 21 15 21M12 4C10.3317 4 9.38784 4.39991 9 4.5C8.47455 4.07463 7.0625 3 5.5 3C5.15625 4 5.21371 5.21921 5.5 6C4.75 7 4.5 8 4.5 9.5C4.5 11.6875 4.98301 13.0822 6 14C7.01699 14.9178 8.1113 15.3749 9.5 15.5C8.84944 16.038 9 17.3743 9 18C9 18.7256 9 21 9 21">
                                            <animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s"
                                                     values="30;0"/>
                                        </path>
                                        <path strokeDasharray="10" strokeDashoffset="10" d="M9 19">
                                            <animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s"
                                                     dur="0.2s" values="10;0"/>
                                            <animate
                                                attributeName="d"
                                                dur="3s"
                                                repeatCount="indefinite"
                                                values="M9 19c-1.406 0-2.844-.563-3.688-1.188C4.47 17.188 4.22 16.157 3 15.5;M9 19c-1.406 0-3-.5-4-.5-.532 0-1 0-2-.5;M9 19c-1.406 0-2.844-.563-3.688-1.188C4.47 17.188 4.22 16.157 3 15.5"/>
                                        </path>
                                    </g>
                                    <rect width="8" height="4" x="8" y="11" fill="currentColor"
                                          mask="url(#lineMdGithubLoop0)">
                                        <animate attributeName="y" dur="10s" keyTimes="0;0.45;0.46;0.54;0.55;1"
                                                 repeatCount="indefinite" values="11;11;7;7;11;11"/>
                                    </rect>
                                </svg>
                            </a>
                        </li>
                        <li className="h-10 w-10 flex items-center justify-center py-2 px-3">
                            <DarkThemeToggle/>
                        </li>
                    </ul>
                </div>
            </Navbar.Collapse>
        </Navbar>
    )
        ;
}
