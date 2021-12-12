import { Dispatch, SetStateAction } from 'react';

export type SetState<T> = Dispatch<SetStateAction<T>>;

export interface DetailPayload {
  arn: string;
  region: string;
}

export interface SecretDetail {
  ARN: string;
  CreatedDate: string;
  DeletedDate?: string;
  Description?: string;
  KmsKeyId?: string;
  LastAccessedDate?: string;
  LastChangedDate?: string;
  LastRotatedDate?: string;
  Name: string;
  OwningService: unknown;
  RotationEnabled: boolean;
  RotationLambdaARN: string;
  RotationRules: {
    AutomaticallyAfterDays: number;
  };
  SecretVersionsToStages?: {
    [key: string]: string[];
  };
  VersionIdsToStages?: {
    [key: string]: string[];
  };
  Tags: {
    Key: string;
    Value: string;
  }[];
}

export interface SecretValue {
  ARN: string;
  CreatedDate: string;
  Name: string;
  SecretBinary?: string;
  SecretString?: string;
  VersionId: string;
  VersionStages: string[];
}
