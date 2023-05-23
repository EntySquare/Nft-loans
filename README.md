Token ID 1-500 AAA  0.7%/day
Token ID 501-3500 AA  0.6%/day
Token ID 3501-10000 A  0.5%/day


/app/login  登录注册
发送 {
"wallet_address": "0xc0822561B310256Aef0032e09b149Ac7cD7b5D55"
}
返回
{
}
/app/myCovenantFlow 我的收益
request
发送 {
}
response
返回 {
"benefit_info": {
"balance":100,
"last_day_benefit":20,
"accumulated_benefit":200
},
"covenant_flows":[
{
"nft_name":"whitetiger",
"pledge_id":"hxdawr",
"interest_rate":0.06,
"accumulated_benefit":200.06,
"pledge_fee":20.06,
"release_fee":10.06,
"release_fee":10.06,
"start_time":1678355904,
"expire_time":1678355904,
"nft_release_time":1678355904,
"flag":"1"
}]
}
/app/myNgt  我的ngt钱包
request
{

}
response
{
"benefit_info": {
"balance":100,
"last_day_benefit":20,
"accumulated_benefit":200
},
"transactions":[
{
"num":123,
"chain":"poly",
"address":"0xc0822561B310256Aef0032e09b149Ac7cD7b5D55"",
"hash":"0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
"ask_for_time":1678355904,
"achieve_time":1678355904,
"transaction_type":1,
"flag":"1"
}]
}
/app/myInvestment 我的邀请
request
{

}
response
{
"uid": "1",
"level": 2,
"accumulated_pledge_count": 2,
"investment_count": 2,
"investment_address": "0xc0822561B310256Aef0032e09b149Ac7cD7b5D55",
"investment_users":[
{
"uid": "3",
"level": 2,
"pledge_count": 2,
}]
}

/app/wallet/deposit  充值
request
{

}
response
{

}
/app/wallet/withdraw  提现
request
{

}
response
{

}

/app/nft/approve  授权
