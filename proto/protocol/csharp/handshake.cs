//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: proto/handshake.proto
namespace proto
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"c2s_handshake")]
  public partial class c2s_handshake : global::ProtoBuf.IExtensible
  {
    public c2s_handshake() {}
    
    private string _encode_challenge;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"encode_challenge", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string encode_challenge
    {
      get { return _encode_challenge; }
      set { _encode_challenge = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"s2c_handshake")]
  public partial class s2c_handshake : global::ProtoBuf.IExtensible
  {
    public s2c_handshake() {}
    
    private int _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int code
    {
      get { return _code; }
      set { _code = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
}