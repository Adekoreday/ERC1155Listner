pragma solidity >=0.4.20;

contract ERC1155 {
   event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 tokens);
   event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] tokens);
}