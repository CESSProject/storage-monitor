"use client";

import Host from "@/app/ui/dashboard/host/host";
import { Suspense } from "react";

export default function Page() {
  return (
    <div className="pl-12 pr-4 bg-white dark:bg-gray-900">
      <Suspense fallback={<div>Loading...</div>}>
        <Host />
      </Suspense>
    </div>
  );
}
