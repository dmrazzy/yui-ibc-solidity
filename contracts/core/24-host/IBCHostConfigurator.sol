// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {IBCClientLib} from "../02-client/IBCClientLib.sol";
import {ILightClient} from "../02-client/ILightClient.sol";
import {IBCModuleManager} from "../26-router/IBCModuleManager.sol";
import {IIBCModule} from "../26-router/IIBCModule.sol";
import {IIBCHostConfigurator} from "./IIBCHostConfigurator.sol";

/**
 * @dev IBCHostConfigurator is a contract that provides the host configuration.
 */
abstract contract IBCHostConfigurator is IIBCHostConfigurator, IBCModuleManager {
    function _setExpectedTimePerBlock(uint64 expectedTimePerBlock_) internal virtual {
        expectedTimePerBlock = expectedTimePerBlock_;
    }

    function _registerClient(string calldata clientType, ILightClient client) internal virtual {
        if (!IBCClientLib.validateClientType(bytes(clientType))) {
            revert IBCHostInvalidClientType(clientType);
        } else if (address(clientRegistry[clientType]) != address(0)) {
            revert IBCHostClientTypeAlreadyExists(clientType);
        }
        if (address(client) == address(this) || !Address.isContract(address(client))) {
            revert IBCHostInvalidLightClientAddress(address(client));
        }
        clientRegistry[clientType] = address(client);
    }

    function _bindPort(string calldata portId, IIBCModule moduleAddress) internal virtual {
        if (!validatePortIdentifier(bytes(portId))) {
            revert IBCHostInvalidPortIdentifier(portId);
        }
        if (address(moduleAddress) == address(this) || !Address.isContract(address(moduleAddress))) {
            revert IBCHostInvalidModuleAddress(address(moduleAddress));
        }
        claimCapability(portCapabilityPath(portId), address(moduleAddress));
    }
}
