# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: blockchain.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import transaction_pb2 as transaction__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10\x62lockchain.proto\x12\x06pactus\x1a\x11transaction.proto\"-\n\x11GetAccountRequest\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\"C\n\x12GetAccountResponse\x12-\n\x07\x61\x63\x63ount\x18\x01 \x01(\x0b\x32\x13.pactus.AccountInfoR\x07\x61\x63\x63ount\"\x1e\n\x1cGetValidatorAddressesRequest\"=\n\x1dGetValidatorAddressesResponse\x12\x1c\n\taddresses\x18\x01 \x03(\tR\taddresses\"/\n\x13GetValidatorRequest\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\"5\n\x1bGetValidatorByNumberRequest\x12\x16\n\x06number\x18\x01 \x01(\x05R\x06number\"K\n\x14GetValidatorResponse\x12\x33\n\tvalidator\x18\x01 \x01(\x0b\x32\x15.pactus.ValidatorInfoR\tvalidator\"_\n\x0fGetBlockRequest\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\x12\x34\n\tverbosity\x18\x02 \x01(\x0e\x32\x16.pactus.BlockVerbosityR\tverbosity\"\x83\x02\n\x10GetBlockResponse\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\x12\x12\n\x04hash\x18\x02 \x01(\x0cR\x04hash\x12\x12\n\x04\x64\x61ta\x18\x03 \x01(\x0cR\x04\x64\x61ta\x12\x1d\n\nblock_time\x18\x04 \x01(\rR\tblockTime\x12/\n\x06header\x18\x05 \x01(\x0b\x32\x17.pactus.BlockHeaderInfoR\x06header\x12\x34\n\tprev_cert\x18\x06 \x01(\x0b\x32\x17.pactus.CertificateInfoR\x08prevCert\x12)\n\x03txs\x18\x07 \x03(\x0b\x32\x17.pactus.TransactionInfoR\x03txs\"-\n\x13GetBlockHashRequest\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\"*\n\x14GetBlockHashResponse\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\"+\n\x15GetBlockHeightRequest\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\"0\n\x16GetBlockHeightResponse\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\"\x1a\n\x18GetBlockchainInfoRequest\"\xd5\x02\n\x19GetBlockchainInfoResponse\x12*\n\x11last_block_height\x18\x01 \x01(\rR\x0flastBlockHeight\x12&\n\x0flast_block_hash\x18\x02 \x01(\x0cR\rlastBlockHash\x12%\n\x0etotal_accounts\x18\x03 \x01(\x05R\rtotalAccounts\x12)\n\x10total_validators\x18\x04 \x01(\x05R\x0ftotalValidators\x12\x1f\n\x0btotal_power\x18\x05 \x01(\x03R\ntotalPower\x12\'\n\x0f\x63ommittee_power\x18\x06 \x01(\x03R\x0e\x63ommitteePower\x12H\n\x14\x63ommittee_validators\x18\x07 \x03(\x0b\x32\x15.pactus.ValidatorInfoR\x13\x63ommitteeValidators\"\x19\n\x17GetConsensusInfoRequest\"O\n\x18GetConsensusInfoResponse\x12\x33\n\tinstances\x18\x01 \x03(\x0b\x32\x15.pactus.ConsensusInfoR\tinstances\"\xad\x02\n\rValidatorInfo\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\x0cR\x04\x64\x61ta\x12\x1d\n\npublic_key\x18\x03 \x01(\tR\tpublicKey\x12\x16\n\x06number\x18\x04 \x01(\x05R\x06number\x12\x14\n\x05stake\x18\x05 \x01(\x03R\x05stake\x12.\n\x13last_bonding_height\x18\x06 \x01(\rR\x11lastBondingHeight\x12\x32\n\x15last_sortition_height\x18\x07 \x01(\rR\x13lastSortitionHeight\x12)\n\x10unbonding_height\x18\x08 \x01(\rR\x0funbondingHeight\x12\x18\n\x07\x61\x64\x64ress\x18\t \x01(\tR\x07\x61\x64\x64ress\"\x81\x01\n\x0b\x41\x63\x63ountInfo\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\x0cR\x04\x64\x61ta\x12\x16\n\x06number\x18\x03 \x01(\x05R\x06number\x12\x18\n\x07\x62\x61lance\x18\x04 \x01(\x03R\x07\x62\x61lance\x12\x18\n\x07\x61\x64\x64ress\x18\x05 \x01(\tR\x07\x61\x64\x64ress\"\xc4\x01\n\x0f\x42lockHeaderInfo\x12\x18\n\x07version\x18\x01 \x01(\x05R\x07version\x12&\n\x0fprev_block_hash\x18\x02 \x01(\x0cR\rprevBlockHash\x12\x1d\n\nstate_root\x18\x03 \x01(\x0cR\tstateRoot\x12%\n\x0esortition_seed\x18\x04 \x01(\x0cR\rsortitionSeed\x12)\n\x10proposer_address\x18\x05 \x01(\tR\x0fproposerAddress\"\x97\x01\n\x0f\x43\x65rtificateInfo\x12\x12\n\x04hash\x18\x01 \x01(\x0cR\x04hash\x12\x14\n\x05round\x18\x02 \x01(\x05R\x05round\x12\x1e\n\ncommitters\x18\x03 \x03(\x05R\ncommitters\x12\x1c\n\tabsentees\x18\x04 \x03(\x05R\tabsentees\x12\x1c\n\tsignature\x18\x05 \x01(\x0cR\tsignature\"{\n\x08VoteInfo\x12$\n\x04type\x18\x01 \x01(\x0e\x32\x10.pactus.VoteTypeR\x04type\x12\x14\n\x05voter\x18\x02 \x01(\tR\x05voter\x12\x1d\n\nblock_hash\x18\x03 \x01(\x0cR\tblockHash\x12\x14\n\x05round\x18\x04 \x01(\x05R\x05round\"\x97\x01\n\rConsensusInfo\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\x12\x16\n\x06\x41\x63tive\x18\x02 \x01(\x08R\x06\x41\x63tive\x12\x16\n\x06height\x18\x03 \x01(\rR\x06height\x12\x14\n\x05round\x18\x04 \x01(\x05R\x05round\x12&\n\x05votes\x18\x05 \x03(\x0b\x32\x10.pactus.VoteInfoR\x05votes*H\n\x0e\x42lockVerbosity\x12\x0e\n\nBLOCK_DATA\x10\x00\x12\x0e\n\nBLOCK_INFO\x10\x01\x12\x16\n\x12\x42LOCK_TRANSACTIONS\x10\x02*\\\n\x08VoteType\x12\x10\n\x0cVOTE_UNKNOWN\x10\x00\x12\x10\n\x0cVOTE_PREPARE\x10\x01\x12\x12\n\x0eVOTE_PRECOMMIT\x10\x02\x12\x18\n\x14VOTE_CHANGE_PROPOSER\x10\x03\x32\xe9\x05\n\nBlockchain\x12=\n\x08GetBlock\x12\x17.pactus.GetBlockRequest\x1a\x18.pactus.GetBlockResponse\x12I\n\x0cGetBlockHash\x12\x1b.pactus.GetBlockHashRequest\x1a\x1c.pactus.GetBlockHashResponse\x12O\n\x0eGetBlockHeight\x12\x1d.pactus.GetBlockHeightRequest\x1a\x1e.pactus.GetBlockHeightResponse\x12X\n\x11GetBlockchainInfo\x12 .pactus.GetBlockchainInfoRequest\x1a!.pactus.GetBlockchainInfoResponse\x12U\n\x10GetConsensusInfo\x12\x1f.pactus.GetConsensusInfoRequest\x1a .pactus.GetConsensusInfoResponse\x12\x43\n\nGetAccount\x12\x19.pactus.GetAccountRequest\x1a\x1a.pactus.GetAccountResponse\x12I\n\x0cGetValidator\x12\x1b.pactus.GetValidatorRequest\x1a\x1c.pactus.GetValidatorResponse\x12Y\n\x14GetValidatorByNumber\x12#.pactus.GetValidatorByNumberRequest\x1a\x1c.pactus.GetValidatorResponse\x12\x64\n\x15GetValidatorAddresses\x12$.pactus.GetValidatorAddressesRequest\x1a%.pactus.GetValidatorAddressesResponseBE\n\x11pactus.blockchainZ0github.com/pactus-project/pactus/www/grpc/pactusb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'blockchain_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021pactus.blockchainZ0github.com/pactus-project/pactus/www/grpc/pactus'
  _BLOCKVERBOSITY._serialized_start=2532
  _BLOCKVERBOSITY._serialized_end=2604
  _VOTETYPE._serialized_start=2606
  _VOTETYPE._serialized_end=2698
  _GETACCOUNTREQUEST._serialized_start=47
  _GETACCOUNTREQUEST._serialized_end=92
  _GETACCOUNTRESPONSE._serialized_start=94
  _GETACCOUNTRESPONSE._serialized_end=161
  _GETVALIDATORADDRESSESREQUEST._serialized_start=163
  _GETVALIDATORADDRESSESREQUEST._serialized_end=193
  _GETVALIDATORADDRESSESRESPONSE._serialized_start=195
  _GETVALIDATORADDRESSESRESPONSE._serialized_end=256
  _GETVALIDATORREQUEST._serialized_start=258
  _GETVALIDATORREQUEST._serialized_end=305
  _GETVALIDATORBYNUMBERREQUEST._serialized_start=307
  _GETVALIDATORBYNUMBERREQUEST._serialized_end=360
  _GETVALIDATORRESPONSE._serialized_start=362
  _GETVALIDATORRESPONSE._serialized_end=437
  _GETBLOCKREQUEST._serialized_start=439
  _GETBLOCKREQUEST._serialized_end=534
  _GETBLOCKRESPONSE._serialized_start=537
  _GETBLOCKRESPONSE._serialized_end=796
  _GETBLOCKHASHREQUEST._serialized_start=798
  _GETBLOCKHASHREQUEST._serialized_end=843
  _GETBLOCKHASHRESPONSE._serialized_start=845
  _GETBLOCKHASHRESPONSE._serialized_end=887
  _GETBLOCKHEIGHTREQUEST._serialized_start=889
  _GETBLOCKHEIGHTREQUEST._serialized_end=932
  _GETBLOCKHEIGHTRESPONSE._serialized_start=934
  _GETBLOCKHEIGHTRESPONSE._serialized_end=982
  _GETBLOCKCHAININFOREQUEST._serialized_start=984
  _GETBLOCKCHAININFOREQUEST._serialized_end=1010
  _GETBLOCKCHAININFORESPONSE._serialized_start=1013
  _GETBLOCKCHAININFORESPONSE._serialized_end=1354
  _GETCONSENSUSINFOREQUEST._serialized_start=1356
  _GETCONSENSUSINFOREQUEST._serialized_end=1381
  _GETCONSENSUSINFORESPONSE._serialized_start=1383
  _GETCONSENSUSINFORESPONSE._serialized_end=1462
  _VALIDATORINFO._serialized_start=1465
  _VALIDATORINFO._serialized_end=1766
  _ACCOUNTINFO._serialized_start=1769
  _ACCOUNTINFO._serialized_end=1898
  _BLOCKHEADERINFO._serialized_start=1901
  _BLOCKHEADERINFO._serialized_end=2097
  _CERTIFICATEINFO._serialized_start=2100
  _CERTIFICATEINFO._serialized_end=2251
  _VOTEINFO._serialized_start=2253
  _VOTEINFO._serialized_end=2376
  _CONSENSUSINFO._serialized_start=2379
  _CONSENSUSINFO._serialized_end=2530
  _BLOCKCHAIN._serialized_start=2701
  _BLOCKCHAIN._serialized_end=3446
# @@protoc_insertion_point(module_scope)
