// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../../contracts/core/02-client/IIBCClientErrors.sol";
import "../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../contracts/proto/MockClient.sol";
import "../../../contracts/clients/mock/MockClient.sol";
import "./helpers/TestableIBCHandler.t.sol";
import "./helpers/IBCTestHelper.t.sol";
import "./helpers/MockClientTestHelper.t.sol";
import {IIBCHostErrors} from "../../../contracts/core/24-host/IIBCHostErrors.sol";

contract TestICS02 is Test, MockClientTestHelper {
    using IBCHeight for Height.Data;

    function testRegisterClient() public {
        TestableIBCHandler handler = defaultIBCHandler();
        MockClient mockClient = new MockClient(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        handler.registerClient("testtes", mockClient);
    }

    function testRegisterClientDuplicatedClientType() public {
        TestableIBCHandler handler = defaultIBCHandler();
        MockClient mockClient = new MockClient(address(handler));
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
        vm.expectRevert(
            abi.encodeWithSelector(IIBCHostErrors.IBCHostClientTypeAlreadyExists.selector, MOCK_CLIENT_TYPE)
        );
        handler.registerClient(MOCK_CLIENT_TYPE, mockClient);
    }

    function testRegisterClientInvalidClientType() public {
        TestableIBCHandler handler = defaultIBCHandler();
        vm.expectRevert(abi.encodeWithSelector(IIBCHostErrors.IBCHostInvalidLightClientAddress.selector, address(0)));
        handler.registerClient(MOCK_CLIENT_TYPE, ILightClient(address(0)));

        MockClient mockClient = new MockClient(address(handler));

        // empty client type
        vm.expectRevert(abi.encodeWithSelector(IIBCHostErrors.IBCHostInvalidClientType.selector, ""));
        handler.registerClient("", mockClient);

        // too short client type
        vm.expectRevert(abi.encodeWithSelector(IIBCHostErrors.IBCHostInvalidClientType.selector, "testte"));
        handler.registerClient("testte", mockClient);

        // first character is not a letter
        vm.expectRevert(abi.encodeWithSelector(IIBCHostErrors.IBCHostInvalidClientType.selector, "-mocktest"));
        handler.registerClient("-mocktest", mockClient);

        // last character is not a letter or number
        vm.expectRevert(abi.encodeWithSelector(IIBCHostErrors.IBCHostInvalidClientType.selector, "mocktest-"));
        handler.registerClient("mocktest-", mockClient);
    }

    function testHeightToUint128(Height.Data memory height) public pure {
        Height.Data memory actual = IBCHeight.fromUint128(IBCHeight.toUint128(height));
        assert(height.eq(actual));
    }

    function testCreateClient() public {
        (TestableIBCHandler handler, MockClient mockClient) = ibcHandlerMockClient();
        {
            string memory clientId = handler.createClient(msgCreateMockClient(1));
            assertEq(clientId, mockClientId(0));
            assertEq(handler.getClientType(clientId), MOCK_CLIENT_TYPE);
            assertEq(handler.getClient(clientId), address(mockClient));
            assertFalse(handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId)) == bytes32(0));
            assertFalse(handler.getCommitment(IBCCommitment.consensusStateCommitmentKey(clientId, 0, 1)) == bytes32(0));
        }
        {
            string memory clientId = handler.createClient(msgCreateMockClient(100));
            assertEq(clientId, mockClientId(1));
            assertEq(handler.getClientType(clientId), MOCK_CLIENT_TYPE);
            assertEq(handler.getClient(clientId), address(mockClient));
            assertFalse(handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId)) == bytes32(0));
            assertFalse(
                handler.getCommitment(IBCCommitment.consensusStateCommitmentKey(clientId, 0, 100)) == bytes32(0)
            );
        }
    }

    function testInvalidCreateClient() public {
        (TestableIBCHandler handler,) = ibcHandlerMockClient();
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.clientType = "";
            vm.expectRevert(abi.encodeWithSelector(IIBCClientErrors.IBCClientUnregisteredClientType.selector, ""));
            handler.createClient(msg_);
        }
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.clientType = "06-solomachine";
            vm.expectRevert(
                abi.encodeWithSelector(IIBCClientErrors.IBCClientUnregisteredClientType.selector, "06-solomachine")
            );
            handler.createClient(msg_);
        }
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.protoClientState = abi.encodePacked(msg_.protoClientState, hex"00");
            vm.expectRevert();
            handler.createClient(msg_);
        }
        {
            IIBCClient.MsgCreateClient memory msg_ = msgCreateMockClient(1);
            msg_.protoConsensusState = abi.encodePacked(msg_.protoConsensusState, hex"00");
            vm.expectRevert();
            handler.createClient(msg_);
        }
    }

    function testUpdateClient() public {
        bytes32 prevClientStateCommitment;
        (TestableIBCHandler handler,) = ibcHandlerMockClient();
        string memory clientId = handler.createClient(msgCreateMockClient(1));
        prevClientStateCommitment = handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId));

        {
            handler.updateClient(msgUpdateMockClient(clientId, 2));
            bytes32 commitment = handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId));
            assertTrue(
                commitment != prevClientStateCommitment && commitment != bytes32(0), "commitment should be updated"
            );
            prevClientStateCommitment = commitment;
        }
        {
            handler.updateClient(msgUpdateMockClient(clientId, 3));
            bytes32 commitment = handler.getCommitment(IBCCommitment.clientStateCommitmentKey(clientId));
            assertTrue(
                commitment != prevClientStateCommitment && commitment != bytes32(0), "commitment should be updated"
            );
            prevClientStateCommitment = commitment;
        }
        {
            // update with the same height should not revert
            handler.updateClient(msgUpdateMockClient(clientId, 3));
            bytes32 prevConsensusStateCommitment = keccak256(mockConsensusState(getBlockTimestampNano()));
            // update with the same height and different consensus state should revert
            uint256 prev = vm.getBlockTimestamp();
            vm.warp(prev + 1);
            bytes32 newConsensusStateCommitment = keccak256(mockConsensusState(getBlockTimestampNano()));
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 3);
            vm.expectRevert(abi.encodeWithSelector(
                IIBCClientErrors.IBCClientInconsistentConsensusStateCommitment.selector,
                IBCCommitment.consensusStateCommitmentKey(clientId, 0, 3),
                newConsensusStateCommitment,
                prevConsensusStateCommitment
            ));
            handler.updateClient(msg_);
            vm.warp(prev);
        }
    }

    function testInvalidUpdateClient() public {
        (TestableIBCHandler handler,) = ibcHandlerMockClient();
        string memory clientId = handler.createClient(msgCreateMockClient(1));
        assertEq(clientId, mockClientId(0));
        {
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 2);
            msg_.clientId = "";
            vm.expectRevert();
            handler.updateClient(msg_);
        }
        {
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 2);
            msg_.clientId = mockClientId(1);
            vm.expectRevert();
            handler.updateClient(msg_);
        }
        {
            IIBCClient.MsgUpdateClient memory msg_ = msgUpdateMockClient(clientId, 2);
            msg_.protoClientMessage = abi.encodePacked(msg_.protoClientMessage, hex"00");
            vm.expectRevert();
            handler.updateClient(msg_);
        }
    }
}
