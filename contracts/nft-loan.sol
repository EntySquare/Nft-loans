pragma solidity >=0.8.0 <0.9.0;

import "./IERC721.sol";
import "./IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
// @author yueliyangzi
contract ngt{
    string public  name;
    string public symbol;
    uint8  public decimals = 4;
    uint256 public  totalSupply = 10000000000;
    mapping (address => uint256)  balances;
    mapping (address => mapping (address => uint256)) allowed;
    event Transfer(address owner,address spender,uint256 value);
    event Approval(address owner,address spender,uint256 value);
    event ChangeMarketStatusEvent(uint8 status);
    
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }
    function approve(address _spender, uint256 _value) public returns (bool success)   
    { 
         require((_value == 0) || (allowed[msg.sender][_spender] == 0));
        allowed[msg.sender][_spender] = _value;
       emit Approval(msg.sender, _spender, _value);
        return true;
    }
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowed[_owner][_spender];//允许_spender从_owner中转出的token数
    }
    function viewtotalSupply() public view returns (uint256){
        return totalSupply;
    }
}
// @author yueliyangzi
contract NGT is ngt,IERC721Receiver{
    using SafeMathCell for uint256;
    address foundation;
    address  owner;
     struct TokenData {
        uint256 tokenId;
        uint256 receivedTime;
        uint256 flag;
    }

    mapping(address => TokenData[]) public receivedTokens;

    event TokenReceived(address indexed from, uint256 indexed tokenId, uint256 receivedTime);

    IERC721 public nftContract;
    
   

     /* Initializes contract with initial supply tokens to the creator of the contract */
    //@notice Contract initial setting
     constructor(
        address _owner,address _nftContract,address _fund)  public {
        uint256 totalSupply = totalSupply * 10 ** uint256(decimals); // Update total supply
        balances[_owner] += totalSupply;  
        foundation = _fund;                     // Give the creator all initial tokens
        name = "NGT";                                      // Set the name for display purposes
        symbol = "NGT";  
        owner = _owner;
        nftContract = IERC721(_nftContract);                                 // Set the symbol for display purposes
    }
    function transfer(address _to, uint256 _value) public payable  returns (bool success){
         uint256 tax = (_value * 0).div(100);
         bool flag = _transfer(foundation,tax,0);
         require(flag,"transfer to foundation flase");
         return _transfer(_to,_value,tax);

    }
    function transferFrom(address _from, address _to, uint256 _value) public payable  returns (bool success){
         uint256 tax = (_value * 0).div(100);
         bool flag = _transferFrom(_from,foundation,tax,0);
         require(flag,"transfer to foundation flase");
         return _transferFrom(_from,_to,_value,tax);
    }
    function _transfer(address _to, uint256 _value,uint256 tax) internal returns (bool success){
        require(balances[msg.sender] >= _value && balances[_to] + _value > balances[_to],"Insufficient funds");
        
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value      
        
        balances[_to] += _value.sub(tax);//往接收账户增加token数量_value
       
        emit Transfer(msg.sender, _to, _value);//触发转币交易事件
       
        return true;
    }
    function _transferFrom(address _from, address _to, uint256 _value,uint256 tax) internal returns 
    (bool success) {
        require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value,"Insufficient funds");
       
        balances[_from] -= _value; //支出账户_from减去token数量_value
        
        balances[_to] += _value.sub(tax);//接收账户增加token数量_value
       
        allowed[_from][msg.sender] -= _value;//消息发送者可以从账户_from中转出的数量减少_value
        emit Transfer(_from, _to, _value);//触发转币交易事件
        return true;
    }
  //@notice burn function internal
  function _burn(uint256 amount) internal {
        require(amount != 0);
        require(amount <= totalSupply);
        totalSupply -= amount;
  }     
function onERC721Received(address operator, address from, uint256 tokenId, bytes calldata data) external override returns(bytes4) {
        // // tokenData[tokenId] = TokenData(from, block.timestamp);
        //  NftLoas Info memory newLoans;
        // //  uint256 nowNumer = loansNumber[from];
        //  newLoans.tokenId = tokenId;
        // //  newLoans.loanTime = block.timestamp;
        // //  newLoans.flag = 1;
        // //  loans[from][nowNumer] = newLoans;
        // //  loansNumber[from] += 1;
        receivedTokens[from].push(TokenData(tokenId, block.timestamp,1));
        emit TokenReceived(from, tokenId, block.timestamp);
        return this.onERC721Received.selector;
}

//赎回nft
function withdrawNFT(address to,uint256 _tokenId) onlyManager external payable {
    for (uint i = 0; i < receivedTokens[to].length; ++i ) {
        if (receivedTokens[to][i].tokenId == _tokenId) {
            require(receivedTokens[to][i].flag == 1, "You already have withdraw");   
            TokenData storage td = receivedTokens[to][i];
             td.flag  = 0;
             nftContract.transferFrom(address(this),to,_tokenId);
        }
    }
    
   
}
function ownerOf(uint256  _tokenId) public returns (address){
            
                return nftContract.ownerOf(_tokenId);
      
}
function loansList(address user) public view returns (TokenData[] memory){
    
    return receivedTokens[user];
}
function loansCount(address user) public view returns (uint256){
    uint256 nowNumer = receivedTokens[user].length;
    for (uint i = 0; i < receivedTokens[user].length; ++i ) {
            if(receivedTokens[user][i].flag == 0){
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

// @title cell library
// @author yueliyangzi
library SafeMathCell {
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
 
        uint256 c = a * b;
        require(c / a == b, "SafeMath:multiplication overflow");
        return c;
    }
 

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "SafeMath:division overflow");
        uint256 c = a / b;
        return c;
    }
 

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "SafeMath: subtraction overflow");
        uint256 c = a - b;
 
        return c;
    }
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");
 
        return c;
    }

    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0, "SafeMath: mod overflow");
        return a % b;
    }
    // _type 1.买入 2.卖出
    function recommender_radio(uint256 _generation,uint256 _type) internal pure returns(uint256 ratio){
              if(_type == 1){
                  if(_generation == 1){
                   return 30;
                   }    
                  if(_generation == 2 ){
                   return 20;
                  }
                  if(_generation >= 3 && _generation <= 8){
                   return 5;
                  }

              }
              if(_type == 2){
                  if(_generation == 1){
                   return 20;
                  }
                  if(_generation == 2 ){
                   return 10;
                  }
              }
              
              
    }
}


