import Link from "next/link";
import { MinerInfoListModel } from "../miner";
import { Button, Spinner } from "flowbite-react";

interface MinerProp {
  host: string | null;
  miners: MinerInfoListModel[] | undefined;
}

export default function Page({ host, miners }: MinerProp) {
  return (
    <section className="pl-12 pr-4 bg-white dark:bg-gray-900 h-full">
      <div className="py-8 px-4 mx-auto max-w-full lg:py-16">
        <h1 className="mb-4 text-xl font-extrabold tracking-tight leading-none text-gray-900 md:text-5xl lg:text-2xl dark:text-white">
          Server: <Link href={`/dashboard/host?host=${host}`}>{host}</Link>
        </h1>
        <div className="flex flex-col w-full space-y-4 sm:flex-row sm:justify-center sm:space-y-0">
          {miners != null ? (
            miners?.map((m) => {
              return <div key={m.AccountId}>Miner {m.AccountId}</div>;
            })
          ) : (
            <div className="flex flex-row gap-3">
              <Button color="gray">
                <Spinner aria-label="Alternate spinner button example" size="sm" />
                <span className="pl-3">Loading...</span>
              </Button>
            </div>
          )}
        </div>
      </div>
    </section>
  );
}
