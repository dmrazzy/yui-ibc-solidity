// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Height} from "../../proto/Client.sol";
import {ConnectionEnd} from "../../proto/Connection.sol";
import {Channel, Upgrade} from "../../proto/Channel.sol";
import {IBCChannelLib} from "../04-channel/IBCChannelLib.sol";
import {IIBCModule, IIBCModuleInitializer} from "../26-router/IIBCModule.sol";

interface IIBCQuerier {
    function getCommitmentPrefix() external view returns (bytes memory);

    function getRevisionNumber() external view returns (uint64);

    function getCommitmentsSlot() external pure returns (bytes32);

    function getCommitment(bytes32 hashedPath) external view returns (bytes32);

    function getExpectedTimePerBlock() external view returns (uint64);

    function getIBCModuleByPort(string calldata portId) external view returns (IIBCModuleInitializer);

    function getIBCModuleByChannel(string calldata portId, string calldata channelId)
        external
        view
        returns (IIBCModule);

    function getClientByType(string calldata clientType) external view returns (address);

    function getClientType(string calldata clientId) external view returns (string memory);

    function getClient(string calldata clientId) external view returns (address);

    function getClientState(string calldata clientId) external view returns (bytes memory, bool);

    function getConsensusState(string calldata clientId, Height.Data calldata height)
        external
        view
        returns (bytes memory, bool);

    function getConnection(string calldata connectionId) external view returns (ConnectionEnd.Data memory, bool);

    function getChannel(string calldata portId, string calldata channelId)
        external
        view
        returns (Channel.Data memory, bool);

    function getNextSequenceSend(string calldata portId, string calldata channelId) external view returns (uint64);

    function getNextSequenceRecv(string calldata portId, string calldata channelId) external view returns (uint64);

    function getNextSequenceAck(string calldata portId, string calldata channelId) external view returns (uint64);

    function getPacketReceipt(string calldata portId, string calldata channelId, uint64 sequence)
        external
        view
        returns (IBCChannelLib.PacketReceipt);

    function getChannelUpgrade(string calldata portId, string calldata channelId)
        external
        view
        returns (Upgrade.Data memory, bool);

    function getCanTransitionToFlushComplete(string calldata portId, string calldata channelId)
        external
        view
        returns (bool);
}
