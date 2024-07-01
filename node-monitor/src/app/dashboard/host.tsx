import React, {Fragment} from "react";
import Miner, {MinerInfoListModel} from "./miner";
import {Table} from "flowbite-react";

export interface HostModel {
    Host: string;
    MinerInfoList: MinerInfoListModel[];
}

interface HostProp {
    host: HostModel;
}

export default function Host({host}: HostProp) {
    return (
        <Fragment>
            <section className="pl-12 pr-4 bg-white dark:bg-gray-900">
                <div className="py-8 px-4 mx-auto max-w-full">
                    <h1 className="mb-4 text-xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-2xl dark:text-white">
                        <mark
                            className="px-2 text-white bg-blue-600 rounded dark:bg-blue-500">Host
                        </mark>
                        &nbsp;&nbsp; {host.Host}
                    </h1>
                    <div className=" w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0">
                        <div className="overflow-x-auto overflow-y-auto w-full">
                            <Table>
                                <Table.Head>
                                    <Table.HeadCell className="w-[200px] text-center">Name</Table.HeadCell>
                                    <Table.HeadCell className="w-[200px]">Signature Account</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center">Status</Table.HeadCell>
                                    <Table.HeadCell className="w-[200px] text-center">Declaration Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[180px] text-center">Available Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center">Idle Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center">Used Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center">Total Reward</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center">Used Reward</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center">Create Time</Table.HeadCell>
                                </Table.Head>
                                <Table.Body className="divide-y">
                                    <Miner host={host.Host} miners={host.MinerInfoList}/>
                                </Table.Body>
                            </Table>
                        </div>
                    </div>
                </div>
            </section>
        </Fragment>
    );
}
