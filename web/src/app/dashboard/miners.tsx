import React, {Fragment, useState} from "react";
import {Badge, Modal, Table} from "antd";
import MinerDescription from "@/app/components/description";
import {CheckCircleOutlined, CloseCircleOutlined} from "@ant-design/icons";
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

    const columns = [
        {
            title: "Name",
            dataIndex: "Name",
            key: "Name",
            align: "center" as const,
            render: (text: string, record: MinerInfoListModel) => (
                <span className="text-blue-600 dark:text-blue-500 cursor-pointer"
                      onClick={() => showModal(record)}>{text}</span>
            ),
        },
        {
            title: () => <div style={{textAlign: 'left'}}>Signature Account</div>,
            dataIndex: "SignatureAcc",
            key: "SignatureAcc",
            align: "left" as const,
            width: 250,
            render: (text: string, record: MinerInfoListModel) => (
                <span
                    className="text-blue-600 dark:text-blue-500 cursor-pointer"
                    onClick={() => showModal(record)}
                >
                        {text}
                    </span>
            ),
        },
        {
            title: "Status",
            dataIndex: ["MinerStat", "status"],
            key: "status",
            align: "center" as const,
            render: (status: string) => (
                status === "positive" ? (
                    <Badge status="success" text="Running" icon={<CheckCircleOutlined/>}/>
                ) : (
                    <Badge status="error" text="Stop" icon={<CloseCircleOutlined/>}/>
                )
            ),
        },
        {
            title: "Declaration Space",
            dataIndex: ["MinerStat", "declaration_space"],
            key: "declaration_space",
            align: "center" as const,
        },
        {
            title: "Idle Space",
            dataIndex: ["MinerStat", "idle_space"],
            key: "idle_space",
            align: "center" as const,
        },
        {
            title: "Used Space",
            dataIndex: ["MinerStat", "service_space"],
            key: "service_space",
            align: "center" as const,
        },
        {
            title: "Total Reward",
            dataIndex: ["MinerStat", "total_reward"],
            key: "total_reward",
            align: "center" as const,
        },
        {
            title: "Claimed Reward",
            dataIndex: ["MinerStat", "reward_issued"],
            key: "reward_issued",
            align: "center" as const,
        },
        {
            title: "Create Time",
            dataIndex: ["CInfo", "created"],
            key: "created",
            align: "center" as const,
            render: (created: number) => unixTimestampToDateFormat(created),
        },
    ];


    return (
        <Fragment>
            <section className="pl-12 pr-4 pt-0 bg-white dark:bg-gray-900">
                <div className="py-8 px-4 mx-auto max-w-full">
                    <div key={host?.Host} className="mb-8 p-4 rounded-lg shadow-md border-2 border-gray-900 bg-white dark:bg-gray-400">
                        <h1 className="text-xl font-bold mb-4 text-blue-600 dark:text-white">
                            <mark className="px-2 text-white bg-blue-900 rounded dark:bg-black">Host</mark>
                            &nbsp;&nbsp; {host?.Host ? host.Host : "Unknown"}
                        </h1>
                        <div className="w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0 ">
                            <div className="overflow-x-auto overflow-y-auto w-full">
                                <Table
                                    columns={columns}
                                    dataSource={host?.MinerInfoList?.sort((a, b) => naturalSort(a.Name, b.Name))}
                                    rowKey="SignatureAcc"
                                    pagination={false}
                                    className="bg-white dark:bg-gray-100 text-white"
                                    rowClassName="hover:bg-gray-100 dark:hover:bg-gray-700"
                                />
                                <Modal
                                    open={isModalVisible}
                                    onCancel={handleCancel}
                                    footer={null}
                                    width="60%"
                                    style={{maxWidth: '1200px'}}
                                    className="bg-white dark:bg-gray-800"
                                >
                                    <MinerDescription miner={selectedMiner}/>
                                </Modal>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </Fragment>
    );
}