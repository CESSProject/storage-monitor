"use client";
import React, {useEffect, useState} from 'react';
import {MinusCircleOutlined, PlusOutlined} from '@ant-design/icons';
import {Button, Card, Form, Input, InputNumber, Switch, Typography} from 'antd';
import {getApiServerUrl} from "@/utils";
import {toast} from "sonner";
import axios from "axios";

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

const {Title} = Typography;

export default function Page() {
    const [form] = Form.useForm();
    const [errorMessage, setErrorMessage] = useState<string>("");
    const [config, setConfig] = useState<Config | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`${getApiServerUrl()}/config`, {});
                if (!response.data) throw new Error(
                    "Server responded with an error. Please check the server status or contact support."
                );
                let config: Config = await response.data;
                setConfig(config);
                form.setFieldsValue(config);
                return "ok"
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };
        fetchData().then(r => console.log(r));
    }, [form]);

    useEffect(() => {
        if (config) {
            form.setFieldsValue(config);
        }
    }, [config, form]);

    const onValuesChange = (allValues: Config) => {
        setConfig(allValues);
    };

    const onFinish = async (values: Config) => {
        try {
            const response = await axios.post(`${getApiServerUrl()}/config`, values, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });
            if (!response.data) {
                throw new Error(
                    "Server responded with an error. Please check the server status or contact support."
                );
            }
            toast.success("Configuration saved!");
            // return response.data;
        } catch (error) {
            console.error("Failed to fetch data:", error);
        }
    };

    return (
        <div style={{
            maxWidth: '800px',
            margin: '0 auto',
            padding: '24px',
            backgroundColor: '#f0f2f5'
        }}>
            <Card>
                <Title level={2} style={{textAlign: 'center', marginBottom: '24px'}}>
                    Watchdog Configuration
                </Title>
                <Form form={form}
                      onValuesChange={onValuesChange}
                      onFinish={onFinish}
                      layout="vertical">
                    <Form.Item>
                        <Form.Item name={['server', 'Http', 'http_port']} label="HTTP Port">
                            <InputNumber disabled/>
                        </Form.Item>
                        <Form.Item name={['server', 'external']} label="External" valuePropName="checked">
                            <Switch disabled/>
                        </Form.Item>
                    </Form.Item>

                    <Form.Item name="scrapeInterval" label="Scrape Interval">
                        <InputNumber/>
                    </Form.Item>

                    <Form.Item label="Hosts">
                        <Form.List name="hosts">
                            {(fields, {add, remove}) => (
                                <>
                                    {fields.map(({key, name, ...restField}) => (
                                        <div key={key} style={{display: 'flex', marginBottom: 8}}>
                                            <Form.Item {...restField} name={[name, 'IP']}
                                                       rules={[{required: true, message: 'IP is required'}]}>
                                                <Input placeholder="IP"/>
                                            </Form.Item>
                                            <Form.Item {...restField} name={[name, 'Port']}
                                                       rules={[{required: true, message: 'Port is required'}]}>
                                                <Input placeholder="Port"/>
                                            </Form.Item>
                                            <Form.Item {...restField} name={[name, 'CAPath']}>
                                                <Input placeholder="CA Path"/>
                                            </Form.Item>
                                            <Form.Item {...restField} name={[name, 'CertPath']}>
                                                <Input placeholder="Cert Path"/>
                                            </Form.Item>
                                            <Form.Item {...restField} name={[name, 'KeyPath']}>
                                                <Input placeholder="Key Path"/>
                                            </Form.Item>
                                            <MinusCircleOutlined onClick={() => remove(name)}/>
                                        </div>
                                    ))}
                                    <Form.Item>
                                        <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined/>}>
                                            Add Host
                                        </Button>
                                    </Form.Item>
                                </>
                            )}
                        </Form.List>
                    </Form.Item>

                    <Form.Item label="Alert Configuration">
                        <Form.Item name={['alert', 'enable']} valuePropName="checked">
                            <Switch/>
                        </Form.Item>
                        <Form.Item label="Webhooks">
                            <Form.List name={['alert', 'webhook']}>
                                {(fields, {add, remove}) => (
                                    <>
                                        {fields.map(({key, name, ...restField}) => (
                                            <div key={key} style={{display: 'flex', marginBottom: 8}}>
                                                <Form.Item {...restField} name={name}
                                                           rules={[{required: true, message: 'Webhook is required'}]}>
                                                    <Input placeholder="Webhook URL"/>
                                                </Form.Item>
                                                <MinusCircleOutlined onClick={() => remove(name)}/>
                                            </div>
                                        ))}
                                        <Form.Item>
                                            <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined/>}>
                                                Add Webhook Url
                                            </Button>
                                        </Form.Item>
                                    </>
                                )}
                            </Form.List>
                        </Form.Item>
                        <Form.Item label="Email Configuration">
                            <Form.Item name={['alert', 'Email', 'smtp_endpoint']} label="SMTP Endpoint">
                                <Input/>
                            </Form.Item>
                            <Form.Item name={['alert', 'Email', 'smtp_port']} label="SMTP Port">
                                <InputNumber/>
                            </Form.Item>
                            <Form.Item name={['alert', 'Email', 'smtp_account']} label="SMTP Account">
                                <Input disabled/>
                            </Form.Item>
                            <Form.Item name={['alert', 'Email', 'smtp_password']} label="SMTP Password">
                                <Input.Password disabled/>
                            </Form.Item>
                            <Form.Item label="Receivers">
                                <Form.List name={['alert', 'Email', 'receiver']}>
                                    {(fields, {add, remove}) => (
                                        <>
                                            {fields.map(({key, name, ...restField}) => (
                                                <div key={key} style={{display: 'flex', marginBottom: 8}}>
                                                    <Form.Item {...restField} name={name}
                                                               rules={[{
                                                                   required: true,
                                                                   message: 'Receiver email is required'
                                                               }]}>
                                                        <Input placeholder="Receiver Email"/>
                                                    </Form.Item>
                                                    <MinusCircleOutlined onClick={() => remove(name)}/>
                                                </div>
                                            ))}
                                            <Form.Item>
                                                <Button type="dashed" onClick={() => add()} block
                                                        icon={<PlusOutlined/>}>
                                                    Add Receiver
                                                </Button>
                                            </Form.Item>
                                        </>
                                    )}
                                </Form.List>
                            </Form.Item>
                        </Form.Item>
                    </Form.Item>

                    <Form.Item>
                        <Button type="primary" htmlType="submit" style={{ position: 'absolute', right: 10, bottom: 10 }}>
                            Save
                        </Button>
                    </Form.Item>
                </Form>
            </Card>
        </div>
    );
}
