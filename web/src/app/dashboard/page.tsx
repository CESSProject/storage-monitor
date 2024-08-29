"use client";
import React, { Fragment, useCallback, useEffect, useState } from "react";
import Miners, { HostModel } from "./miners";
import axios from "axios";
import { env } from "next-runtime-env";
import { toast } from "sonner";

const API_URL = env("NEXT_PUBLIC_API_URL") || "http://localhost:13081";

export default function Page() {
    const [data, setData] = useState<HostModel[]>([]);
    const [search, setSearch] = useState<string>("");
    const [pageIndex, setPageIndex] = useState<number>(1);
    const [pageSize, setPageSize] = useState<number>(5);

    const urlSafeSearch = encodeURIComponent(search);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearch(e.target.value);
    };

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === "Enter") {
            handleSearch().catch(console.error);
        }
    };

    const handleSearch = useCallback(async () => {
        try {
            const response = await axios.get(`${API_URL}/list?host=${urlSafeSearch}&pageindex=${pageIndex}&pagesize=${pageSize}`);
            if (!response.data || !Array.isArray(response.data)) {
                throw new Error(
                    "Server responded with invalid data. Please check the API or contact support."
                );
            }
            const data: HostModel[] = response.data;
            setData(data);
        } catch (error) {
            console.error("Failed to fetch data", error);
            toast.error("Failed to fetch data. Please check your connection or try again later.");
        }
    }, [urlSafeSearch, pageIndex, pageSize]);

    useEffect(() => {
        let isMounted = true;
        const fetchData = async () => {
            try {
                await handleSearch();
                if (isMounted) {
                    // No need to log the result here
                }
            } catch (error) {
                console.error("Error fetching data", error);
            }
        };
        fetchData();
        return () => {
            isMounted = false;
        };
    }, [handleSearch]);

    const filterDataBySearch = (data: HostModel[], search: string): HostModel[] => {
        if (search.length > 0) {
            return data.filter((d) => d.Host.includes(search));
        }
        return data;
    };

    // Apply filtering after pagination
    const filteredData = filterDataBySearch(data, search);

    return (
        <Fragment>
            <section className="pl-12 pr-4 bg-white dark:bg-gray-900">
                <div className="py-8 px-4 mx-auto max-w-full lg:pt-16 flex items-center justify-center">
                    <div className="relative">
                        <input
                            type="text"
                            id="search"
                            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            placeholder="Search by host IP"
                            required
                            value={search}
                            onChange={handleInputChange}
                            onKeyDown={handleKeyDown}
                        />
                    </div>
                    <button
                        type="submit"
                        onClick={handleSearch}
                        className="p-2.5 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                    >
                        <svg
                            className="w-4 h-4"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 20 20"
                        >
                            <path
                                stroke="currentColor"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth="2"
                                d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                            />
                        </svg>
                        <span className="sr-only">Search</span>
                    </button>
                </div>
            </section>
            <div className="mt-4">
                {filteredData.map((host) => (
                    <Miners key={host.Host} host={host} />
                ))}
            </div>
            <div className="flex justify-center mt-4">
                <button
                    onClick={() => setPageIndex(pageIndex - 1)}
                    disabled={pageIndex === 1}
                    className="px-4 py-2 mr-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                    Previous
                </button>
                <div
                    className="bg-gray-200 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded-lg px-4 py-2 text-sm">
                    Current Page: {pageIndex}
                </div>
                <button
                    onClick={() => setPageIndex(pageIndex + 1)}
                    disabled={filteredData.length < pageSize}
                    className="px-4 py-2 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                    Next
                </button>
            </div>
        </Fragment>
    );
}
