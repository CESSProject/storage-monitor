"use client";
import React, {Fragment, useCallback, useEffect, useState} from "react";
import Miners, {HostModel} from "./miners";
import axios from "axios";
import {env} from "next-runtime-env";
import {notification, Select} from 'antd';
import {DesktopOutlined} from "@ant-design/icons";

const {Option} = Select;
const API_URL = env("NEXT_PUBLIC_API_URL") || "http://localhost:13081";

export default function Page() {
    const [hosts, setHosts] = useState<string[]>([]);
    const [hostData, setHostData] = useState<HostModel | null>(null);
    const [selectedHost, setSelectedHost] = useState<string | null>("");

    const fetchHosts = useCallback(async () => {
        try {
            const response = await axios.get(`${API_URL}/hosts`);
            if (Array.isArray(response.data)) {
                setHosts(response.data);
                if (response.data.length > 0) {
                    setSelectedHost(response.data[0]);
                    return response.data[0];
                } else {
                    setSelectedHost("");
                }
            } else {
                throw new Error("Invalid host list data");
            }
        } catch (error) {
            console.error("Failed to fetch host list", error);
            notification.error({
                message: 'Failed to fetch host list',
            });
        }
    }, []);
    const fetchHostData = useCallback(async (host: string) => {
        if (!host) {
            notification.error({
                message: 'Please select a host',
            });
            return;
        }
        try {
            const response = await axios.get(`${API_URL}/list?host=${encodeURIComponent(host)}`);
            if (!response.data || !Array.isArray(response.data)) {
                throw new Error(
                    "Server responded with invalid data. Please check the API or contact support."
                );
            }
            const data: HostModel = response.data[0];
            setHostData(data);
        } catch (error) {
            console.error("Failed to fetch data", error);
            notification.error({
                message: 'Failed to fetch host data',
            });
        }
    }, []);

    useEffect(() => {
        fetchHosts().then(firstHost => {
            if (firstHost) {
                fetchHostData(firstHost).then(r => {
                });
            }
        }).catch(console.error);
    }, [fetchHosts, fetchHostData]);

    const handleHostChange = (host: string) => {
        setSelectedHost(host);
        fetchHostData(host).then(r => {
        });
    };

    return (
        <Fragment>
            <div style={{display: 'flex', justifyContent: 'center', marginTop: '20px'}}>
                {hosts.length > 0 && (
                    <>
                        <DesktopOutlined className="mr-2 text-2xl text-black dark:text-white"/>
                        <Select
                            placeholder="select a host"
                            onChange={handleHostChange}
                            value={selectedHost || undefined}
                            style={{width: 250, height: 40}}
                            className="bg-white dark:bg-gray-800"
                        >
                            {hosts.map((host) => (
                                <Option key={host} value={host}>
                                    {host}
                                </Option>
                            ))}
                        </Select>
                    </>
                )}
            </div>
            {hostData && (
                <div className="mt-4 space-y-12 dark:bg-gray-800 text-gray-900 dark:text-gray-100">
                    <Miners key={hostData?.Host ? hostData.Host : ""} host={hostData}/>
                </div>
            )}
        </Fragment>
    );
}
