pragma solidity >= 0.5.0;

contract BinaryTree{

    uint private nodeID=0;
    struct Node{
        uint id;
        uint256 data;
        uint left;
        uint right;
        bool flag;
    }

    event printNode(uint id, uint256 data, uint left, uint right);
    event error(uint id,uint parent,string msg);

    mapping(uint=>Node) public nodeList;

    /*
    * addNode
    * input
    *   data     - int256
    *   parent   - uint parentId
    *   position - int -1 for left and 1 for right
    */
    function addNode(uint256 data,uint parent, int position ) public {

        bool isEmpty= isTreeEmpty();
        nodeID++;
        Node memory node=Node(nodeID,data,0,0,true);
        if (isEmpty){
            nodeList[nodeID]=node;
            return;
        }

        bool parentExists= checkIfParentExists(parent);

        //checking that parent does not have any preassigned left or right child
        //by checking nodeList[parent].left==0
        //currently not throwing error if left/right position is already accoupied
        if (parentExists) {
            if (position== -1 && nodeList[parent].left==0 ){
                nodeList[parent].left=nodeID;
            }
            else if (position ==1 && nodeList[parent].right==0) {
                nodeList[parent].right=nodeID;
            }
            nodeList[nodeID]=node;
        }
        else {
            emit error(nodeID,parent,"Parent is not present");
        }
    }

    //check if parentId is present or not
    function checkIfParentExists(uint parentID) internal view returns (bool){
        if (nodeList[parentID].flag == true){
            return true;
        }
        else {
            return false;
        }
    }

    //returns total no of nodes in a tree
    function getLastID() public view returns (uint) {
        return nodeID;
    }

    //check if tree is empty
    function isTreeEmpty() public view returns (bool){
        if (nodeID==0) {
            return true;
        }
        else {
            return false;
        }
    }

    //Print node details and there relationship with other nodes
    function printNodes() public  {
        for(uint i=1; i <= nodeID; i++){
            Node memory node= nodeList[i];
            emit printNode(node.id,node.data, node.left, node.right);
        }
    }

} 