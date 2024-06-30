import { Table } from "flowbite-react";
import Link from "next/link";
import React, { Fragment } from "react";


import {
  CheckCircleOutlined,
  CloseCircleOutlined,
} from '@ant-design/icons';
import { Tag } from 'antd';

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

export default function Miner({ host, miners }: MinerProp) {
  return miners?.map((miner) => {
    return (
      <Table.Row key={miner.SignatureAcc} className="bg-white dark:border-gray-700 dark:bg-gray-800">
        <Table.Cell className="text-center">{miner.Name}</Table.Cell>
        <Table.Cell className="w-24">{miner.SignatureAcc}</Table.Cell>
        <Table.Cell className="text-center">
          {miner.MinerStat.status == "positive" ? (
              <Tag icon={<CheckCircleOutlined />} color="success">Running</Tag>
          ) : (
            <Tag icon={<CloseCircleOutlined />} color="error">Stop</Tag>
          )}
        </Table.Cell>
        <Table.Cell className="text-center">{miner.MinerStat.declaration_space}</Table.Cell>
        <Table.Cell className="text-center">{miner.Conf.UseSpace} GiB</Table.Cell>
        <Table.Cell className="text-center">{miner.MinerStat.total_reward} GiB</Table.Cell>
        <Table.Cell className="text-center">{miner.MinerStat.reward_issued} GiB</Table.Cell>
        <Table.Cell className="text-center">{miner.MinerStat.idle_space}</Table.Cell>
        <Table.Cell className="text-center">{miner.MinerStat.service_space}</Table.Cell>
        <Table.Cell className="text-center">{miner.CInfo.mem_usage} MiB</Table.Cell>
        <Table.Cell className="text-center">{miner.CInfo.cpu_percent * 100}% </Table.Cell>
      </Table.Row>
    );
  });
}