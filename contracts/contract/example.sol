pragma solidity ^0.4.22;

contract example {
    
    address public sender;
    uint256 public times;
    
    event Test(address sender, uint256 times);
    
    function test() {
        sender = msg.sender;
        times++;
        Test(sender, times);
    }
}