//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: proto/challenge.proto
namespace proto
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"s2c_challenge")]
  public partial class s2c_challenge : global::ProtoBuf.IExtensible
  {
    public s2c_challenge() {}
    
    private string _challenge;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"challenge", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string challenge
    {
      get { return _challenge; }
      set { _challenge = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
}