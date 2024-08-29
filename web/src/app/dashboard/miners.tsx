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
    is_punished: PunishmentModel;
    total_reward: number;
    reward_issued: number;
}

interface PunishmentModel {
    block_id: number;
    extrinsic_hash: string;
    extrinsic_name: string;
    block_hash: string;
    account: string;
    recv_account: string;
    amount: string;
    type: number;
    timestamp: number;
}

export interface MinerInfoListModel {
    Name: string;
    SignatureAcc: string;
    Conf: ConfModel;
    CInfo: CInfoModel;
    MinerStat: MinerStatModel;
}

function naturalSort(a: string, b: string): number {
    const regex = /(\d+)|(\D+)/g;
    const aParts = a.match(regex) || [];
    const bParts = b.match(regex) || [];
    for (let i = 0; i < Math.min(aParts.length, bParts.length); i++) {
        const aPart = aParts[i];
        const bPart = bParts[i];
        let result = 0;
        if (aPart !== bPart) {
            const aIsNumber = !isNaN(Number(aPart));
            const bIsNumber = !isNaN(Number(bPart));

            if (aIsNumber && bIsNumber) {
                result = Number(aPart) - Number(bPart);
            } else {
                result = aPart.localeCompare(bPart);
            }
        }
        if (result !== 0) {
            return result;
        }
    }
    return aParts.length - bParts.length;
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

    function renderTableRow(miner: MinerInfoListModel) {
        return (
            <Table.Row key={miner.SignatureAcc} className="bg-white dark:border-gray-700 dark:bg-gray-800">
                <Table.Cell
                    className="text-center font-medium text-blue-600 dark:text-blue-500"
                    onClick={() => showModal(miner)}
                >
                    {miner.Name}
                </Table.Cell>
                <Table.Cell className="w-24 text-blue-600 dark:text-blue-500" onClick={() => showModal(miner)}>
                    {miner.SignatureAcc}
                </Table.Cell>
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
                <Table.Cell className="text-center">{miner.MinerStat.declaration_space}</Table.Cell>
                <Table.Cell className="text-center">{miner.Conf.UseSpace} GiB</Table.Cell>
                <Table.Cell className="text-center">{miner.MinerStat.idle_space}</Table.Cell>
                <Table.Cell className="text-center">{miner.MinerStat.service_space}</Table.Cell>
                <Table.Cell className="text-center">{miner.MinerStat.total_reward}</Table.Cell>
                <Table.Cell className="text-center">{miner.MinerStat.reward_issued}</Table.Cell>
                <Table.Cell className="text-center">{unixTimestampToDateFormat(miner.CInfo.created)}</Table.Cell>
            </Table.Row>
        );
    }

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
                                    <Table.HeadCell
                                        className="w-[200px] text-center preserve-case">Name</Table.HeadCell>
                                    <Table.HeadCell className="w-[200px] preserve-case">Signature
                                        Account</Table.HeadCell>
                                    <Table.HeadCell
                                        className="w-[150px] text-center preserve-case">Status</Table.HeadCell>
                                    <Table.HeadCell className="w-[200px] text-center preserve-case">Declaration
                                        Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[180px] text-center preserve-case">Available
                                        Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center preserve-case">Idle
                                        Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center preserve-case">Used
                                        Space</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center preserve-case">Total
                                        Reward</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center preserve-case">Claimed
                                        Reward</Table.HeadCell>
                                    <Table.HeadCell className="w-[150px] text-center preserve-case">Create
                                        Time</Table.HeadCell>
                                </Table.Head>

                                <Table.Body className="divide-y">
                                    {host?.MinerInfoList
                                        ? host.MinerInfoList
                                            .sort((a, b) => naturalSort(a.Name, b.Name))
                                            .map(renderTableRow)
                                        : null}
                                </Table.Body>
                            </Table>
                            <Modal
                                mask={true}
                                confirmLoading={true}
                                footer={null}
                                open={isModalVisible}
                                onCancel={handleCancel}
                                keyboard={true}
                                onOk={handleCancel}
                                maskClosable={true}
                                styles={{
                                    body: {backgroundColor: 'darkgray'},
                                    header: {backgroundColor: 'darkgray'},
                                    content: {backgroundColor: 'darkgray'},
                                }}
                                width="60%"
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
