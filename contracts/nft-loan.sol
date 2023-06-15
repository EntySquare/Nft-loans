pragma solidity >=0.8.0 <0.9.0;

import "./IERC721.sol";

contract NFTLoan {
    address  owner;
    struct NftLoasInfo {
        uint256 tokenId;
        uint loanTime;
        uint256 flag;            
    }

    IERC721 public nftContract;
    mapping (address => NftLoasInfo[]) public loans;
    mapping (address => uint256) public loansNumber;
    
constructor(address _owner) {
    owner = _owner;
    
}
// //生成通行卡
// function mintPassCard() external onlyManager {

// }
// //合成通行卡
// function composePassCard() external{
  
// }
// //购买通行卡
// function buyPassCard() external{
     

// }
// //转移通行卡
// function shiftPassCard() external{
     

// }
// //授权通行卡（租借功能）
// function approvePasscard(uint _cardId,address to) external{
//      address owner = passcard[_cardId];
//         require(
//             msg.sender == owner,
//             "not owner for you"
//         );
//         _approve(owner, to, tokenId);
// }
// function _approve(【
//         address owner,
//         address to,
//         uint cardId
//     ) private {
//         passcardApprovals[_cardId] = to;
//     }
//抵押nft
function depositNFT(uint256  _tokenId,address _nftContract) external payable {

                nftContract = IERC721(_nftContract);
                require(nftContract.ownerOf(_tokenId) == msg.sender, "You don't own this NFT");
                NftLoasInfo memory newLoans;
                uint256 nowNumer = loansNumber[msg.sender];
                newLoans.tokenId = _tokenId;
                newLoans.loanTime = block.timestamp;
                newLoans.flag = 1;
                loans[msg.sender][nowNumer] = newLoans;
                loansNumber[msg.sender] += 1;
                nftContract.approve(address(this), _tokenId);
      
}
//赎回nft
function withdrawNFT(uint256 _tokenId,address _nftContract) external payable {
    nftContract = IERC721(_nftContract);
    uint256 nowNumer = loansNumber[msg.sender];
    for (uint i = 0; i < nowNumer; ++i ) {
        if (loans[msg.sender][i].tokenId == _tokenId) {
            require(loans[msg.sender][i].flag == 1, "You already have withdraw");   
             loans[msg.sender][_tokenId].flag  = 0;
             nftContract.approve(msg.sender, _tokenId);
        }
    }
   
}
function ownerOf(uint256  _tokenId,address _nftContract) public returns (address){
                nftContract = IERC721(_nftContract);
                return nftContract.ownerOf(_tokenId);
      
}
function loansList(uint start,uint limit,address _nftContract) public returns (uint256[] memory){
    nftContract = IERC721(_nftContract);
    uint256 nowNumer = loansNumber[msg.sender];
    uint l = 0;
    uint256[] memory list;
    for (uint i = 0; i < nowNumer; ++i ) {
            if(loans[msg.sender][i].flag == 1){
                if (l >= start && l < start + limit){
                list[l] = loans[msg.sender][i].tokenId;
                l++;
                }
                if(l > start + limit){
                    return list;
                }
            }   
           
    }
    return list;
}
function loansCount() public returns (uint256){
    uint256 nowNumer = loansNumber[msg.sender];
    for (uint i = 0; i < nowNumer; ++i ) {
            if(loans[msg.sender][i].flag == 0){
                nowNumer--;
                }
                 
           
    }
    return nowNumer;
}

    modifier onlyManager() {
        require(
            msg.sender == owner,
            "Only owner can call this."
        );
        _;
    }
}


