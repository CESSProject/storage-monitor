import {Badge, Table} from "flowbite-react";
import React, {useState} from "react";
import {HiCheck, HiX} from 'react-icons/hi';
import {unixTimestampToDateFormat} from '../util/util';
import {Modal} from 'antd';
import MinerDescription from "@/app/components/description";

// Conf container config
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
    cpu_percent: number;
    memory_percent: number;
    mem_usage: number;
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

interface MinerProp {
    host: string;
    miners: MinerInfoListModel[] | null;
}

export default function Miner({host, miners}: MinerProp) {
    const [isModalVisible, setIsModalVisible] = useState(false);

    const showModal = () => {
        setIsModalVisible(true);
    };

    const handleCancel = () => {
        setIsModalVisible(false);
    };

    return miners
        ?.sort((a, b) => a.Name.localeCompare(b.Name))
        .map((miner) => {
            return (
                <Table.Row key={miner.Name} className="bg-white dark:border-gray-700 dark:bg-gray-800">
                    <Table.Cell
                        className="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                        onClick={showModal}>
                        {miner.Name}
                    </Table.Cell>
                    <Modal
                        open={isModalVisible}
                        onCancel={handleCancel}
                        keyboard={true}
                        onOk={handleCancel}
                        maskClosable={true}
                        width="50%"
                        style={{ maxWidth: '1200px' }}
                    >
                        <MinerDescription miner={miner}></MinerDescription>
                    </Modal>
                    <Table.Cell className="w-24">{miner.SignatureAcc}</Table.Cell>
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
                    <Table.Cell className="text-center">{miner.MinerStat.total_reward} GiB</Table.Cell>
                    <Table.Cell className="text-center">{miner.MinerStat.reward_issued} GiB</Table.Cell>
                    <Table.Cell className="text-center">{miner.MinerStat.idle_space}</Table.Cell>
                    <Table.Cell className="text-center">{miner.MinerStat.service_space}</Table.Cell>
                    <Table.Cell className="text-center">{unixTimestampToDateFormat(miner.CInfo.created)}</Table.Cell>
                </Table.Row>
            );
        });
}