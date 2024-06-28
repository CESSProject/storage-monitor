"use client";

import Link from "next/link";
import { MinerInfoListModel } from "../miner";
import { Button, Spinner, Table } from "flowbite-react";
import { Fragment, useEffect, useState } from "react";

interface MinerProp {
  host: string | null;
  miners: MinerInfoListModel[] | undefined;
}

export default function Page({ host, miners }: MinerProp) {
  const [filteredMiners, setFilteredMiners] = useState<
    MinerInfoListModel[] | undefined
  >(undefined);
  const [search, setSearch] = useState<string>("");

  useEffect(() => {
    if (miners) {
      setFilteredMiners(miners);
    }
  }, [miners]);

  useEffect(() => {
    if (miners) {
      setFilteredMiners(
        miners?.filter((m) => {
          return m.Name.toLowerCase().includes(search.toLowerCase());
        })
      );
    }
  }, [search, miners]);

  return (
    <Fragment>
      <section className="pr-4 bg-white dark:bg-gray-900">
        <div className="py-8 px-4 mx-auto max-w-full lg:pt-16">
          <input
            className="border text-sm rounded-lg block max-w-screen-xl w-64 p-2.5"
            type="text"
            placeholder="Search by Miner Name"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
        </div>
      </section>
      <section className="pr-4 bg-white dark:bg-gray-900 h-full">
        <div className="py-8 px-4 mx-auto max-w-full">
          <h1 className="mb-4 text-xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-2xl dark:text-white">
            Server:{" "}
            <Link
              className="text-blue-400 hover:text-blue-500"
              href={`/dashboard/host?host=${host}`}
            >
              {host}
            </Link>
          </h1>
          <div className="flex flex-col w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0">
            {filteredMiners != null ? (
              <div className="overflow-x-auto w-full">
                <Table>
                  <Table.Head>
                    <Table.HeadCell className="whitespace-nowrap">Miner Name</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Signature Account</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Earnings Account</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Staking Account</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">RPC</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Space Available</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Workspace</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">CPU Cores</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Tee List</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Bootnode</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Container Id</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Miner Nodes</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Network</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Docker Image Id</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Command</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Created</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">State</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Container Status</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">CPU Utilization</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Memory Utilization</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Memory Usage</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Peer Id</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Staking</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Debt</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Miner Status</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Declaration Space</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Idle Space</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Service Space</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Lock Space</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Is Punished</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Total Reward</Table.HeadCell>
                    <Table.HeadCell className="whitespace-nowrap">Reward Issued</Table.HeadCell>
                  </Table.Head>
                  <Table.Body className="divide-y">
                    {filteredMiners?.map((m) => {
                      return (
                        <Fragment key={m.Name}>
                          <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                            <Table.Cell className="text-nowrap">{m.Name}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.SignatureAcc}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.EarningsAcc}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.StakingAcc}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.Rpc}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.UseSpace}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.Workspace}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.UseCpu}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.TeeList.join(", ")}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.Conf.Boot}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.id}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.names.join(", ")}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.image.split(":")[1]}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.image_id}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.command}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.created}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.state}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.status}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.cpu_percent}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.memory_percent}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.CInfo.mem_usage}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.peer_id}</Table.Cell>
                            <Table.Cell className="text-nowrap">
                              {m.MinerStat.collaterals.toString()}
                            </Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.debt}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.status}</Table.Cell>
                            <Table.Cell className="text-nowrap">
                              {m.MinerStat.declaration_space}
                            </Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.idle_space}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.service_space}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.lock_space}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.is_punished.join(", ")}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.total_reward}</Table.Cell>
                            <Table.Cell className="text-nowrap">{m.MinerStat.reward_issued}</Table.Cell>
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
    </Fragment>
  );
}
