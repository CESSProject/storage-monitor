"use client";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  // Navigate to the dashboard page
  router.push("/dashboard");

  return null;
}
