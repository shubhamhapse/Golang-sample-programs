pragma solidity >= 0.5.0;

contract SignatureValidator {
    struct Person {
        address addr;
        bytes32 name;
        bytes32 id;
    }
    address private owner;

    //map to maintain signed data
    mapping (bytes32=>bool) private signedDataHash;

    constructor() public {
        owner=msg.sender;
    }

    //split signature into s,r,v variables
    function splitSignatureInRSV(bytes memory sig) internal pure returns (bool, uint8, bytes32, bytes32){

        if (sig.length != 65){
            return (false,0,0,0);
        }
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
        return (true,v, r, s);
    }

    //setter function
    function setPerson(bytes memory data, bytes memory signatureBytes) public returns (bool){
        bytes32 r;
        bytes32 s;
        uint8 v;
        bool status;
        (status, v,r,s)=splitSignatureInRSV(signatureBytes);

        if (status == false){
            //SIG length is not 65
            return false;
        }

        bytes32 dataHash=keccak256(data);

        if (owner==ecrecover(dataHash,v, r, s)){
            signedDataHash[dataHash]=true;
            return true;
        }
        else {
            return false;
        }
    }

    function yourTransaction(bytes memory data) view public {
        if (isSigned(data)==true) {
            //data is signed, write your code here
        }
        else {
            //data is not signed
        }
    }

    //Check if data is signed or not
    function isSigned(bytes memory data) public view returns (bool){
        return signedDataHash[keccak256(data)];
    }


}
