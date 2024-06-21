import { Fragment } from "react";
import Miner, { MinerInfoListModel } from "./miner";
import { Accordion, Table } from "flowbite-react";
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
        <div className="py-8 px-4 mx-auto max-w-full lg:py-16">
          <h1 className="mb-4 text-xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-2xl dark:text-white">
            Server: <Link href={`dashboard/host?host=${host.Host}`}>{host.Host}</Link>
          </h1>
          <div className="flex flex-col w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0">
            <div className="overflow-x-auto w-full">
              <Table>
                <Table.Head>
                  <Table.HeadCell className="w-16">Host</Table.HeadCell>
                  <Table.HeadCell className="w-24">Miner Id</Table.HeadCell>
                  <Table.HeadCell>Miner Name</Table.HeadCell>
                  <Table.HeadCell>Miner Status</Table.HeadCell>
                  <Table.HeadCell>
                    <span className="sr-only">Edit</span>
                  </Table.HeadCell>
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
