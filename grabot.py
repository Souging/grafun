from web3 import Web3
from eth_utils import to_checksum_address

# BSC rpc节点
bsc_provider = 'https://bsc-mainnet.core.chainstack.com/' 
caller_address = '你的地址'
private_key = '你的私钥' 

w3 = Web3(Web3.HTTPProvider(bsc_provider))
if w3.isConnected():
    print("Successfully connected to BSC")
else:
    print("Failed to connect to BSC")
    exit()
contract_address = '0x8341b19a2A602eAE0f22633b6da12E1B016E6451 '#gra盘内合约地址
def pizhun():
    token_abi = '''
    [
        {
        "constant": false,
        "inputs": [
            {
                "name": "_spender",
                "type": "address"
            },
            {
                "name": "_value",
                "type": "uint256"
            }
        ],
        "name": "approve",
        "outputs": [
            {
                "name": "",
                "type": "bool"
            }
        ],
        "payable":  false,
        "stateMutability": "nonpayable",
        "type": "function"
        }
    ]
    '''

    token_contract = w3.eth.contract(address=token_address, abi=token_abi)

    # 批准的代币数量
    amount_to_approve = 56289999919999999999999000000000000000000  # 使用您尝试出售的相同数量

    # 构建批准交易
    approve_txn = token_contract.functions.approve(
        contract_address,  
        amount_to_approve
    ).build_transaction({
        'from': caller_address,
        'gas': 100000,  # 如有需要，调整 gas 限制
        'gasPrice': w3.toWei('1', 'gwei'), 
        'nonce': w3.eth.get_transaction_count(caller_address),
    })
    signed_approve_txn = w3.eth.account.sign_transaction(approve_txn, private_key)
    approve_tx_hash = w3.eth.send_raw_transaction(signed_approve_txn.rawTransaction)
    approve_receipt = w3.eth.wait_for_transaction_receipt(approve_tx_hash)

    if approve_receipt['status'] == 1:
        print("批准成功！")
    else:
        print("批准失败！")

def sell(val):
    contract_abi = '''
    [
        {
        "inputs": [
            {
                "internalType": "address",
                "name": "varg0",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "varg1",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "varg2",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "varg3",
                "type": "address"
            }
        ],
        "name": "sell",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
        }
    ]
    '''
    contract = w3.eth.contract(address=contract_address, abi=contract_abi)
    amount =  w3.toWei(val, 'ether') 
    tsdkaddr = '0xd4c32eD71ab6f32d1BE170F2C4511cfAcc616c2c'
    min_refund = w3.toWei(0.00001, 'ether') 
    buy_txn = contract.functions.sell(
        token_address,
        amount,
        min_refund,
        tsdkaddr
    ).build_transaction({
        'from': caller_address,
        'gas': 200000,  
        'gasPrice': w3.toWei('1', 'gwei'), 
        'nonce': w3.eth.get_transaction_count(caller_address),
    })

    signed_txn = w3.eth.account.sign_transaction(buy_txn, private_key)
    tx_hash = w3.eth.send_raw_transaction(signed_txn.rawTransaction)
    print(f'交易哈希: {tx_hash.hex()}')
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f'交易收据: {receipt}')


def buy(val):
    contract_abi = '''
    [
        {
        "inputs": [
            {
                "internalType": "address",
                "name": "contractAddress",
                "type": "address"
            },
            {
                "internalType": "uint256",
                "name": "tokenAmount",
                "type": "uint256"
            },
            {
                "internalType": "address",
                "name": "varg2",
                "type": "address"
            }
        ],
        "name": "buy",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
        }
    ]
    '''

    
    contract = w3.eth.contract(address=contract_address, abi=contract_abi)
    amount = 79061231758838210569000
    tsdkaddr = '0xd4c32eD71ab6f32d1BE170F2C4511cfAcc616c2c'
    value = w3.toWei(val, 'ether') 
    buy_txn = contract.functions.buy(token_address,amount,tsdkaddr).build_transaction({'from': caller_address,'value': value,   'gas': 200000,    'gasPrice': w3.toWei('1', 'gwei'), 'nonce': w3.eth.get_transaction_count(caller_address),})
    #amount = w3.toWei(1405279.412, 'ether') 
    #print(amount)
    #return
    signed_txn = w3.eth.account.sign_transaction(buy_txn, private_key)
    tx_hash = w3.eth.send_raw_transaction(signed_txn.rawTransaction)
    print(f'交易哈希: {tx_hash.hex()}')
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)

    print(f'交易收据: {receipt}')

#内盘要打的token合约地址
token_address = to_checksum_address('0x5ef9f735dd08a6903424e51c2582bd9f986e2719')

def sellv1():
  pizhun()
  sell(6636886908.500962)#单位是币的数量
def buyv1():
  buy(0.025)#单位是bnb

#要买就用  buyv1()    要卖就用sellv1()



