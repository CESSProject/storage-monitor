"use client";

import Link from "next/link";
import { MinerInfoListModel } from "../miner";
import { Button, Spinner, Table } from "flowbite-react";
import { Fragment } from "react";

interface MinerProp {
  host: string | null;
  miners: MinerInfoListModel[] | undefined;
}

export default function Page({ host, miners }: MinerProp) {
  return (
    <section className="pl-12 pr-4 bg-white dark:bg-gray-900 h-full">
      <div className="py-8 px-4 mx-auto max-w-full lg:py-16">
        <h1 className="mb-4 text-xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-2xl dark:text-white">
          Server: <Link className="text-blue-400 hover:text-blue-500" href={`/dashboard/host?host=${host}`}>{host}</Link>
        </h1>
        <div className="flex flex-col w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0">
          {miners != null ? (
            <div className="overflow-x-auto w-full">
              <Table>
                <Table.Head>
                  <Table.HeadCell>Miner Name</Table.HeadCell>
                  <Table.HeadCell>Account Id</Table.HeadCell>
                  <Table.HeadCell>Port</Table.HeadCell>
                  <Table.HeadCell>Earnings Account</Table.HeadCell>
                  <Table.HeadCell>Staking Account</Table.HeadCell>
                  <Table.HeadCell>RPC</Table.HeadCell>
                  <Table.HeadCell>Use Space</Table.HeadCell>
                  <Table.HeadCell>Workspace</Table.HeadCell>
                  <Table.HeadCell>CPU Cores</Table.HeadCell>
                  <Table.HeadCell>Tee List</Table.HeadCell>
                  <Table.HeadCell>Bootnode</Table.HeadCell>
                  <Table.HeadCell>Container Id</Table.HeadCell>
                  <Table.HeadCell>Miner Nodes</Table.HeadCell>
                  <Table.HeadCell>Docker Image</Table.HeadCell>
                  <Table.HeadCell>Docker Image Id</Table.HeadCell>
                  <Table.HeadCell>Command</Table.HeadCell>
                  <Table.HeadCell>Created</Table.HeadCell>
                  <Table.HeadCell>State</Table.HeadCell>
                  <Table.HeadCell>Container Status</Table.HeadCell>
                  <Table.HeadCell>CPU Utilization</Table.HeadCell>
                  <Table.HeadCell>Memory Utilization</Table.HeadCell>
                  <Table.HeadCell>Memory Usage</Table.HeadCell>
                  <Table.HeadCell>Peer Id</Table.HeadCell>
                  <Table.HeadCell>Staking</Table.HeadCell>
                  <Table.HeadCell>Debt</Table.HeadCell>
                  <Table.HeadCell>Miner Status</Table.HeadCell>
                  <Table.HeadCell>Declaration Space</Table.HeadCell>
                  <Table.HeadCell>Idle Space</Table.HeadCell>
                  <Table.HeadCell>Service Space</Table.HeadCell>
                  <Table.HeadCell>Lock Space</Table.HeadCell>
                  <Table.HeadCell>Is Punished</Table.HeadCell>
                  <Table.HeadCell>Total Reward</Table.HeadCell>
                  <Table.HeadCell>Reward Issued</Table.HeadCell>
                  <Table.HeadCell>
                    <span className="sr-only">Edit</span>
                  </Table.HeadCell>
                </Table.Head>
                <Table.Body className="divide-y">
                  {miners?.map((m) => {
                    // return <div key={m.AccountId}>Miner {m.AccountId}</div>;
                    return (
                      <Fragment>
                        <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                          <Table.Cell>{m.Name}</Table.Cell>
                          <Table.Cell>{m.AccountId}</Table.Cell>
                          <Table.Cell>{m.Conf.Port}</Table.Cell>
                          <Table.Cell>{m.Conf.EarningsAcc}</Table.Cell>
                          <Table.Cell>{m.Conf.StakingAcc}</Table.Cell>
                          <Table.Cell>{m.Conf.Rpc}</Table.Cell>
                          <Table.Cell>{m.Conf.UseSpace}</Table.Cell>
                          <Table.Cell>{m.Conf.Workspace}</Table.Cell>
                          <Table.Cell>{m.Conf.UseCpu}</Table.Cell>
                          <Table.Cell>{m.Conf.TeeList.join(', ')}</Table.Cell>
                          <Table.Cell>{m.Conf.Boot}</Table.Cell>
                          <Table.Cell>{m.CInfo.id}</Table.Cell>
                          <Table.Cell>{m.CInfo.names}</Table.Cell>
                          <Table.Cell>{m.CInfo.image}</Table.Cell>
                          <Table.Cell>{m.CInfo.image_id}</Table.Cell>
                          <Table.Cell>{m.CInfo.command}</Table.Cell>
                          <Table.Cell>{m.CInfo.created}</Table.Cell>
                          <Table.Cell>{m.CInfo.state}</Table.Cell>
                          <Table.Cell>{m.CInfo.status}</Table.Cell>
                          <Table.Cell>{m.CInfo.cpu_percent}</Table.Cell>
                          <Table.Cell>{m.CInfo.memory_percent}</Table.Cell>
                          <Table.Cell>{m.CInfo.mem_usage}</Table.Cell>
                          <Table.Cell>{m.MinerStat.peer_id}</Table.Cell>
                          <Table.Cell>
                            {m.MinerStat.collaterals.toString()}
                          </Table.Cell>
                          <Table.Cell>{m.MinerStat.debt}</Table.Cell>
                          <Table.Cell>{m.MinerStat.status}</Table.Cell>
                          <Table.Cell>
                            {m.MinerStat.declaration_space}
                          </Table.Cell>
                          <Table.Cell>{m.MinerStat.idle_space}</Table.Cell>
                          <Table.Cell>{m.MinerStat.service_space}</Table.Cell>
                          <Table.Cell>{m.MinerStat.lock_space}</Table.Cell>
                          <Table.Cell>{m.MinerStat.is_punished}</Table.Cell>
                          <Table.Cell>{m.MinerStat.total_reward}</Table.Cell>
                          <Table.Cell>{m.MinerStat.reward_issued}</Table.Cell>
                          <Table.Cell>
                            <a
                              href="#"
                              className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                            >
                              Edit
                            </a>
                          </Table.Cell>
                        </Table.Row>
                      </Fragment>
                    );
                  })}
                </Table.Body>
              </Table>
            </div>
          ) : (
            <div className="flex flex-row gap-3">
              <Button color="gray">
                <Spinner
                  aria-label="Alternate spinner button example"
                  size="sm"
                />
                <span className="pl-3">Loading...</span>
              </Button>
            </div>
          )}
        </div>
      </div>
    </section>
  );
}
