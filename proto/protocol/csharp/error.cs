//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: proto/error.proto
namespace proto
{
    [global::ProtoBuf.ProtoContract(Name=@"error")]
    public enum error
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"SUCCESS", Value=0)]
      SUCCESS = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"INTERNAL", Value=1)]
      INTERNAL = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"UNKNOWN_CMD", Value=2)]
      UNKNOWN_CMD = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PROTO_UNSERIALIZATION_FAILED", Value=3)]
      PROTO_UNSERIALIZATION_FAILED = 3,
            
      [global::ProtoBuf.ProtoEnum(Name=@"LOGIN_CLIENT_KEY_LEN_ILLEGAL", Value=1000)]
      LOGIN_CLIENT_KEY_LEN_ILLEGAL = 1000,
            
      [global::ProtoBuf.ProtoEnum(Name=@"LOGIN_HANDSHAKE_FAILED", Value=1001)]
      LOGIN_HANDSHAKE_FAILED = 1001,
            
      [global::ProtoBuf.ProtoEnum(Name=@"LOGIN_PROCESSING_IN_OTHER_PLACE", Value=1002)]
      LOGIN_PROCESSING_IN_OTHER_PLACE = 1002,
            
      [global::ProtoBuf.ProtoEnum(Name=@"REGISTER_DB_ERR", Value=1003)]
      REGISTER_DB_ERR = 1003,
            
      [global::ProtoBuf.ProtoEnum(Name=@"CREATE_PLAYER_DB_ERR", Value=1004)]
      CREATE_PLAYER_DB_ERR = 1004,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ACCOUNT_PLAYER_NOT_EXIST", Value=1005)]
      ACCOUNT_PLAYER_NOT_EXIST = 1005,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PLAYER_ID_NOT_EXIT", Value=10000)]
      PLAYER_ID_NOT_EXIT = 10000,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PLAYER_NOT_LOGIN", Value=10001)]
      PLAYER_NOT_LOGIN = 10001,
            
      [global::ProtoBuf.ProtoEnum(Name=@"CHANGE_PLAYER_NAME_DB_ERR", Value=10002)]
      CHANGE_PLAYER_NAME_DB_ERR = 10002
    }
  
}