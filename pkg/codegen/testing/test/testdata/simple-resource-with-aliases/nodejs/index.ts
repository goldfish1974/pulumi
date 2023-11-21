// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { BasicResourceArgs } from "./basicResource";
export type BasicResource = import("./basicResource").BasicResource;
export const BasicResource: typeof import("./basicResource").BasicResource = null as any;
utilities.lazyLoad(exports, ["BasicResource"], () => require("./basicResource"));

export { BasicResourceV2Args } from "./basicResourceV2";
export type BasicResourceV2 = import("./basicResourceV2").BasicResourceV2;
export const BasicResourceV2: typeof import("./basicResourceV2").BasicResourceV2 = null as any;
utilities.lazyLoad(exports, ["BasicResourceV2"], () => require("./basicResourceV2"));

export { BasicResourceV3Args } from "./basicResourceV3";
export type BasicResourceV3 = import("./basicResourceV3").BasicResourceV3;
export const BasicResourceV3: typeof import("./basicResourceV3").BasicResourceV3 = null as any;
utilities.lazyLoad(exports, ["BasicResourceV3"], () => require("./basicResourceV3"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "example:index:BasicResource":
                return new BasicResource(name, <any>undefined, { urn })
            case "example:index:BasicResourceV2":
                return new BasicResourceV2(name, <any>undefined, { urn })
            case "example:index:BasicResourceV3":
                return new BasicResourceV3(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("example", "index", _module)
pulumi.runtime.registerResourcePackage("example", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:example") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
