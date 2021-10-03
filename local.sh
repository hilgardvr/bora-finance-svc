pabUrl=http://127.0.0.1:9080/api/new/contract/instance/
echo "setting BORA_PAB_URL to ${pabUrl}"
export BORA_PAB_URL=$pabUrl

minterPwd=/home/hilgard/workspace/Bora-Finance-Property-Sale/Minter.cid
echo "setting BORA_CID_MINTER_FILE to ${minterPwd}"
export BORA_CID_MINTER_FILE=$minterPwd

sellerPwd=/home/hilgard/workspace/Bora-Finance-Property-Sale/Seller.cid
echo "setting BORA_CID_SELLER_FILE to ${sellerPwd}"
export BORA_CID_SELLER_FILE=$sellerPwd

buyer2Pwd=/home/hilgard/workspace/Bora-Finance-Property-Sale/Buyer2.cid
echo "setting BORA_CID_BUYER2_FILE to ${buyer2Pwd}"
export BORA_CID_BUYER2_FILE=$buyer2Pwd

echo "setting BORA_MINTER_CID to Minter.cid"
export BORA_MINTER_CID=''

echo "setting BORA_SELLER_CID to Seller.cid"
export BORA_SELLER_CID=''
