import { Table } from "flowbite-react";
import Link from "next/link";
import { Fragment } from "react";

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
      <Table.Row
        key={miner.SignatureAcc}
        className="bg-white dark:border-gray-700 dark:bg-gray-800"
      >
        <Table.Cell className="w-16 whitespace-nowrap font-medium text-gray-900 dark:text-white">
          <Link
            className="text-blue-400 hover:text-blue-500"
            href={`dashboard/host?host=${host}`}
          >
            {host}
          </Link>
        </Table.Cell>
        <Table.Cell className="w-24">{miner.MinerStat.peer_id}</Table.Cell>
        <Table.Cell>{miner.Name}</Table.Cell>
        <Table.Cell className="justify-center">
          {miner.MinerStat.status == "positive" ? (
            <span className="flex w-3 h-3 me-3 bg-green-500 rounded-full"></span>
          ) : (
            <span className="flex w-3 h-3 me-3 bg-red-500 rounded-full"></span>
          )}
        </Table.Cell>
        <Table.Cell>
          <a
            href="#"
            className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
          >
            Edit
          </a>
        </Table.Cell>
      </Table.Row>
    );
  });
}
