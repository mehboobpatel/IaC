import { Stack, StackProps, CfnOutput } from 'aws-cdk-lib';
import { Construct } from 'constructs';

import * as ec2 from 'aws-cdk-lib/aws-ec2';
import * as ssm from 'aws-cdk-lib/aws-ssm';

import { SSM_PREFIX } from '../../config';

const ACCOUNT: string = '683627797067';
const REGION: string = 'us-east-1';

export class VpcStack extends Stack {
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, {
            ...props,
            env: {
                account: ACCOUNT,
                region: REGION,
            },
        });

        // Retrieve values from SSM parameters
        const vpcId = ssm.StringParameter.valueFromLookup(this, ${SSM_PREFIX}/JAWSvpcID);
        const subnet1Id = ssm.StringParameter.valueFromLookup(this, ${SSM_PREFIX}/JAWSsubnet1);
        const subnet2Id = ssm.StringParameter.valueFromLookup(this, ${SSM_PREFIX}/JAWSsubnet2);
        const region = ssm.StringParameter.valueFromLookup(this, ${SSM_PREFIX}/JAWSregion);

        // Define the VPC
        const jvpc = ec2.Vpc.fromVpcAttributes(this, 'JenkinsVPC', {
            vpcId: vpcId,
            availabilityZones: [${region}c, ${region}d],
            privateSubnetIds: [subnet1Id],
            isolatedSubnetIds: [subnet2Id],
        });

        // Define other resources using retrieved values
        // You can create route tables, network ACLs, security groups, etc. based on your requirements.

        // Example CfnOutput
        new CfnOutput(this, 'JenkinsVPCOutput', { value: vpcId });
    }
}
