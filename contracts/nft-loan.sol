pragma solidity >=0.8.0 <0.9.0;

import "./IERC721.sol";

contract NFTLoan {
    address  owner;
    struct NftLoasInfo {
        uint16 loanAmount;
        uint32 loanDeadLine;            
    }
    struct UserInfo {
       uint8 userClass;
       uint16 integral;
      
    }
    uint256 interestRateAAA; // */1000 day
    uint256 interestRateAA; // */1000 day
    uint256 interestRateA; // */1000 day
    IERC721 public nftContract;
    mapping (address => mapping(uint256=>NftLoasInfo)) public loans;
    mapping (uint => address)  public passcard;
    mapping(uint => address) public passcardApprovals;
    
constructor(address _owner) {
    owner = _owner;
    interestRateAAA =  7;
    interestRateAA =  6;
    interestRateA =  5;
}
//生成通行卡
function mintPassCard() external onlyManager {

}
//合成通行卡
function composePassCard() external{
  
}
//购买通行卡
function buyPassCard() external{
     

}
//转移通行卡
function shiftPassCard() external{
     

}
//授权通行卡（租借功能）
function approvePasscard(uint _cardId,address to) external{
     address owner = passcard[_cardId];
        require(
            msg.sender == owner,
            "not owner for you"
        );
        _approve(owner, to, tokenId);
}
function _approve(【
        address owner,
        address to,
        uint cardId
    ) private {
        passcardApprovals[_cardId] = to;
    }
//抵押nft
function depositNFT(uint256[] _tokenIds,address _nftContract,uint16 _loanAmount,uint32 _loanDeadLine) external payable {
     for (
            uint j = 0;
            j <= _tokenIds.length - 1;
            ++ j
            ) {
                _tokenId = _tokenIds[j];
                 require(nftContract.ownerOf(_tokenId) == msg.sender, "You don't own this NFT");
                loans[msg.sender][_tokenId].loanAmount = _loanAmount;
                loans[msg.sender][_tokenId].loanDeadLine = _loanDeadLine;
                nftContract = IERC721(_nftContract);
                nftContract.transferFrom(msg.sender, address(this), _tokenId);
            }
   
}
//赎回nft
function withdrawNFT(uint256 _tokenId,address _nftContract) public {
    require(loans[msg.sender][_tokenId].loanAmount != 0, "You don't have any NFT deposited");
    require(loans[msg.sender][_tokenId].loanDeadLine < block.timestamp, "Your NFT loan has out of date");
    nftContract = IERC721(_nftContract);
    loans[msg.sender][_tokenId].loanAmount  = 0;
    loans[msg.sender][_tokenId].loanDeadLine  = 0;
    nftContract.transferFrom(address(this), msg.sender, _tokenId);
}
    modifier onlyManager() {
        require(
            msg.sender == owner,
            "Only owner can call this."
        );
        _;
    }
}

