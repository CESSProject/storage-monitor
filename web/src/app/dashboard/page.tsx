"use client";
import React, {Fragment, useCallback, useEffect, useState} from "react";
import Miners, {HostModel} from "./miners";
import {getApiServerUrl} from "@/utils";
import axios from "axios";

export default function Page() {
    const [data, setData] = useState<HostModel[]>([]);
    const [filteredData, setFilteredData] = useState<HostModel[]>([]);
    const [search, setSearch] = useState<string>("");
    const params = new URLSearchParams({host: search});
    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearch(e.target.value);
    };
    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === "Enter") {
            handleSearch().then(r => console.log(r));
        }
    };
    const handleSearch = useCallback(async () => {
        try {
            const response = await axios.get(`${getApiServerUrl()}/list?${params.toString()}`, {});
            if (!response.data) {
                throw new Error(
                    "Server responded with an error. Please check the watchdog status or contact support."
                );
            }
            let data: HostModel[] = response.data;
            setData(data);
            setFilteredData(data);
        } catch (error) {
            console.error("Failed to fetch data:", error);
        }
    }, []);
    useEffect(() => {
        handleSearch().then(r => console.log(r));
    }, [handleSearch]);
    useEffect(() => {
        if (search.length > 0) {
            setFilteredData(
                data.filter((d) => {
                    return d.Host.includes(search);
                })
            );
        } else {
            setFilteredData(data);
        }
    }, [search]);
    return (
        <Fragment>
            <section className="pl-12 pr-4 bg-white dark:bg-gray-900">
                <div className="py-8 px-4 mx-auto max-w-full lg:pt-16 flex items-center justify-center">
                    <div className="relative ">
                        <input type="text" id="search"
                               className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                               placeholder="Search by host IP"
                               required
                               value={search}
                               onChange={handleInputChange}
                               onKeyDown={handleKeyDown}/>
                    </div>
                    <button type="submit"
                            onClick={handleSearch}
                            className="p-2.5 ms-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                        <svg className="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                             viewBox="0 0 20 20">
                            <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                                  d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"/>
                        </svg>
                        <span className="sr-only">Search</span>
                    </button>
                </div>
            </section>
            {filteredData.map((host) => {
                return <Miners key={host.Host} host={host}/>;
            })}
        </Fragment>
    );
}
