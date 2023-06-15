pragma solidity >=0.8.0 <0.9.0;

import "./IERC721.sol";
import "./IERC20.sol";
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
contract NGT is ngt {
    using SafeMathCell for uint256;
    using AddressArrayLimitOnee for address[31];
    using AddressArrayOnee for address[];
    uint256 constant exchange_rate_usdt = 100000;
    mapping(uint8 => TokenInfo) tokens;
    address foundation;
    mapping (address => AddressStatus) details;
    struct TokenInfo{
        string token_name;
        address token_address;
        uint256 exchange_rate; 
    }
    struct AddressStatus{
        uint256 locked_balances;
        uint256 avilable_balances;
        address recommender;
    }
    address  owner;
    struct NftLoasInfo {
        uint256 tokenId;
        uint loanTime;
        uint256 flag;            
    }
    IERC721 public nftContract;
    mapping (address => NftLoasInfo[]) public loans;
    mapping (address => uint256) public loansNumber;
    address nft;
   
    
     /* Initializes contract with initial supply tokens to the creator of the contract */
    //@notice Contract initial setting
     constructor(
        address _owner,address _nftContract,address _fund)  public {
        uint256 totalSupply = totalSupply * 10 ** uint256(decimals); // Update total supply
        balances[_owner] += totalSupply;  
        foundation = _fund;                     // Give the creator all initial tokens
        name = "ONEE";                                      // Set the name for display purposes
        symbol = "ONEE";  
        owner = _owner;
        nft = _nftContract;                                 // Set the symbol for display purposes
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
        require(details[msg.sender].avilable_balances >= _value,"Insufficient funds");
        balances[msg.sender] -= _value;//从消息发送者账户中减去token数量_value      
        details[msg.sender].avilable_balances -= _value;
        balances[_to] += _value.sub(tax);//往接收账户增加token数量_value
        details[_to].avilable_balances += _value.sub(tax);
        emit Transfer(msg.sender, _to, _value);//触发转币交易事件
       
        return true;
    }
    function _transferFrom(address _from, address _to, uint256 _value,uint256 tax) internal returns 
    (bool success) {
        require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value,"Insufficient funds");
        require(details[_from].avilable_balances >= _value,"Insufficient funds");
        balances[_from] -= _value; //支出账户_from减去token数量_value
        details[_from].avilable_balances -= _value;
        balances[_to] += _value.sub(tax);//接收账户增加token数量_value
        details[_to].avilable_balances += _value.sub(tax);
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
  //@notice bind user with recommender (one can only have one recommender)
  function bindRecommender(address _recommender) external returns(bool){
      require(details[msg.sender].recommender == address(0),"this address has been bound");
      _bind(msg.sender,_recommender);
      return true;
  }
  //@notice bind function internal
  function _bind(address _self,address _to) internal{
      details[_self].recommender = _to;
  }
//抵押nft
function depositNFT(uint256  _tokenId) external payable {

                nftContract = IERC721(nft);
                require(nftContract.ownerOf(_tokenId) == msg.sender, "You don't own this NFT");
                // NftLoasInfo memory newLoans;
                // uint256 nowNumer = loansNumber[msg.sender];
                // newLoans.tokenId = _tokenId;
                // newLoans.loanTime = block.timestamp;
                // newLoans.flag = 1;
                // loans[msg.sender][nowNumer] = newLoans;
                // loansNumber[msg.sender] += 1;
                // nftContract.approve(address(this), _tokenId);
      
}
//赎回nft
function withdrawNFT(uint256 _tokenId) onlyManager external payable {
    nftContract = IERC721(nft);
    uint256 nowNumer = loansNumber[msg.sender];
    for (uint i = 0; i < nowNumer; ++i ) {
        if (loans[msg.sender][i].tokenId == _tokenId) {
            require(loans[msg.sender][i].flag == 1, "You already have withdraw");   
             loans[msg.sender][_tokenId].flag  = 0;
             nftContract.approve(msg.sender, _tokenId);
        }
    }
   
}
function ownerOf(uint256  _tokenId) public returns (address){
                nftContract = IERC721(nft);
                return nftContract.ownerOf(_tokenId);
      
}
function loansList(uint start,uint limit,address user) public returns (uint256[] memory){
    nftContract = IERC721(nft);
    uint256 nowNumer = loansNumber[user];
    uint l = 0;
    uint256[] memory list;
    for (uint i = 0; i < nowNumer; ++i ) {
            if(loans[user][i].flag == 1){
                if (l >= start && l < start + limit){
                list[l] = loans[user][i].tokenId;
                l++;
                }
                if(l > start + limit){
                    return list;
                }
            }   
           
    }
    return list;
}
function loansCount(address user) public view returns (uint256){
    uint256 nowNumer = loansNumber[user];
    for (uint i = 0; i < nowNumer; ++i ) {
            if(loans[user][i].flag == 0){
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
// @title array limit library
// @author yueliyangzi
library AddressArrayLimitOnee{
    function pushLimit(address[31] memory origin, address input) internal pure returns (address[31] memory result) {
        for(
          uint i = 0; i < 30; i++
            ){
             origin[i] = origin[i+1];   
        }
        origin[30] = input;
        return origin;
    }
}
// @title array library
// @author yueliyangzi
library AddressArrayOnee{
    function deleteAddress(address[] memory origin,address pointer) internal pure returns(address[] memory result){
        if(containAddress(origin,pointer)){
        uint index;
        for(
            uint i = 0;i < origin.length; i++
            ){
              if(origin[i] == pointer){
                  index = i;
              }    
            }
        delete origin[index];
        for(
            uint i = index;i < origin.length-1;i++
            ){
                origin[i]=origin[i+1];
            }
        }
        return origin;
    }
    function containAddress(address[] memory origin,address pointer) internal pure returns(bool){
          for(
            uint i = 0;i < origin.length; i++
            ){
              if(origin[i] == pointer){
                  return true;
              }    
            }
            return false;
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


