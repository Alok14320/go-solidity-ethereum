// MyToken.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract MintTokenMTK is ERC721 {
    uint256 public mintPrice = 1;
    uint256 public totalMint = 1;

    mapping(address => uint256) public walletMints;

    constructor() ERC721("myToken", "MTK") {}

    function MintMTK(uint256 quantity) public payable {
        require(msg.value != 0, "invalid amount");
        require(quantity * mintPrice == msg.value);
        walletMints[msg.sender] += quantity;
        // _safeMint(msg.sender, totalMint);

        totalMint++;
    }

    function CheckTotalMint() public view returns (uint256) {
        return walletMints[msg.sender];
    }
}
