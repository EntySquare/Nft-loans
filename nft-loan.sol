pragma solidity >=0.8.0 <0.9.0;

import "./IERC721.sol";

contract NFTLoan {
    address public owner;
    struct NftLoasInfo {
        uint256 loanAmount;
        uint256 loanDeadLine;            
    }
    uint256 interestRate;
    IERC721 public nftContract;
    mapping (address => mapping(uint256=>NftLoasInfo)) public loans;
    
constructor(address _owner,uint256 _interestRate) {
    owner = _owner;
    interestRate =  _interestRate;
}
function depositNFT(uint256 _tokenId,uint256 _loanAmount,uint256 _loanDeadLine,address _nftContract) public {
    require(nftContract.ownerOf(_tokenId) == msg.sender, "You don't own this NFT");
    loans[msg.sender][_tokenId].loanAmount = _loanAmount;
    loans[msg.sender][_tokenId].loanDeadLine = _loanDeadLine;
    nftContract = IERC721(_nftContract);
    nftContract.transferFrom(msg.sender, address(this), _tokenId);
}
function withdrawNFT(uint256 _tokenId,address _nftContract) public {
    require(loans[msg.sender][_tokenId].loanAmount != 0, "You don't have any NFT deposited");
    require(loans[msg.sender][_tokenId].loanDeadLine < block.timestamp, "Your NFT loan has out of date");
    nftContract = IERC721(_nftContract);
    loans[msg.sender][_tokenId].loanAmount  = 0;
    loans[msg.sender][_tokenId].loanDeadLine  = 0;
    nftContract.transferFrom(address(this), msg.sender, _tokenId);
}

}

