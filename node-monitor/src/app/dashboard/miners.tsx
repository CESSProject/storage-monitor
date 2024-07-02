import React, {Fragment, useState} from "react";
import {Badge, Table} from "flowbite-react";
import MinerDescription from "@/app/components/description";
import {Modal} from "antd";
import {HiCheck, HiX} from "react-icons/hi";
import {unixTimestampToDateFormat} from "@/utils";

export interface HostModel {
    Host: string;
    MinerInfoList: MinerInfoListModel[];
}

interface HostProp {
    host: HostModel;
}

interface ConfModel {
    Name: string;
    Port: number;
    EarningsAcc: string;
    StakingAcc: string;
    Mnemonic: string;
    Rpc: string[];
    UseSpace: number;
    Workspace: string;
    UseCpu: number;
    TeeList: string[];
    Boot: string[];
}

// CInfo container info
interface CInfoModel {
    id: string;
    names: string[];
    name: string;
    image: string;
    image_id: string;
    command: string;
    created: number;
    state: string;
    status: string;
    cpu_percent: string;
    memory_percent: string;
    mem_usage: string;
}

// Miner Stat
interface MinerStatModel {
    peer_id: string;
    collaterals: BigInt;
    debt: number;
    status: string;
    declaration_space: number;
    idle_space: number;
    service_space: number;
    lock_space: number;
    is_punished: boolean[][];
    total_reward: number;
    reward_issued: number;
}

export interface MinerInfoListModel {
    Name: string;
    SignatureAcc: string;
    Conf: ConfModel;
    CInfo: CInfoModel;
    MinerStat: MinerStatModel;
}

export default function Miners({host}: HostProp) {


    const emptyConf = {} as MinerInfoListModel;

    const [isModalVisible, setIsModalVisible] = useState(false);
    const [selectedMiner, setSelectedMiner] = useState<MinerInfoListModel>(emptyConf);
    const showModal = (miner: MinerInfoListModel) => {
        setSelectedMiner(miner);
        setIsModalVisible(true);
    };
    const handleCancel = () => {
        setIsModalVisible(false);
        setSelectedMiner(emptyConf);
    };

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
                    <div className=" w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0 ">
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
                                    {host.MinerInfoList
                                        ?.sort((a, b) => a.Name.localeCompare(b.Name))
                                        .map((miner) => {
                                            return (
                                                <Table.Row key={miner.SignatureAcc}
                                                           className="bg-white dark:border-gray-700 dark:bg-gray-800">
                                                    <Table.Cell
                                                        className="text-center font-medium text-blue-600 dark:text-blue-500"
                                                        onClick={() => showModal(miner)}
                                                    >
                                                        {miner.Name}
                                                    </Table.Cell>
                                                    <Table.Cell className="w-24 text-blue-600 dark:text-blue-500"
                                                                onClick={() => showModal(miner)}>{miner.SignatureAcc}</Table.Cell>
                                                    <Table.Cell className="text-center">
                                                        {miner.MinerStat.status === "positive" ? (
                                                            <Badge color="success" icon={HiCheck}>
                                                                Running
                                                            </Badge>
                                                        ) : (
                                                            <Badge color="failure" icon={HiX}>
                                                                Stop
                                                            </Badge>
                                                        )}
                                                    </Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{miner.MinerStat.declaration_space}</Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{miner.Conf.UseSpace} GiB</Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{miner.MinerStat.idle_space}</Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{miner.MinerStat.service_space}</Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{miner.MinerStat.total_reward} </Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{miner.MinerStat.reward_issued} </Table.Cell>
                                                    <Table.Cell
                                                        className="text-center">{unixTimestampToDateFormat(miner.CInfo.created)}</Table.Cell>
                                                </Table.Row>
                                            );
                                        })}
                                </Table.Body>
                            </Table>
                            <Modal
                                open={isModalVisible}
                                onCancel={handleCancel}
                                keyboard={true}
                                onOk={handleCancel}
                                maskClosable={true}
                                width="50%"
                                style={{maxWidth: '1200px'}}
                            >
                                <MinerDescription miner={selectedMiner}></MinerDescription>
                            </Modal>
                        </div>
                    </div>
                </div>
            </section>
        </Fragment>
    );
}
