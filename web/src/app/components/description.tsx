import React from 'react';
import {Descriptions, DescriptionsProps, Divider} from 'antd';
import {MinerInfoListModel} from "@/app/dashboard/miners";

const minerInfoToDescriptionItems = (minerInfo: MinerInfoListModel): DescriptionsProps['items'] => [
    {
        label: 'Miner Name',
        children: minerInfo.Name,
        contentStyle: {fontSize: "16px", color: 'black', fontWeight: "bold"},
        style: {borderBottom: '2px solid black',}
    },
    {
        label: 'Signature Account',
        children: minerInfo.SignatureAcc,
        contentStyle: {fontSize: "16px", color: 'black', fontWeight: "bold"},
        style: {borderBottom: '2px solid black'}
    },
    {
        label: 'Configuration',
        span: {xs: 1, sm: 1, md: 1, lg: 1, xl: 1, xxl: 2},
        contentStyle: {fontSize: "16px", color: 'black', fontWeight: "bold"},
        style: {borderBottom: '2px solid black', borderRight: '2px solid black', borderLeft: '2px solid black'},
        children: (
            <>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Name: </strong>{minerInfo.Conf.Name}</p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Port: </strong>{minerInfo.Conf.Port}</p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">EarningsAcc: </strong>{minerInfo.Conf.EarningsAcc}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">StakingAcc: </strong>{minerInfo.Conf.StakingAcc}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Rpc: </strong>{minerInfo.Conf.Rpc.join(', ')}</p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Use Space: </strong>{minerInfo.Conf.UseSpace}GiB
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Use Core: </strong>{minerInfo.Conf.UseCpu}
                </p>
            </>
        ),
    },
    {
        label: 'Container Info',
        span: {xs: 1, sm: 1, md: 1, lg: 1, xl: 1, xxl: 2},
        contentStyle: {fontSize: "16px", color: 'black', fontWeight: "bold"},
        style: {borderBottom: '2px solid black', borderRight: '2px solid black', borderLeft: '2px solid black'},
        children: (
            <>
                <p><strong className="text-sm font-bold text-gray-900 dark:text-black">Container
                    ID: </strong>{minerInfo.CInfo.id}</p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Name: </strong>{minerInfo.CInfo.name}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Status: </strong>{minerInfo.CInfo.state}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Up
                    Time: </strong>{minerInfo.CInfo.status}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">CPU Percent: </strong>{minerInfo.CInfo.cpu_percent}%
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Memory Usage
                    Usage: </strong>{minerInfo.CInfo.mem_usage}MiB
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Memory
                    Usage: </strong>{minerInfo.CInfo.memory_percent}%
                </p>
            </>
        ),
    },
    {
        label: 'Miner Statistics',
        span: {xs: 1, sm: 1, md: 1, lg: 1, xl: 1, xxl: 2},
        contentStyle: {fontSize: "16px", color: 'black', fontWeight: "bold",},
        style: {borderBottom: '2px solid black', borderRight: '2px solid black', borderLeft: '2px solid black'},
        children: (
            <>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">P2P Peer
                    ID: </strong>{minerInfo.MinerStat.peer_id}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Status: </strong>{minerInfo.MinerStat.status}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Declaration
                    Space: </strong>{minerInfo.MinerStat.declaration_space}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Idle
                    Space: </strong>{minerInfo.MinerStat.idle_space}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Service
                    Space: </strong>{minerInfo.MinerStat.service_space}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Total
                    Reward: </strong>{minerInfo.MinerStat.total_reward}
                </p>
                <Divider></Divider>
                <p><strong
                    className="text-sm font-bold text-gray-900 dark:text-black">Claimed Reward: </strong>{minerInfo.MinerStat.reward_issued}
                </p>
            </>
        ),
    },
];

// xs: 1: On very small screens (such as mobile phones), only 1 column is displayed.
// sm: 2: On small screens (such as tablets), 2 columns are displayed.
// md: 3: On medium screens, 3 columns are displayed.
// lg: 3: On large screens, 3 columns are displayed.
// xl: 4: On extra large screens, 4 columns are displayed.
// xxl: 4: On extra large screens, 4 columns are displayed.

const MinerDescription: React.FC<{ miner: MinerInfoListModel }> = ({miner}) => (
    <Descriptions
        labelStyle={{color: 'black', fontWeight: "bold", fontSize: "18px"}}
        title="Miner Information"
        bordered
        column={{xs: 1, sm: 1, md: 1, lg: 1, xl: 1, xxl: 2}}
        items={minerInfoToDescriptionItems(miner)}
    />
);

export default MinerDescription;