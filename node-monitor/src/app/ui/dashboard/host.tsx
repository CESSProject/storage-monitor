import { Fragment } from "react";
import Miner, { MinerInfoListModel } from "./miner";
import { Table } from "flowbite-react";
import Link from "next/link";

export interface HostModel {
  Host: string;
  MinerInfoList: MinerInfoListModel[];
}

interface HostProp {
  host: HostModel;
}

export default function Host({ host }: HostProp) {
  return (
    <Fragment>
      <section className="pl-12 pr-4 bg-white dark:bg-gray-900">
        <div className="py-8 px-4 mx-auto max-w-full">
          <h1 className="mb-4 text-xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-2xl dark:text-white">
            Host: <Link className="text-blue-400 hover:text-blue-500" href={`dashboard/host?host=${host.Host}`}>{host.Host}</Link>
          </h1>
          <div className="flex flex-col w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0">
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
                  <Table.HeadCell className="w-[160px] text-center">Used Memory</Table.HeadCell>
                  <Table.HeadCell className="w-[150px] text-center">Used CPU</Table.HeadCell>
                </Table.Head>
                <Table.Body className="divide-y">
                  <Miner host={host.Host} miners={host.MinerInfoList} />
                </Table.Body>
              </Table>
            </div>
          </div>
        </div>
      </section>
    </Fragment>
  );
}
