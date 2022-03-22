pragma solidity^0.8.7;
contract hello {
    string mymsg;
    constructor(string memory _msg) {
        mymsg = _msg;
    }
    function getMsg() public view returns (string memory) {
        return mymsg;
    }
    function setMsg(string memory _msg) public {
        mymsg = _msg;
    }
}
