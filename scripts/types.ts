// types.ts

export interface TerraformConfig {
  provider: string;
  region: string;
  accessKey: string;
  secretKey: string;
}

export interface InfraResource {
  id: string;
  name: string;
  type: string;
  status: string;
}

export enum InfraResourceStatus {
  PENDING = 'pending',
  RUNNING = 'running',
  STOPPED = 'stopped',
  ERROR = 'error',
}

export interface InfraTerraformState {
  version: string;
  terraformVersion: string;
  serial: number;
  lineage: string;
  outputs: { [key: string]: any };
  resources: InfraResource[];
}

export interface InfraTerraformOutput {
  name: string;
  value: any;
  description: string;
  sensitive: boolean;
}