import { useCallback, useEffect, useState } from "react";
import { HostModel } from "../host";
import { useSearchParams } from "next/navigation";
import { getApiServerUrl } from "@/utils";
import Miner from "./miner";

export default function Page() {
  const [data, setData] = useState<HostModel | null>(null);
  const searchParams = useSearchParams();
  const host = searchParams.get("host");

  const refreshData = useCallback(async () => {
    try {
      const response = await fetch(`${getApiServerUrl()}/list?host=${host}`, {
        method: "GET",
      });

      if (!response.ok) {
        throw new Error(
          "Server responded with an error. Please check the server status or contact support."
        );
      }

      let data: HostModel[] = await response.json();
      setData(data[0]);
    } catch (error) {
      setData(null);
      console.error("Failed to fetch data:", error);
    }
  }, []);

  useEffect(() => {
    refreshData();
  }, [refreshData]);
  return (
    <Miner host={host} miners={data?.MinerInfoList}/>
  );
}
