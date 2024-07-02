"use client";
import {useRouter} from "next/navigation";

export default function Home() {
    const router = useRouter();
    router.push("/dashboard");
    router.push("/system");
    return <div>{}</div>;
}
