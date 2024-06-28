"use client";

import { getApiServerUrl } from "@/utils";
import { Button, Label, Textarea } from "flowbite-react";
import { useCallback, useEffect, useState } from "react";
import { Toaster, toast } from 'sonner';

interface Hosts {
  ip: string;
  port: string;
  ca_path: string;
  cert_path: string;
  key_path: string;
}

interface Email {
  smtp_endpoint: string;
  smtp_port: number;
  smtp_account: string;
  smtp_password: string;
  receiver: string[];
}

interface Alert {
  enable: boolean;
  webhook: string[];
  email: Email;
}

interface Http {
  http_port: number
}

interface Https {
  https_port: number;
  cert_path: string;
  key_path: string;
}

interface Server {
  http: Http;
  https: Https;
  external: boolean;
}

interface Config {
  server: Server;
  scrapeInterval: number;
  hosts: Hosts[];
  alert: Alert;
}

let config_sample = `{
  "scrapeInterval": 30,
  "hosts": [
    {
      "ip": "127.0.0.1",
      "port": "2375"
    },
    {
      "ip": "84.247.176.100",
      "port": "2375",
      "ca_path": "/etc/docker/84.247.176.100/ca.pem",
      "cert_path": "/etc/docker/84.247.176.100/cert.pem",
      "key_path": "/etc/docker/84.247.176.100/key.pem"
    }
  ],
  "alert": {
    "enable": true,
    "webhook": [
      "https://open.larksuite.com/open-apis/bot/v2/hook/*"
    ],
    "email": {
      "smtp_endpoint": "smtpdm-ap-1.aliyuncs.com",
      "smtp_port": 80,
      "smtp_account": "autome@cess.cloud",
      "smtp_password": "***********",
      "receiver": [
        "z1092280043@gmail.com"
      ]
    }
  }
}`;

let sample = `{
  "alert": {
    "email": {
      "receiver": [
        "string"
      ],
      "senderAddr": "string",
      "smtpEndpoint": "string",
      "smtpPassword": "string",
      "smtpPort": number
    },
    "enable": boolean,
    "webhook": [
      "string"
    ]
  },
  "hosts": [
    {
      "capath": "string",
      "certPath": "string",
      "ip": "string",
      "keyPath": "string",
      "port": "string"
    }
  ],
  "scrapeInterval": number
}`;

export default function Page() {

  const [isValidJSON, setIsValidJSON] = useState<boolean>(true);
  const [configString, setConfigString] = useState<string>("");
  const [errorMessage, setErrorMessage] = useState<string>("");

  const refreshConfig = useCallback(async () => {
    try {
      const response = await fetch(`${getApiServerUrl()}/config`, {
        method: "GET",
      });

      if (!response.ok) {
        throw new Error(
          "Server responded with an error. Please check the server status or contact support."
        );
      }

      let config_with_server: Config = await response.json();
      let { server, ...config } = config_with_server;
      setConfigString(JSON.stringify(config, null, 2));
    } catch (error) {
      setConfigString(config_sample);
      console.error("Failed to fetch data:", error);
    }
  }, []);
  

  useEffect(() => {
    refreshConfig();
  }, []);

  useEffect(() => {
    if (configString.length > 0) {
      try {
        JSON.parse(configString);
        setIsValidJSON(true);
        setErrorMessage("");
      } catch (e) {
        if (e instanceof Error) {
          setErrorMessage(e.message);
        }
        setIsValidJSON(false);
      }
    }
  }, [configString]);

  const saveConfigHandler = async () => {
    try {
      const response = await fetch(`${getApiServerUrl()}/config`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body:JSON.stringify(JSON.parse(configString)),
      });

      if (!response.ok) {
        throw new Error(
          "Server responded with an error. Please check the server status or contact support."
        );
      }
      console.log("toast");
      toast.success("Configuration saved!");
      const data = await response.json();
      return data;
    } catch (error) {
      console.error("Failed to fetch data:", error);
    }
  };

  return (
    <div className="pl-12 pr-4 bg-white dark:bg-gray-900">
      <section className="pr-4 bg-white dark:bg-gray-900">
        <div className="flex flex-col md:flex-row py-8 px-4 mx-auto max-w-full lg:pt-16">
          <div className="basis-1/2 p-4">
            <div className="mb-3 block">
              <Label htmlFor="comment" value="Host Configuration" />
            </div>
            <Textarea
              id="comment"
              className={`${
                isValidJSON
                  ? "focus:border-green-500 focus:ring-green-500"
                  : "focus:border-red-500 focus:ring-red-500"
              }`}
              placeholder="Edit Configuration..."
              required
              rows={20}
              value={configString}
              onChange={(e) => setConfigString(e.target.value)}
            />
            <div className="flex justify-start pt-4">
              <label className="text-red-500">{errorMessage}</label>
            </div>
            <div className="flex justify-end pt-4">
              <Button onClick={saveConfigHandler}>Save</Button>
            </div>
          </div>
          <div className="basis-1/2 p-4 h-full">
            <div className="mb-3 block">
              <Label htmlFor="comment" value="Sample Configuration" />
            </div>
            <Textarea
              id="comment"
              className="bg-gray-300 resize-none"
              placeholder="Leave a comment..."
              required
              rows={35}
              value={sample}
              readOnly
            />
          </div>
        </div>
      </section>
      <Toaster richColors position="top-right"/>
    </div>
  );
}
